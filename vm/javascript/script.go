package javascript

import (
	"strings"

	"github.com/kildevaeld/scaffolt"
)

type jsscript struct {
	path   string
	engine scaffolt.ScriptEngine
}

func (self jsscript) Type() scaffolt.IntepreterType {
	return scaffolt.Javascript
}

func (self *jsscript) Run(ctx scaffolt.Context) error {
	e := self.engine.Run(self.path, ctx)
	return e
}

func (self *jsscript) Init(g scaffolt.Generator) error {

	if strings.HasPrefix(self.path, "/") {
		self.path = "." + self.path
	} else if !strings.HasPrefix(self.path, "./") {
		self.path = "./" + self.path
	}

	return nil
}

func New(engine scaffolt.ScriptEngine, path string) scaffolt.Script {
	return &jsscript{
		engine: engine,
		path:   path,
	}
}
