package statusserver

import (
	"fmt"
	"net/http"
)

type statusServer struct {
	delay, port, statusCode int
}

func (server *statusServer) Run() error {
	addr := fmt.Sprintf(":%v", server.port)
	handler := &statusHandler{delay: server.delay, statusCode: server.statusCode}
	return http.ListenAndServe(addr, handler)
}
