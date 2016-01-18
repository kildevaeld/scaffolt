package engine

import (
	"errors"
	"io/ioutil"
	"log"
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

func (self *context) Target() string {
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

	log.Printf("Creating file %s", fullpath)
	dir := filepath.Dir(fullpath)

	if !Exists(dir) {
		log.Printf("  Creating dir %s", dir)
		os.MkdirAll(dir, 0755)
	}
	log.Printf("  Writing file %s", fullpath)
	e := ioutil.WriteFile(fullpath, content, 0755)

	if e != nil {
		return e
	}

	log.Printf("  Written %d", len(content))

	return nil
}

func (self *context) Locals() map[string]interface{} {
	return self.m
}

func NewContext(gen scaffolt.Generator, target string) scaffolt.Context {
	return &context{
		target:    target,
		generator: gen,
		m:         utils.NewMap(),
	}
}
