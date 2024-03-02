package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

const tokenName = "SPINITRON_API_KEY"
const spinitronBaseURL = "https://spinitron.com"

type SimpleProxy struct {
	Proxy *httputil.ReverseProxy
}

func NewProxy() *SimpleProxy {
	token := os.Getenv(tokenName)
	if token == "" {
		panic(tokenName + " environment variable is empty")
	}

	url, _ := url.Parse(spinitronBaseURL)

	s := &SimpleProxy{httputil.NewSingleHostReverseProxy(url)}
	baseDirector := s.Proxy.Director
	s.Proxy.Director = func(req *http.Request) {
		baseDirector(req)
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("accept", "application/json")
	}

	return s
}

func (s *SimpleProxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	s.Proxy.ServeHTTP(rw, req)
}

func main() {
	proxy := NewProxy()
	http.Handle("GET /api/", proxy)
	err := http.ListenAndServe(":8080", nil)
	panic(err)
}
