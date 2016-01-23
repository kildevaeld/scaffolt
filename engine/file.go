package engine

import (
	"path/filepath"

	"github.com/Sirupsen/logrus"
	"github.com/kildevaeld/scaffolt"
)

type file struct {
	desc   scaffolt.FileDescription
	source string
	l      *logrus.Entry
}

func (self *file) Run(ctx scaffolt.Context, dryrun bool) (err error) {

	source := filepath.Join(ctx.Source(), InterpolateOrDefault(randomString(20), self.desc.Source, ctx))
	target := filepath.Join(ctx.Target(), InterpolateOrDefault(randomString(20), self.desc.Target, ctx))

	hasContent := self.desc.Content != ""
	self.l.Printf("Creating file: %s", target)
	if dryrun {
		return nil
	}
	if self.desc.Interpolate {
		var str string
		if hasContent {
			str, err = Interpolate(randomString(20)+"name", self.desc.Content, ctx)
		} else {
			str, err = InterpolateFile(source, ctx)
		}

		if err == nil {
			err = createFileBytes(target, []byte(str), false)
		}

	} else {
		if hasContent {
			err = createFileBytes(target, []byte(self.desc.Content), false)
		} else {
			err = createFile(target, source, false)
		}
	}

	return err
}

func (self *file) Init(g scaffolt.Generator) error {
	/*if self.desc.Content != "" {
		return nil
	}

	fp := filepath.Join(g.Root(), self.desc.Source)

	if !IsFile(fp) {
		return fmt.Errorf("[file] Template does not exits: %s", fp)
	}*/

	return nil
}

func NewFile(desc scaffolt.FileDescription, l *logrus.Entry) scaffolt.File {
	return &file{
		desc: desc,
		l:    l,
	}
}
