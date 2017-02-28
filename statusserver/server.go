package statusserver

import (
	"fmt"
	"net/http"
)

type statusServer struct {
	port, statusCode int
}

func (server *statusServer) Run() error {
	addr := fmt.Sprintf(":%v", server.port)
	handler := &statusHandler{statusCode: server.statusCode}
	return http.ListenAndServe(addr, handler)
}
