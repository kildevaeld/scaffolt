package scaffolt

type Question interface {
}

type Questions struct {
	desc map[string]QuestionDescription
}

func (self *Questions) Run(ctx Context) error {

	return nil
}

func (self *Questions) Init() error {
	return nil
}

func NewQuestions(desc map[string]QuestionDescription) *Questions {
	return &Questions{desc}
}
