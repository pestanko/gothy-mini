package resp

import "net/http"

func MkUnauthorized() ErrorResponse {
	return MkErrResp(http.StatusUnauthorized, "You need to log in")
}

func MkForbidden() ErrorResponse {
	return MkErrResp(http.StatusForbidden, "You are not authorized for this operation")
}
