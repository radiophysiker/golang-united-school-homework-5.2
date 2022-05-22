package cache

import "time"

type item struct {
	value   string
	expires time.Time
}

type Cache struct {
	mapCache map[string]item
}

func NewCache() Cache {
	return Cache{mapCache: make(map[string]item)}
}

func (c Cache) Get(key string) (string, bool) {
	item, ok := c.mapCache[key]
	if !ok {
		return "", false
	}
	if item.expires != (time.Time{}) && time.Now().After(item.expires) {
		delete(c.mapCache, key)
		return "", false
	}
	return item.value, true
}

func (c Cache) Put(key, value string) {
	c.mapCache[key] = item{
		expires: time.Time{},
		value:   value,
	}
}

func (c Cache) Keys() []string {
	var keys []string
	for key, item := range c.mapCache {
		if item.expires != (time.Time{}) && time.Now().After(item.expires) {
			delete(c.mapCache, key)
		} else {
			keys = append(keys, key)
		}

	}
	return keys
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	c.mapCache[key] = item{
		expires: deadline,
		value:   value,
	}
}
