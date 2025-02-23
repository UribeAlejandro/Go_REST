package handlers

import (
	"encoding/json"
	"go_rest.com/rest-ws/api/responses"
	"go_rest.com/rest-ws/api/server"
	"net/http"
)

func HomeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(responses.JsonResponse{
			Message: "Hello World!",
			Status:  true,
		})
	}
}
