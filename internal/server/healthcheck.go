package server

import "net/http"

// `PING is the default ping operation
// `Usage`: ping used to do health checks
// from external resources
func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
