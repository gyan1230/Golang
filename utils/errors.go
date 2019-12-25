package utils

import "net/http"

//NewBadrequest :
func NewBadrequest(msg string) *APIResponse {
	return &APIResponse{
		HTTPCode:   http.StatusBadRequest,
		Msg:        msg,
		HTTPStatus: "bad request",
	}
}

//InternalServerError :
func InternalServerError(msg string) *APIResponse {
	return &APIResponse{
		HTTPCode:   http.StatusInternalServerError,
		Msg:        msg,
		HTTPStatus: "Internal Server Error",
	}
}
