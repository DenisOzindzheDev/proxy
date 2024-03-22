package main

import (
	"rproxy/internal/server"
	logger "rproxy/pkg"
)

func main() {
	logger.Infof("[PROXY]Starting proxy server...")
	if err := server.Run(); err != nil {
		logger.Errorf("error running server: %v", err)
	}
}
