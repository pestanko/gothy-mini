package restutl

import (
	"github.com/pestanko/gothy-mini/pkg/rest/resp"
	"net/http"
)

func WrapErrHandler(status func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := status(w, r)
		if err != nil {
			handleErrFromHandler(w, r, err)
		}
	}
}

func handleErrFromHandler(w http.ResponseWriter, r *http.Request, err error) {
	switch err.(type) {
	case *resp.ErrorDto:
		WriteErrorResp(w, &resp.ErrorResp{
			StatusCode: http.StatusBadRequest,
			ErrDto:     *err.(*resp.ErrorDto),
		})
	case resp.ErrorResponse:
		WriteErrorResp(w, err.(resp.ErrorResponse))
	default:
		// TODO: merge these two funcs
		WriteErrResp(w, http.StatusInternalServerError, err)
	}
}
