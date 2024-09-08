package web

type ErrorResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Errors any    `json:"errors"`
}
