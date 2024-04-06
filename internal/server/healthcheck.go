package server

import (
	"net/http"
	responses "rproxy/pkg/api_utils"
	logger "rproxy/pkg/logger"
)

// `PING is the default ping operation
// `Usage`: ping used to do health checks
// from external resources
func ping(w http.ResponseWriter, r *http.Request) {
	responses.StatusOk(w, map[string]string{
		"status":  "OK",
		"Message": "Here is a random message",
	})
	logger.Infof("[PROXY] ping health check request from %s", r.RemoteAddr)
}
