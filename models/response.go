package models

type Response struct {
	Code    float64   `json:"code"`
	Message string    `json:"message"`
	Data    []Profile `json:"data"`
	Error   Error     `json:"error"`
}

type ResponseProfileList struct {
	Code    float64   `json:"code"`
	Message string    `json:"message"`
	Data    []Profile `json:"data"`
	Error   Error     `json:"error"`
}

type ResponseProfile struct {
	Code    float64 `json:"code"`
	Message string  `json:"message"`
	Data    Profile `json:"data"`
	Error   Error   `json:"error"`
}

type ResponseImage struct {
	Code    float64 `json:"code"`
	Message string  `json:"message"`
	Data    []Image `json:"data"`
	Error   Error   `json:"error"`
}
