package javascript

import "github.com/kildevaeld/scaffolt"

type jsscript struct {
	path string
}

func (self jsscript) Type() scaffolt.IntepreterType {
	return scaffolt.Javascript
}

func (self *jsscript) Run(ctx scaffolt.Context) error {
	return nil
}

func (self jsscript) Init(g scaffolt.Generator) error {
	return nil
}

func New(path string) scaffolt.Script {
	return &jsscript{path}
}
