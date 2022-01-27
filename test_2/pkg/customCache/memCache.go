package customCache

type memCache struct {
	data map[string]interface{}
}

func NewMemCache() Cache {
	var cache memCache
	cache.data = make(map[string]interface{})
	return &cache
}

func (c *memCache) Get(key string) interface{} {
	return c.data[key]
}

func (c *memCache) Set(key string, value interface{}) error {
	c.data[key] = value
	return nil
}

func (c *memCache) Clear() {
	c.data = make(map[string]interface{})
	c.data.Clear()
}
