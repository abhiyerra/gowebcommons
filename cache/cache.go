package cache

import (
	"encoding/json"
	"log"
)

type Cache struct {
	cache map[string][]byte
}

func (c *Cache) Get(cacheKey string, runner func() interface{}) []byte {
	if cacheKey == "" {
		return nil
	}

	if b, ok := c.cache[cacheKey]; ok {
		log.Println("Hitting cache:", cacheKey)
		return b
	}

	b, err := json.Marshal(runner())
	if err != nil {
		log.Println("error:", err)
	}

	c.cache[cacheKey] = b
	return c.cache[cacheKey]
}

func New() (c Cache) {
	c.cache = make(map[string][]byte)
	return
}
