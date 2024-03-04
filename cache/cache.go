package cache

import (
	"fmt"

	"github.com/Yiling-J/theine-go"
	"github.com/wcbn/spinitron-proxy/api"
)

type CacheStore struct {
	Cache *theine.Cache[string, []byte]
}

func (cs *CacheStore) Init() {
	if cs.Cache != nil {
		return
	}

	cache, err := theine.NewBuilder[string, []byte](2000).RemovalListener(func(key string, value []byte, reason theine.RemoveReason) {

		if api.IsResourcePath(key) {
			cs.Cache.Range(func(key string, value []byte) bool {
				fmt.Println("ranging through cache...")
				return true
			})
		}

		// reason could be REMOVED/EVICTED/EXPIRED
		fmt.Println("Key:", key, "Value:", string(value))
	}).Build()

	if err != nil {
		panic(err)
	}

	cs.Cache = cache

}

// func (t *AppTransport) InvalidateCollection(s string) {

// 	t.Cache.Range(func(key string, value any) bool {
// 		if api.IsResourcePath(key) {
// 			t.Cache.Delete(key)
// 		}
// 		return true
// 	})
// }
