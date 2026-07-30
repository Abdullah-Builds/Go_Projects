package cache

<<<<<<< HEAD
type Cache struct{}
=======
import (
	"sync"
	"time"
)

type Cache struct {
	data map[string]Item
	mu   sync.RWMutex
}

func New() *Cache {
	return &Cache{data: map[string]Item{}}
}

func (c *Cache) Set(key, value string, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item := Item{
		Value: value,
	}

	if ttl > 0 {
		item.ExpiresAt = time.Now().Add(ttl)
	}

	c.data[key] = item
}

func (c *Cache) Get(key string) (string, bool) {

	c.mu.RLock()
	defer c.mu.RUnlock()

	item, ok := c.data[key]
	if !ok {
		return "", false
	}

	if !item.ExpiresAt.IsZero() && time.Now().After(item.ExpiresAt) {
		delete(c.data, key)
		return "", false
	}

	return item.Value, true
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

// Now expired keys disappear automatically.
func (c *Cache) StartCleanup(interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		defer ticker.Stop()

		for range ticker.C {
			c.cleanupExpired()
		}
	}()
}

func (c *Cache) cleanupExpired() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()

	for key, item := range c.data {
		if !item.ExpiresAt.IsZero() && now.After(item.ExpiresAt) {
			delete(c.data, key)
		}
	}
}
>>>>>>> 43439c3 (feat: cache version 1)
