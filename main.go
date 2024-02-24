package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

const spinitronBaseURL = "https://spinitron.com"

type SimpleProxy struct {
	Proxy *httputil.ReverseProxy
}

func NewProxy(token string) *SimpleProxy {
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

	// Only allow GET requests
	if req.Method != http.MethodGet {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("404 page not found\n"))
		return
	}

	s.Proxy.ServeHTTP(rw, req)
}

func main() {
	token := os.Getenv("SPINITRON_API_KEY")
	if token == "" {
		panic("SPINITRON_API_KEY is empty")
	}

	proxy := NewProxy(token)

	http.Handle("/api/", proxy)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
