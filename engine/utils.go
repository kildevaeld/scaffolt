package engine

import (
	"bytes"
	"os"
	"text/template"

	"github.com/kildevaeld/scaffolt"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

func IsDir(path string) bool {
	if stat, err := os.Stat(path); err == nil {
		return stat.IsDir()
	}
	return false
}

func IsFile(path string) bool {
	if stat, err := os.Stat(path); err == nil {
		return !stat.IsDir()
	}
	return false
}

func Interpolate(name, str string, ctx scaffolt.Context) (string, error) {
	t, err := template.New(name).Parse(str)

	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = t.Execute(&buf, ctx.Locals())
	if err != nil {
		return "", err
	}
	return string(buf.Bytes())
}
