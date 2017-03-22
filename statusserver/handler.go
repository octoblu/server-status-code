package statusserver

import (
	"net/http"
	"time"
)

type statusHandler struct {
	delay, statusCode int
}

func (handler *statusHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	time.Sleep(time.Duration(handler.delay) * time.Millisecond)
	response.WriteHeader(handler.statusCode)
	response.Write([]byte(http.StatusText(handler.statusCode)))
}
