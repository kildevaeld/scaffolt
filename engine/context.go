package engine

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/kildevaeld/blueprint/store/utils"
	"github.com/kildevaeld/scaffolt"
)

type context struct {
	target    string
	generator scaffolt.Generator
	m         utils.Map
	lock      sync.RWMutex
}

func (self *context) Source() string {
	return self.generator.(*generator).path
}

func (self *Context) Target() string {
	return self.target
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

func (self *context) CreateFile(path string, content []byte) error {
	fullpath := filepath.Join(self.target, path)
	if IsDir(fullpath) {
		return errors.New("exists")
	}

	base := filepath.Base(fullpath)

	if !Exists(base) {
		os.MkdirAll(base, 0755)
	}

	ioutil.WriteFile(fullpath, content, 0755)

	return nil
}

func NewContext(gen scaffolt.Generator, target string) scaffolt.Context {
	return &context{
		target:    target,
		generator: gen,
		m:         utils.NewMap(),
	}
}
