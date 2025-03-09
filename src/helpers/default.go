package helpers

import (
	"net/http"
)

func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"response": "OK",
	}
	SendResponse(w, data, "/health", http.StatusOK)
}
