package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pestanko/gothy-mini/pkg/cfg"
	"github.com/pestanko/gothy-mini/pkg/client"
	"github.com/pestanko/gothy-mini/pkg/rest/handler"
	"github.com/pestanko/gothy-mini/pkg/user"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
)

type restServer struct {
	config       *cfg.AppCfg
	data         cfg.DataTemplate
	userGetter   user.Getter
	clientGetter client.Getter
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
		s.registerApiAuthRoutes(r)
		s.registerApiUserRoutes(r)
		s.registerApiClientRoutes(r)
	})
}

func (s *restServer) registerApiAuthRoutes(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Use(middleware.NoCache)
		r.Post("/login/credentials", handler.HandleAuthLoginCredentials())
		r.Post("/login/token", handler.HandleAuthLoginApiToken())
	})
	r.Route("/oauth2", func(r chi.Router) {
		r.Get("/authorize", handler.HandleOAuth2Authorize())
		r.Post("/token", handler.HandleOAuth2Token())
	})
}

func (s *restServer) registerApiUserRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		// TODO: require admin auth
		r.Get("/", handler.HandleUserList(s.userGetter))
		r.Get("/{username}", handler.HandleUserGet(s.userGetter))
	})
}

func (s *restServer) registerApiClientRoutes(r chi.Router) {
	r.Route("/clients", func(r chi.Router) {
		// TODO: require admin auth
		r.Get("/", handler.HandleClientList(s.clientGetter))
		r.Get("/{clientId}", handler.HandleClientGet(s.clientGetter))
	})
}
