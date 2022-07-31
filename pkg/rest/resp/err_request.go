package resp

import (
	"fmt"
	"net/http"
)

// MkInvalidRequest make an invalid request error
func MkInvalidRequest(desc string) ErrorResponse {
	return MkErrResp(http.StatusBadRequest, desc)
}

// MkParamMissing invalid request - required parameter missing
func MkParamMissing(param string) ErrorResponse {
	return MkInvalidRequest(fmt.Sprintf("required paramter missing: %s", param))
}

// MkParamInvalid invalid parameter value provided
func MkParamInvalid(param, val string) ErrorResponse {
	return MkInvalidRequest(fmt.Sprintf("paramter invalid '%s': %s", param, val))
}
