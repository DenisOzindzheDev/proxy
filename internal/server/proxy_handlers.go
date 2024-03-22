package server

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	logger "rproxy/pkg"
	"strings"
	"time"
)

func NewProxy(target *url.URL) *httputil.ReverseProxy {
	proxy := httputil.NewSingleHostReverseProxy(target)
	return proxy
}
func ProxyRequestHandler(proxy *httputil.ReverseProxy, url *url.URL, endpoint string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("[PROXY] Request recived from %s at %v", r.URL.String(), time.Now().UTC())
		r.URL.Host = url.Host
		r.URL.Scheme = url.Scheme
		r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
		r.Host = url.Host
		path := r.URL.Path
		r.URL.Path = strings.TrimLeft(path, endpoint)
		logger.Infof("[PROXY] Redirecting to request %v at %v", r.URL.String(), time.Now().UTC())
		proxy.ServeHTTP(w, r)
	}
}
