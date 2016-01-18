package scaffolt

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/hashicorp/go-multierror"
)

type Generator interface {
	Description() GeneratorDescription
	Init() error
	Run(path string) error
}

type generator struct {
	path  string
	desc  *GeneratorDescription
	once  sync.Once
	tasks []Task
}

func (self *generator) Description() GeneratorDescription {
	return *self.desc
}

func (self *generator) Run(path string) error {

	ctx := NewContext(self, path)

	var err error
	if path, err = filepath.Abs(path); err != nil {
		return err
	}

	if IsFile(path) || IsDir(path) {
		return fmt.Errorf("Target path already exists: %s", path)
	}

	if err = os.MkdirAll(path, 0755); err != nil {
		return err
	}

	for _, task := range self.tasks {
		if err = task.Run(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (self *generator) Init() error {
	if self.desc == nil {
		return errors.New("no description")
	}

	var result error

	self.once.Do(func() {

		for _, desc := range self.desc.Tasks {
			task := NewTask(desc)
			if err := task.Init(); err != nil {
				result = multierror.Append(result, err)
			}
			self.tasks = append(self.tasks, task)
		}

	})

	return result
}

func NewGenerator(path string, description *GeneratorDescription) Generator {
	return &generator{
		path: path,
		desc: description,
	}
}
