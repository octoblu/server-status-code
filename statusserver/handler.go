package statusserver

import (
	"net/http"
	"time"
)

type statusHandler struct {
	delay, statusCode int
}

func (handler *statusHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/healthcheck" {
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		response.Write([]byte("{\"online\": true}"))
	} else {
		time.Sleep(time.Duration(handler.delay) * time.Millisecond)
		response.WriteHeader(handler.statusCode)
		response.Write([]byte(http.StatusText(handler.statusCode)))
	}
}
