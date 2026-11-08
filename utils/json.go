package utils

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(writer http.ResponseWriter, status int, obj any) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(obj)
}
