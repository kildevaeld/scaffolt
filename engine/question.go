package engine

import "github.com/kildevaeld/scaffolt"

type Question interface {
}

type Questions struct {
	desc map[string]scaffolt.QuestionDescription
}

func (self *Questions) Run(ctx scaffolt.Context) error {

	return nil
}

func (self *Questions) Init() error {
	return nil
}

func NewQuestions(desc map[string]scaffolt.QuestionDescription) *Questions {
	return &Questions{desc}
}
