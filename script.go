package scaffolt

type Script interface {
	Type() IntepreterType
	Run(ctx Context) error
	Init() error
}

type jsscript struct {
	path string
}

func (self jsscript) Type() IntepreterType {
	return Javascript
}

func (self *jsscript) Run(ctx Context) error {
	return nil
}

func (self jsscript) Init() error {
	return nil
}

func NewScript(desc ScriptDescription) Script {
	switch desc.Type {
	case Javascript:
		return &jsscript{desc.Path}
	}
	return nil
}
