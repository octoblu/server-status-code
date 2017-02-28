package statusserver

import "net/http"

type statusHandler struct {
	statusCode int
}

func (handler *statusHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(handler.statusCode)
	response.Write([]byte(http.StatusText(handler.statusCode)))
}
