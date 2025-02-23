package handlers

import (
	"encoding/json"
	"go_rest.com/rest-ws/responses"
	"go_rest.com/rest-ws/server"
	"net/http"
)

func HealthHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(responses.JsonResponse{
			Message: "OK",
			Status:  true,
		})
	}
}
