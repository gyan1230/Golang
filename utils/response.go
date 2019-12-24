package utils

//APIResponse :
type APIResponse struct {
	HTTPCode   int    `json:"code"`
	Msg        string `json:"message"`
	HTTPStatus string `json:"status_message"`
}

//APIData :
type APIData struct {
	Data APIResponse `json:"data"`
}
