package scaffolt

import (
	"sync"

	"github.com/kildevaeld/blueprint/store/utils"
)

type Context interface {
	Get(key string) interface{}
	Set(key string, value interface{})
}

type context struct {
	target    string
	generator Generator
	m         utils.Map
	lock      sync.RWMutex
}

func (self *context) Set(key string, value interface{}) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.m.Set(key, value)
}

func (self *context) Get(key string) interface{} {
	self.lock.RLock()
	defer self.lock.RUnlock()
	return self.m.Get(key)
}

func NewContext(gen Generator, target string) Context {
	return &context{
		target:    target,
		generator: gen,
		m:         utils.NewMap(),
	}
}
