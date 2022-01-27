package customCache

import (
	"encoding/json"
	"os"
)

type fileCache struct {
	memCache
	filePath string
}

func (c *fileCache) loadFromFile() error {
	serializedCache, readErr := os.ReadFile(c.filePath)

	if os.IsNotExist(readErr) {
		return nil
	}

	if readErr != nil {
		return readErr
	}

	serialErr := json.Unmarshal(serializedCache, &c.memCache.data)

	return serialErr
}

func (c *fileCache) writeToFile() error {
	serializedCache, serialErr := json.Marshal(c.memCache.data)
	if serialErr != nil {
		return serialErr
	}

	writeErr := os.WriteFile(c.filePath, serializedCache, 0644)
	return writeErr
}

func NewFileCache(path string) (Cache, error) {
	filePath := path + "/fileCache.json"

	cache := fileCache{filePath: filePath}
	cache.memCache.data = make(map[string]interface{})

	err := os.MkdirAll(path, 0755)

	if err != nil {
		return &cache, err
	}

	loadErr := cache.loadFromFile()

	return &cache, loadErr
}

func (c *fileCache) Get(key string) interface{} {
	return c.memCache.Get(key)
}

func (c *fileCache) Set(key string, value interface{}) error {
	c.memCache.Set(key, value)
	writeErr := c.writeToFile()

	return writeErr
}

func (c *fileCache) Clear() {
	c.memCache.Clear()
	os.Remove(c.filePath)
}
