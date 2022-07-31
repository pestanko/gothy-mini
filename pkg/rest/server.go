package rest

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pestanko/gothy-mini/pkg/auth/session"
	"github.com/pestanko/gothy-mini/pkg/cfg"
	"github.com/pestanko/gothy-mini/pkg/client"
	"github.com/pestanko/gothy-mini/pkg/rest/handler"
	"github.com/pestanko/gothy-mini/pkg/rest/resp"
	"github.com/pestanko/gothy-mini/pkg/rest/restutl"
	"github.com/pestanko/gothy-mini/pkg/security"
	"github.com/pestanko/gothy-mini/pkg/user"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
	"time"
)

type restServer struct {
	config       *cfg.AppCfg
	data         cfg.DataTemplate
	userGetter   user.Getter
	clientGetter client.Getter
	pwdHasher    security.PasswordHasher
	sessionStore session.Store
}

func CreateResetServer(config cfg.AppCfg) http.Handler {
	server := makeWebServer(&config)
	r := chi.NewRouter()

	server.registerMiddleWares(r)
	server.registerRoutes(r)

	printRoutes(r)

	return r
}

func makeWebServer(config *cfg.AppCfg) restServer {
	data, err := cfg.LoadDataTemplate(config)
	if err != nil {
		log.Fatal().
			Str("env", cfg.Vars.Env).
			Str("data", config.Data.Load).
			Err(err).
			Msg("unable to load data template")
	}

	return restServer{
		config:       config,
		data:         data,
		userGetter:   user.NewGetter(data.Users),
		clientGetter: client.NewGetter(data.Clients),
		pwdHasher:    security.NewPasswordHasher(),
		sessionStore: session.NewStore(),
	}
}

func printRoutes(r *chi.Mux) {
	walkFunc := func(
		method string,
		route string,
		handler http.Handler,
		middlewares ...func(http.Handler) http.Handler,
	) error {
		route = strings.Replace(route, "/*/", "/", -1)
		log.Debug().
			Str("method", method).
			Str("route", route).
			Send()

		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		log.Error().
			Err(err).
			Msg("unable to print routes")
	}
}

func (s *restServer) registerMiddleWares(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(ChiZeroLog(&log.Logger))
	r.Use(middleware.Recoverer)
}

func (s *restServer) registerRoutes(r *chi.Mux) {
	r.Route("/healthz", func(r chi.Router) {
		r.Use(middleware.NoCache)
		r.Get("/", handler.HandleHealth())
	})

	r.Route("/api", func(r chi.Router) {
		r.Use(s.sessionMiddleware)
		s.registerApiAuthRoutes(r)
		s.registerApiUserRoutes(r)
		s.registerApiClientRoutes(r)
	})
}

func (s *restServer) registerApiAuthRoutes(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Use(middleware.NoCache)
		r.Post("/login/credentials", restutl.WrapErrHandler(handler.HandleAuthLoginCredentials(
			s.userGetter,
			s.pwdHasher,
			s.sessionStore,
		)))
		r.Post("/login/token", restutl.WrapErrHandler(handler.HandleAuthLoginApiToken()))
		r.Get("/session/status", restutl.WrapErrHandler(handler.HandleAuthSessionStatus()))
	})
	r.Route("/oauth2", func(r chi.Router) {
		r.Get("/authorize", restutl.WrapErrHandler(handler.HandleOAuth2Authorize()))
		r.Post("/token", restutl.WrapErrHandler(handler.HandleOAuth2Token(
			s.userGetter,
		)))
	})
}

func (s *restServer) registerApiUserRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Use(s.checkAccess(requireAdmin))
		r.Get("/", restutl.WrapErrHandler(handler.HandleUserList(s.userGetter)))
		r.Get("/{username}", restutl.WrapErrHandler(handler.HandleUserGet(s.userGetter)))
	})
}

func (s *restServer) registerApiClientRoutes(r chi.Router) {
	r.Route("/clients", func(r chi.Router) {
		r.Use(s.checkAccess(requireAdmin))
		r.Get("/", restutl.WrapErrHandler(handler.HandleClientList(s.clientGetter)))
		r.Get("/{clientId}", restutl.WrapErrHandler(handler.HandleClientGet(s.clientGetter)))
	})
}

func (s *restServer) sessionMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		sess, err := restutl.ExtractSessionFromReqCookie(r, s.sessionStore)
		if err != nil {
			log.Warn().Err(err).Msg("unable to get a session")
		}
		if err == nil && sess != nil && sess.IsValid(time.Now()) {
			ctx = context.WithValue(ctx, restutl.SessionKey, sess)
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

func (s *restServer) checkAccess(
	sessValidate func(sess *session.Session) bool,
) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			sess := restutl.GetSessionFromReq(r)
			if sess == nil {
				restutl.WriteErrorResp(w, resp.MkUnauthorized())
				return
			}
			if !sessValidate(sess) {
				restutl.WriteErrorResp(w, resp.MkForbidden())
				return
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

func requireAdmin(sess *session.Session) bool {
	return sess.UserType == user.TypeAdmin
}
