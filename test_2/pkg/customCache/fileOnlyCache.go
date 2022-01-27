package customCache

import (
	"encoding/json"
	"fmt"
	"os"
)

type fileOnlyCache struct {
	path string
	data map[string]interface{}
}

func NewFileOnlyCache(path string) (Cache, error) {
	cache := fileOnlyCache{path: path}

	err := os.MkdirAll(path, 0755)

	return &cache, err
}

func (c *fileOnlyCache) Get(key string) interface{} {
	filePath := fmt.Sprintf("%s/%s", c.path, key)

	serializedValue, readErr := os.ReadFile(filePath)

	if os.IsNotExist(readErr) {
		return nil
	}

	// file could not be read
	if readErr != nil {
		return nil
	}

	var value interface{}
	serialErr := json.Unmarshal(serializedValue, &value)

	// value could not be deserialized
	if serialErr != nil {
		return nil
	}

	return value
}

func (c *fileOnlyCache) Set(key string, value interface{}) error {
	filePath := fmt.Sprintf("%s/%s", c.path, key)

	serializedValue, serialErr := json.Marshal(value)
	if serialErr != nil {
		return serialErr
	}

	writeErr := os.WriteFile(filePath, serializedValue, 0644)

	return writeErr
}

func (c *fileOnlyCache) Clear() {
	os.RemoveAll(c.path)
	os.MkdirAll(c.path, 0755)
}
