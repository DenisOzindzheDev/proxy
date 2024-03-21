package server

import (
	"fmt"
	"net/http"
	"net/url"
	"rproxy/internal/configs"
)

func Run() error {
	config, err := configs.NewConfig()
	if err != nil {
		return fmt.Errorf("Can't create config: %v", err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping)
	for _, resource := range config.Resources {
		url, _ := url.Parse(resource.Destination_URL)
		proxy := NewProxy(url)
		mux.HandleFunc(resource.Endpoint, ProxyRequestHandler(proxy, url, resource.Endpoint))
	}
	if err := http.ListenAndServe(config.Server.Host+config.Server.Listen_port, mux); err != nil {
		return fmt.Errorf("Can't start server: %v", err)
	}
	return nil
}
