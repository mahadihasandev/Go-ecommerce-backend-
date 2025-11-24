package util

import (
	"encoding/json"
	"net/http"
)

func SendData(write http.ResponseWriter, data interface{}, statusCode int) {
	write.WriteHeader(statusCode)
	encoder := json.NewEncoder(write)
	encoder.Encode(data)
}