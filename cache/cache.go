package cache

import (
	"fmt"
	"time"

	"github.com/Yiling-J/theine-go"
	"github.com/wcbn/spinitron-proxy/api"
)

type Cache struct {
	tcache *theine.Cache[string, []byte]
}

func (c *Cache) evictCollection(name string) {
	c.tcache.Range(func(k string, v []byte) bool {
		fmt.Println(k, name, api.GetCollectionName(k))
		if api.GetCollectionName(k) == name {
			c.tcache.Delete(k)
		}
		return true
	})
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

// getTTL contains the cache expiration rules for each endpoint
func getTTL(key string) time.Duration {
	if api.IsResourcePath(key) {
		return 3 * time.Minute
	}

	return 5 * time.Second // for testing TODO remove this line

	// c := api.GetCollectionName(key)

	// var ttl = map[string]time.Duration{
	// 	"personas":  2 * time.Minute,
	// 	"shows":     2 * time.Minute,
	// 	"playlists": 2 * time.Minute,
	// 	"spins":     30 * time.Second,
	// }

	// return ttl[c]
}
