package cache

import (
	"net/http"
	"time"

	"github.com/Yiling-J/theine-go"
	"github.com/wcbn/spinitron-proxy/api"
)

type Cache struct {
	tcache *theine.Cache[string, []byte]
}

func (c *Cache) Init() {
	if c.tcache != nil {
		return
	}

	cache, err := theine.NewBuilder[string, []byte](2000).RemovalListener(func(k string, v []byte, r theine.RemoveReason) {

		// when a collection path expires
		if api.IsCollectionPath(k) && r == theine.EXPIRED {

			// remove all entries for said collection
			c.evictCollection(api.GetCollectionName(k))
		}

	}).Build()

	if err != nil {
		panic(err)
	}

	c.tcache = cache
}

func (c *Cache) Get(key string) ([]byte, bool) {
	return c.tcache.Get(key)
}

func (c *Cache) Set(key string, value []byte) bool {
	ttl := getTTL(key)
	return c.tcache.SetWithTTL(key, value, 0, ttl)
}

func (c *Cache) MakeCacheKey(req *http.Request) string {
	result := req.URL.Path
	if api.IsCollectionPath(result) {
		result += "?" + req.URL.Query().Encode()
	}
	return result
}

// getTTL contains the cache expiration rules for each endpoint
func getTTL(key string) time.Duration {
	if api.IsResourcePath(key) {
		return 3 * time.Minute
	}

	c := api.GetCollectionName(key)

	var ttl = map[string]time.Duration{
		"personas":  5 * time.Minute,
		"shows":     5 * time.Minute,
		"playlists": 3 * time.Minute,
		"spins":     30 * time.Second,
	}

	return ttl[c]
}

func (c *Cache) evictCollection(name string) {
	shouldDelete := []string{}
	c.tcache.Range(func(k string, v []byte) bool {
		if api.GetCollectionName(k) == name {
			shouldDelete = append(shouldDelete, k)
		}
		return true
	})
	for _, key := range shouldDelete {
		c.tcache.Delete(key)
	}
}
