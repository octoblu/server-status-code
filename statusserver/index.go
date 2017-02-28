package statusserver

// Server accepts HTTP requests and responds with status codes
type Server interface {
	// Run runs the server and blocks until the server stops
	Run() error
}

// New constructs a new Server instance
func New(port, statusCode int) Server {
	return &statusServer{port: port, statusCode: statusCode}
}
