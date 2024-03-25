package server

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"rproxy/internal/configs"
	logger "rproxy/pkg"
)

func Run() error {
	config, err := configs.NewConfig()
	if err != nil {
		return fmt.Errorf("Can't create config: %v", err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping)
	if config.Mode == "nodb" {
		logger.Warnf("[PROXY] WARNING: running in nodb mode it may be not safe")
		for _, resource := range config.Resources {
			url, _ := url.Parse(resource.Destination_URL)
			proxy := NewProxy(url)
			mux.HandleFunc(resource.Endpoint, ProxyRequestHandler(proxy, url, resource.Endpoint))
		}
	}
	if config.Mode == "db" {

	} else {
		logger.Errorf("[PROXY] ERROR: invalid runing mode, got %s, waited = [db||nodb]", config.Mode)
		os.Exit(1)
	}

	if err := http.ListenAndServe(config.Server.Host+config.Server.Listen_port, mux); err != nil {
		return fmt.Errorf("Can't start server: %v", err)
	}

	return nil
}
