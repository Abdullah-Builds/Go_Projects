package cache

import (
	"container/list"
	"encoding/json"
	"os"
	"time"
)

func (c *Cache) Save(filename string) error {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// convert internal map[string]*list.Element to map[string]Item for snapshot
	snapshotData := make(map[string]Item, len(c.data))
	for k, elem := range c.data {
		if elem == nil {
			continue
		}
		switch v := elem.Value.(type) {
		case Item:
			snapshotData[k] = v
		case *Item:
			snapshotData[k] = *v
		}
	}

	snapshot := Snapshot{
		Version: 1,
		SavedAt: time.Now(),
		Data:    snapshotData,
	}

	data, err := json.MarshalIndent(snapshot, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func (c *Cache) Load(filename string) error {

	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	var snapshot Snapshot

	if err := json.Unmarshal(data, &snapshot); err != nil {
		return err
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	// rebuild internal map from snapshot Data (map[string]Item) to map[string]*list.Element
	newData := make(map[string]*list.Element, len(snapshot.Data))
	for k, it := range snapshot.Data {
		item := it
		newData[k] = &list.Element{Value: item}
	}

	c.data = newData
	return nil
}

func (c *Cache) StartAutoSave(filename string, interval time.Duration) {
	go func() {

		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for range ticker.C {
			c.Save(filename)
		}

	}()
}
