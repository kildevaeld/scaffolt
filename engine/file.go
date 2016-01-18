package engine

import (
	"text/template"

	"github.com/kildevaeld/scaffolt"
)

type file struct {
	desc     scaffolt.FileDescription
	template *template.Template
}

func (self *file) Run(ctx scaffolt.Context) error {
	return nil
}

func (self file) Init(g scaffolt.Generator) error {

	return nil
}

func NewFile(desc scaffolt.FileDescription) scaffolt.File {
	return &file{
		desc: desc,
	}
}
