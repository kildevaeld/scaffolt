package scaffolt

type Context interface {
	Generator() Generator
	Get(key string) interface{}
	Set(key string, value interface{})
	CreateFile(path string, content []byte) error
	Exec(cmd string, args ...string) error
	Move(source, target string, interpolate bool)
	Target() string
	Source() string
	Locals() map[string]interface{}
}

type File interface {
	Run(ctx Context, dryrun bool) error
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
	Name() string
}

type Generator interface {
	Description() GeneratorDescription
	Init() error
	Run(path string) error
	Root() string
	Engine(IntepreterType) ScriptEngine
	AddFile(file FileDescription) error
}

type ScriptEngine interface {
	Init(g Generator) error
	Run(path string, ctx Context) error
	Close()
}
