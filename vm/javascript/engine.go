package javascript

import (
	"github.com/kildevaeld/motto"
	_ "github.com/kildevaeld/motto/underscore"
	"github.com/kildevaeld/scaffolt"
	"github.com/robertkrimen/otto"
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
	/*val, err := self.engine.ToValue(ctx)
	if err != nil {
		return err
	}*/
	c, _ := self.engine.ToValue(&context{ctx, self.engine})
	//self.engine.Set("context", c)
	value, err := self.engine.Require(path, ctx.Source())

	if err != nil {
		return err
	}

	if value.IsFunction() {
		_, err := value.Call(otto.UndefinedValue(), c)
		if err != nil {
			return err
		}
	} else if value.IsObject() {
		v := value.Object()
		prop := "run"
		if Contains(v.Keys(), "run") {
			prop = "run"
		} else if Contains(v.Keys(), "Run") {
			prop = "Run"
		}
		_, err := v.Call(prop, c)
		if err != nil {
			return err
		}
	}

	return nil
}

func (self *engine) Close() {
	self.engine.RunLoop.Stop()
}

func NewEngine() scaffolt.ScriptEngine {
	return &engine{}
}
