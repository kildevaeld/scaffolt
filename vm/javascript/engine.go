package javascript

import (
	"github.com/kildevaeld/motto"
	"github.com/kildevaeld/scaffolt"
)

type engine struct {
	engine *motto.Motto
}

func (self *engine) Init(g scaffolt.Generator) error {

	engine := motto.New()

	self.engine = engine

	go engine.RunLoop.Run()

	return nil
}

func (self *engine) Run(path string, ctx scaffolt.Context) error {

	self.engine.Require(path, ctx.Source())

	return nil
}

func (self *engine) Close() {
	self.engine.RunLoop.Stop()
}

func NewEngine() scaffolt.ScriptEngine {
	return &engine{}
}
