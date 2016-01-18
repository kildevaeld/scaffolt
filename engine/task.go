package engine

import (
	"log"
	"strings"
	"sync"

	"github.com/hashicorp/go-multierror"
	"github.com/kildevaeld/blueprint/store/utils"
	"github.com/kildevaeld/scaffolt"
	"github.com/kildevaeld/scaffolt/vm"
)

type Hook int

const (
	Before Hook = iota
	After
)

type taskContext struct {
	ctx  scaffolt.Context
	task string
}

func (self *taskContext) Source() string {
	return self.ctx.Source()
}

func (self *taskContext) Target() string {
	return self.ctx.Target()
}

func (self *taskContext) Set(key string, value interface{}) {
	fk := key
	if !strings.HasPrefix(key, self.task+".") {
		fk = self.task + "." + key
	}
	self.ctx.Set(fk, value)
}

func (self *taskContext) Get(key string) interface{} {
	fk := key
	if key != self.task && !strings.Contains(key, ".") {
		fk = self.task + "." + key
	}
	return self.ctx.Get(fk)
}

func (self *taskContext) CreateFile(path string, content []byte) error {
	return self.ctx.CreateFile(path, content)
}

func (self *taskContext) Locals() map[string]interface{} {
	return self.ctx.Locals()
}

func (self *taskContext) Exec(c string, args ...string) error {
	return self.ctx.Exec(c, args...)
}

func (self *taskContext) Move(source, target string, interpolate bool) {
	self.ctx.Move(source, target, interpolate)
}

type task struct {
	desc      scaffolt.TaskDescription
	scripts   map[Hook]scaffolt.Script
	files     []scaffolt.File
	questions *Questions
	once      sync.Once
}

func (self *task) Run(ctx scaffolt.Context) error {
	ctx.Set(self.desc.Name, utils.NewMap())
	ctx = &taskContext{ctx, self.desc.Name}

	if err := self.runHook(Before, ctx); err != nil {
		return err
	}

	if err := self.questions.Run(ctx); err != nil {
		return err
	}

	if err := self.runFiles(ctx); err != nil {
		return err
	}

	if err := self.runHook(After, ctx); err != nil {
		return err
	}

	return nil
}

func (self *task) runFiles(ctx scaffolt.Context) error {
	var wg sync.WaitGroup
	var lock sync.Mutex
	var err error
	for _, file := range self.files {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if e := file.Run(ctx); e != nil {
				lock.Lock()
				err = multierror.Append(err, e)
				lock.Unlock()
			}
		}()
	}
	wg.Wait()

	return err

}

func (self *task) runHook(hook Hook, ctx scaffolt.Context) error {

	if script, ok := self.scripts[hook]; ok {
		return script.Run(ctx)
	}

	return nil
}

func (self *task) Name() string {
	return self.desc.Name
}

func (self *task) Init(g scaffolt.Generator) error {
	log.Printf("Initialize task: %s\n", self.desc.Name)
	self.scripts = make(map[Hook]scaffolt.Script)

	if self.desc.Before.Path != "" {
		log.Printf("  Adding before: script %s, type: %s\n", self.desc.Before.Path, self.desc.Before.Type)
		self.scripts[Before] = vm.NewScript(g.Engine(self.desc.Before.Type), self.desc.Before)
	}
	if self.desc.After.Path != "" {
		log.Printf("  Adding after: script %s, type: %s\n", self.desc.After.Path, self.desc.Before.Type)
		self.scripts[After] = vm.NewScript(g.Engine(self.desc.After.Type), self.desc.After)
	}

	for _, script := range self.scripts {
		log.Printf("  Initializing script: %s\n", script.Type())
		if err := script.Init(g); err != nil {
			return err
		}
	}

	for _, fileDesc := range self.desc.Files {
		log.Printf("  Initializing file: %s\n", fileDesc.Source)
		file := NewFile(fileDesc)
		if err := file.Init(g); err != nil {
			return err
		}
		self.files = append(self.files, file)
	}

	self.questions = NewQuestions(self.desc.Questions)
	if err := self.questions.Init(); err != nil {
		return err
	}

	return nil
}

func NewTask(desc scaffolt.TaskDescription) scaffolt.Task {
	return &task{
		desc: desc,
	}
}
