package scaffolt

type Context interface {
	Get(key string) interface{}
	Set(key string, value interface{})
	CreateFile(path string, content []byte) error
	Target() string
	Source() string
	Locals() map[string]interface{}
}

type File interface {
	Run(ctx Context) error
	Init(g Generator) error
}

type Script interface {
	Type() IntepreterType
	Run(ctx Context) error
	Init(g Generator) error
}

type Task interface {
	Run(ctx Context) error
	Init(g Generator) error
}

type Generator interface {
	Description() GeneratorDescription
	Init() error
	Run(path string) error
	Root() string
	Engine(IntepreterType) ScriptEngine
}

type ScriptEngine interface {
	Init(g Generator) error
	Run(path string, ctx Context) error
	Close()
}
