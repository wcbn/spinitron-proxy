package proxy

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/wcbn/spinitron-proxy/cache"
)

type TransportWithCache struct {
	Transport http.RoundTripper
	Cache     *cache.Cache
}

func (t *TransportWithCache) RoundTrip(req *http.Request) (*http.Response, error) {
	key := t.Cache.MakeCacheKey(req)
	value, found := t.Cache.Get(key)

	if found {
		resp := &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewReader(value)),
		}

		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	}

	// Do HTTP fetch from target server
	resp, err := t.Transport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return resp, err
	}

	var data []byte
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	resp.Body = io.NopCloser(bytes.NewReader(data))

	t.Cache.Set(key, data)

	return resp, err
}

func NewReverseProxy(tokenEnvVarName string, target *url.URL) *httputil.ReverseProxy {
	t := os.Getenv(tokenEnvVarName)
	if t == "" {
		panic(tokenEnvVarName + " environment variable is empty")
	}

	rp := httputil.NewSingleHostReverseProxy(target)

	d := rp.Director
	rp.Director = func(req *http.Request) {
		d(req)
		req.Header.Set("Authorization", "Bearer "+t)
		req.Header.Set("accept", "application/json")
	}

	c := &cache.Cache{}
	c.Init()

	rp.Transport = &TransportWithCache{
		Transport: http.DefaultTransport,
		Cache:     c,
	}

	return rp
}
