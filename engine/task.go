package engine

import (
	"strings"
	"sync"

	"github.com/Sirupsen/logrus"
	"github.com/kildevaeld/dict"
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

func (self *taskContext) Generator() scaffolt.Generator {
	return self.ctx.Generator()
}

type task struct {
	desc      scaffolt.TaskDescription
	scripts   map[Hook]scaffolt.Script
	files     []scaffolt.File
	questions *Questions
	once      sync.Once
	l         *logrus.Entry
}

func (self *task) Run(ctx scaffolt.Context) error {
	ctx.Set(self.desc.Name, dict.NewMap())
	ctx = &taskContext{ctx, self.desc.Name}

	if err := self.runHook(Before, ctx); err != nil {
		return err
	}
	self.l.Printf("Running questions")
	if err := self.questions.Run(ctx); err != nil {
		return err
	}
	self.l.Printf("Running files")
	if err := self.runFiles(ctx); err != nil {
		return err
	}

	if err := self.runHook(After, ctx); err != nil {
		return err
	}

	return nil
}

func (self *task) runFiles(ctx scaffolt.Context) error {
	/*var wg sync.WaitGroup
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
	*/
	return nil

}

func (self *task) runHook(hook Hook, ctx scaffolt.Context) error {
	var h string
	switch hook {
	case Before:
		h = "before"
	case After:
		h = "after"
	}

	if script, ok := self.scripts[hook]; ok {
		self.l.WithFields(logrus.Fields{
			"type": script.Type().String(),
		}).Printf("Running hook %s", h)
		return script.Run(ctx)
	}

	return nil
}

func (self *task) Name() string {
	return self.desc.Name
}

func (self *task) Init(g scaffolt.Generator) error {
	self.l.Printf("Initialize task: %s", self.desc.Name)
	self.scripts = make(map[Hook]scaffolt.Script)

	if self.desc.Before.Path != "" {
		self.l.Printf("Adding before: script %s, type: %s", self.desc.Before.Path, self.desc.Before.Type)
		self.scripts[Before] = vm.NewScript(g.Engine(self.desc.Before.Type), self.desc.Before)
	}
	if self.desc.After.Path != "" {
		self.l.Printf("Adding after: script %s, type: %s", self.desc.After.Path, self.desc.Before.Type)
		self.scripts[After] = vm.NewScript(g.Engine(self.desc.After.Type), self.desc.After)
	}

	for _, script := range self.scripts {
		self.l.Printf("Initializing script: %s", script.Type())
		if err := script.Init(g); err != nil {
			return err
		}
	}

	for _, fileDesc := range self.desc.Files {
		/*self.l.Printf("Initializing file: %s", fileDesc.Source)
		file := NewFile(fileDesc)
		if err := file.Init(g); err != nil {
			return err
		}
		self.files = append(self.files, file)*/
		g.AddFile(fileDesc)
	}

	self.questions = NewQuestions(self.desc.Questions)
	if err := self.questions.Init(); err != nil {
		return err
	}

	return nil
}

func NewTask(desc scaffolt.TaskDescription, l *logrus.Entry) scaffolt.Task {
	return &task{
		desc: desc,
		l:    l,
	}
}
