package scaffolt

type File interface {
	Run(ctx Context) error
	Init() error
}

type file struct {
}

func (self *file) Run(ctx Context) error {
	return nil
}

func (self file) Init() error {
	return nil
}

func NewFile(desc FileDescription) File {
	return &file{}
}
