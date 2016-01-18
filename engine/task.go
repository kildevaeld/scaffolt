package engine

import (
	"sync"

	"github.com/hashicorp/go-multierror"
	"github.com/kildevaeld/scaffolt"
	"github.com/kildevaeld/scaffolt/vm"
)

type Hook int

const (
	Before Hook = iota
	After
)

type task struct {
	desc      scaffolt.TaskDescription
	scripts   map[Hook]scaffolt.Script
	files     []scaffolt.File
	questions *Questions
	once      sync.Once
}

func (self *task) Run(ctx scaffolt.Context) error {

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

func (self *task) Init(g scaffolt.Generator) error {
	self.scripts = make(map[Hook]scaffolt.Script)

	if self.desc.Before.Path != "" {
		self.scripts[Before] = vm.NewScript(self.desc.Before)
	}
	if self.desc.After.Path != "" {
		self.scripts[After] = vm.NewScript(self.desc.After)
	}

	for _, script := range self.scripts {
		if err := script.Init(g); err != nil {
			return err
		}
	}

	for _, fileDesc := range self.desc.Files {
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
