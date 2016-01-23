package engine

import (
	"os"
	"os/exec"
	"sync"

	"github.com/Sirupsen/logrus"
	"github.com/kildevaeld/blueprint/store/utils"
	"github.com/kildevaeld/scaffolt"
)

type context struct {
	target    string
	generator scaffolt.Generator
	m         utils.Map
	lock      sync.RWMutex
	l         *logrus.Entry
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

func (self *context) Exec(c string, args ...string) error {
	cmd := exec.Command(c, args...)
	cmd.Env = os.Environ()
	cmd.Dir = self.target
	return cmd.Run()
}

func (self *context) Move(source, target string, interpolate bool) {

	self.generator.AddFile(scaffolt.FileDescription{
		Interpolate: interpolate,
		Source:      source,
		Target:      target,
	})

}

func (self *context) CreateFile(path string, content []byte) error {

	self.generator.AddFile(scaffolt.FileDescription{
		Interpolate: true,
		Target:      path,
		Content:     string(content),
	})

	return nil
}

func (self *context) Generator() scaffolt.Generator {
	return self.generator
}

func (self *context) Locals() map[string]interface{} {
	return self.m
}

func NewContext(gen scaffolt.Generator, target string, log *logrus.Entry) scaffolt.Context {
	return &context{
		target:    target,
		generator: gen,
		m:         utils.NewMap(),
		l:         log,
	}
}
