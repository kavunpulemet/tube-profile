package utils

import (
	"encoding/json"
	"net/http"
)

const (
	ContentType     = "Content-Type"
	ApplicationJSON = "application/json"
)

type StatusResponse struct {
	Status string `json:"status"`
}

func WriteResponse(w http.ResponseWriter, statusCode int, data any) error {
	w.Header().Set(ContentType, ApplicationJSON)
	//w.Header().Set("Access-Control-Allow-Origin", "*") // можно заменить на конкретный домен
	//w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	//w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(data)
}
