package main

import (
	"net/http"
	"net/url"

	"github.com/wcbn/spinitron-proxy/proxy"
)

const tokenEnvVarName = "SPINITRON_API_KEY"
const spinitronBaseURL = "https://spinitron.com"

func main() {
	url, _ := url.Parse(spinitronBaseURL)
	proxy := proxy.NewReverseProxy(tokenEnvVarName, url)
	http.Handle("GET /api/", proxy)
	http.Handle("GET /images/", proxy)
	err := http.ListenAndServe(":8080", nil)
	panic(err)
}
