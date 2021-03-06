package engine

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/Sirupsen/logrus"
	"github.com/hashicorp/go-multierror"
	"github.com/kildevaeld/scaffolt"
	"github.com/kildevaeld/scaffolt/vm/javascript"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

type generator struct {
	path        string
	desc        scaffolt.GeneratorDescription
	once        sync.Once
	tasks       []scaffolt.Task
	files       []scaffolt.File
	initialized bool
	engines     map[scaffolt.IntepreterType]scaffolt.ScriptEngine
	l           *logrus.Logger
}

func (self *generator) Root() string {
	return self.path
}

func (self *generator) Description() scaffolt.GeneratorDescription {
	return self.desc
}

func (self *generator) Run(path string) error {

	if !self.initialized {
		return errors.New("not initialized")
	}

	for k, engine := range self.engines {
		self.l.Printf("Starting engine: %s", k)
		if err := engine.Init(self); err != nil {
			return err
		}
		defer engine.Close()
	}

	ctx := NewContext(self, path, self.l.WithField("prefix", "context"))

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

	base := filepath.Base(path)
	ctx.Set("BASE", base)
	ctx.Set("TARGET_DIR", path)

	self.l.Printf("Running tasks")
	for _, task := range self.tasks {
		if err = task.Run(ctx); err != nil {
			os.RemoveAll(path)
			return err
		}
	}

	log.Printf("Running files")
	if err := self.runFiles(ctx, false); err != nil {
		os.RemoveAll(path)
		return err
	}

	return nil
}

func (self *generator) Engine(engine scaffolt.IntepreterType) scaffolt.ScriptEngine {
	return self.engines[engine]
}

func (self *generator) AddFile(file scaffolt.FileDescription) error {

	self.files = append(self.files, NewFile(file, self.l.WithField("prefix", "file")))
	return nil
}

func (self *generator) Init() error {

	self.l.Formatter = new(prefixed.TextFormatter)

	var result error

	self.once.Do(func() {
		self.l.Printf("Initializing script engines")
		self.engines = make(map[scaffolt.IntepreterType]scaffolt.ScriptEngine)
		self.engines[scaffolt.Javascript] = javascript.NewEngine()

		self.l.Printf("Initialize tasks")
		for _, desc := range self.desc.Tasks {

			task := NewTask(desc, self.l.WithFields(logrus.Fields{
				"prefix": "task:" + desc.Name,
				//"task":   desc.Name,
			}))

			if err := task.Init(self); err != nil {
				result = multierror.Append(result, err)
			}

			self.tasks = append(self.tasks, task)
		}

		for _, fileDesc := range self.desc.Files {
			self.l.Printf("Initializing file: %s", fileDesc.Source)
			file := NewFile(fileDesc, self.l.WithField("prefix", "file"))
			if err := file.Init(self); err != nil {
				result = multierror.Append(result, err)
			}
			self.files = append(self.files, file)
		}

		self.initialized = true

	})

	return result
}

func (self *generator) runFiles(ctx scaffolt.Context, dryrun bool) error {
	var wg sync.WaitGroup
	var lock sync.Mutex
	var err error
	for _, file := range self.files {
		wg.Add(1)
		ff := file
		go func() {
			defer wg.Done()
			if e := ff.Run(ctx, dryrun); e != nil {
				lock.Lock()
				err = multierror.Append(err, e)
				lock.Unlock()
			}
		}()
	}
	wg.Wait()

	return err

}

func NewGenerator(path string, description scaffolt.GeneratorDescription) scaffolt.Generator {
	return &generator{
		path: path,
		desc: description,
		l:    logrus.New(),
	}
}
