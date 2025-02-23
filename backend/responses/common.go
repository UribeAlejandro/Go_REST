package responses

type JsonResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}
