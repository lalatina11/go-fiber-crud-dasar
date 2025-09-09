package lib

type JsonResponse struct {
	Error   bool   `json:"error"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}
