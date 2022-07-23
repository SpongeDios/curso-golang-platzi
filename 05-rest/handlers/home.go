package handlers

import (
	"05-rest/server"
	"encoding/json"
	"net/http"
)

type HomeResponse struct {
	Message string `json:"message"`
	Status bool `json:"status"`
}

func HomeHandle(s server.Server) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		err := json.NewEncoder(writer).Encode(HomeResponse{
			Message: "Hola mundo",
			Status:  true,
		})
		if err != nil {
			return 
		}
	}
}