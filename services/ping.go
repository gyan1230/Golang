package services

import (
	"net/http"

	"github.com/gyan1230/Golang/utils"
)

//PingTest :
func PingTest() utils.APIResponse {
	return utils.APIResponse{
		HTTPCode:   http.StatusOK,
		Msg:        "Ping Sucessful",
		HTTPStatus: "StatusOK",
	}

}
