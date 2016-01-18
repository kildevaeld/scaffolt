package engine

import (
	"bytes"
	"errors"
	"io/ioutil"
	"path/filepath"
	"text/template"

	"github.com/kildevaeld/scaffolt"
)

type file struct {
	desc     scaffolt.FileDescription
	template *template.Template
}

func (self *file) Run(ctx scaffolt.Context) error {
	var buf bytes.Buffer
	err := self.template.Execute(&buf, ctx.Locals())
	if err != nil {
		return err
	}

	fp := filepath.Join(ctx.Target(), self.desc.Target)

	err = ctx.CreateFile(fp, buf.Bytes())

	return err
}

func (self file) Init(g scaffolt.Generator) error {

	fp := filepath.Join(g.Root(), self.desc.Source)

	if !IsFile(fp) {
		return errors.New("not exists")
	}

	if !self.desc.Interpolate {
		return nil
	}
	bs, err := ioutil.ReadFile(fp)
	if err != nil {
		return err
	}
	base := filepath.Base(self.desc.Source)
	self.template, err = template.New(base).Parse(string(bs))
	if err != nil {
		return err
	}

	return nil
}

func NewFile(desc scaffolt.FileDescription) scaffolt.File {
	return &file{
		desc: desc,
	}
}
