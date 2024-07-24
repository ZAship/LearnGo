package demo

type Cache struct {
	Key_num   int
	Map_Cache map[string]string
}

func New(Max_key int) *Cache {
	return &Cache{
		Key_num:   Max_key,
		Map_Cache: make(map[string]string),
	}
}
func (c *Cache) Add(key string, value string) {
	if _, ok := c.Map_Cache[key]; ok {
		return
	}
	if len(c.Map_Cache) == c.Key_num {
		return
	}
	c.Map_Cache[key] = value
	return
}
func (c *Cache) Delete(key string) {
	if _, ok := c.Map_Cache[key]; !ok {
		return
	}
	delete(c.Map_Cache, key)
	return
}
