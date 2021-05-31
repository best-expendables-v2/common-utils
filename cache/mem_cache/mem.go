package mem_cache

import (
	commonCache "github.com/best-expendables-v2/common-utils/cache"
	"github.com/patrickmn/go-cache"
	"reflect"
	"time"
)

type Mem struct {
	c *cache.Cache
}

func NewMem(ttl time.Duration) *Mem {
	return &Mem{cache.New(ttl, 10*time.Minute)}
}

func (m *Mem) Get(key string, obj interface{}) error {
	value, found := m.c.Get(key)
	if found {
		v := reflect.ValueOf(obj).Elem()
		rv := reflect.ValueOf(value)
		if rv.Kind() == reflect.Ptr {
			rv = rv.Elem()
		}
		v.Set(rv)
		return nil
	}
	return commonCache.Nil
}

func (m *Mem) Set(key string, obj interface{}) error {
	m.c.Set(key, obj, cache.DefaultExpiration)
	return nil
}

func (m *Mem) Delete(key string) error {
	m.c.Delete(key)
	return nil
}
