//go:generate stringer -type=IntepreterType
package scaffolt

import (
	"errors"
	"strings"
)

type IntepreterType int

const (
	Javascript IntepreterType = iota
	Lua
)

func (i IntepreterType) MarshalJSON() ([]byte, error) {
	return []byte("\"" + i.String() + "\""), nil
}

func (i *IntepreterType) UnmarshalJSON(bs []byte) error {
	str := strings.Trim(string(bs), "\"")
	switch str {
	case "Javascript":
		*i = Javascript
	case "Lua":
		*i = Lua
	default:
		return errors.New("could not " + str)

	}
	return nil
}
