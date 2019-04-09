// Reference:
// https://github.com/patrickmn/go-cache
// https://stackoverflow.com/questions/25484122/map-with-ttl-option-in-go
// https://groups.google.com/g/golang-nuts/c/avSIKqUKKAM?pli=1
// https://golang.org/pkg/sync/#Map

package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// item is the data to be cached.
type item struct {
	value   string
	expires int64
}

// Cache is a struct for caching.
type Cache struct {
	items map[string]*item
	mu    sync.Mutex
}

func New() *Cache {
	c := &Cache{items: make(map[string]*item)}
	go func() {
		t := time.NewTicker(time.Second)
		defer t.Stop()
		for {
			select {
			case <-t.C:
				c.mu.Lock()
				for k, v := range c.items {
					if v.Expired(time.Now().UnixNano()) {
						log.Printf("%v has expires at %d", c.items, time.Now().UnixNano())
						delete(c.items, k)
					}
				}
				c.mu.Unlock()
			}
		}
	}()
	return c
}

// Expired determines if it has expires.
func (i *item) Expired(time int64) bool {
	if i.expires == 0 {
		return true
	}
	return time > i.expires
}

// Get gets a value from a cache.
func (c *Cache) Get(key string) string {
	c.mu.Lock()
	var s string
	if v, ok := c.items[key]; ok {
		s = v.value
	}
	c.mu.Unlock()
	return s
}

// Put puts a value to a cache. If a key and value exists, overwrite it.
func (c *Cache) Put(key string, value string, expires int64) {
	c.mu.Lock()
	if _, ok := c.items[key]; !ok {
		c.items[key] = &item{
			value:   value,
			expires: expires,
		}
	}
	c.mu.Unlock()
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fk := "first-key"
		sk := "second-key"

		cache := New()

		cache.Put(fk, "first-value", time.Now().Add(2*time.Second).UnixNano())
		fmt.Println(cache.Get(fk))

		time.Sleep(10 * time.Second)

		if len(cache.Get(fk)) == 0 {
			cache.Put(sk, "second-value", time.Now().Add(100*time.Second).UnixNano())
		}
		fmt.Println(cache.Get(sk))
	})
	http.ListenAndServe(":8080", nil)
}
