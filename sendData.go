package main

import (
	"encoding/json"
	"net/http"
)

func sendData(write http.ResponseWriter, data interface{}, statusCode int) {
	write.WriteHeader(statusCode)
	encoder := json.NewEncoder(write)
	encoder.Encode(data)
}