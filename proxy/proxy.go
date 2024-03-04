package proxy

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"

	"github.com/wcbn/spinitron-proxy/api"
	"github.com/wcbn/spinitron-proxy/cache"
)

type TransportWithCache struct {
	Transport  http.RoundTripper
	CacheStore *cache.CacheStore
}

func (t *TransportWithCache) RoundTrip(req *http.Request) (*http.Response, error) {
	key := req.URL.Path
	value, found := t.CacheStore.Cache.Get(key)
	var data []byte

	if found {
		resp := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewReader(value)),
		}

		resp.Header.Set("Content-Type", "application/json")

		return resp, nil
	}

	fmt.Println("MISS")

	// Do HTTP fetch from target server
	resp, err := t.Transport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return resp, err
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	resp.Body = io.NopCloser(bytes.NewReader(data))

	if api.IsResourcePath(key) {
		t.CacheStore.Cache.SetWithTTL(key, data, 1, 3*time.Minute)
	}

	return resp, err
}

func NewReverseProxy(tokenEnvVarName string, target *url.URL) *httputil.ReverseProxy {
	token := os.Getenv(tokenEnvVarName)
	if token == "" {
		panic(tokenEnvVarName + " environment variable is empty")
	}

	rp := httputil.NewSingleHostReverseProxy(target)

	d := rp.Director
	rp.Director = func(req *http.Request) {
		d(req)
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("accept", "application/json")
	}

	cs := &cache.CacheStore{}
	cs.Init()

	rp.Transport = &TransportWithCache{
		Transport:  http.DefaultTransport,
		CacheStore: cs,
	}

	return rp
}
