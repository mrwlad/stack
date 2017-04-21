package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

// CreateProxy creates a new reverse proxy
func CreateProxy(target *url.URL) *httputil.ReverseProxy {
	targetQuery := target.RawQuery
	director := func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = singleJoiningSlash(target.Path, req.URL.Path)
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
		// set the x-extra custom header
		req.Header.Set("X-Extra", "hello")
	}
	return &httputil.ReverseProxy{Director: director}
}

func main() {
	proxy := CreateProxy(&url.URL{Scheme: "http", Host: "api:4000"})
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handler)
	mux.Handle("/", proxy)
	http.ListenAndServe(":8080", mux)
}