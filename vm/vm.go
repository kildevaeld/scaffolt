package vm

import (
	"github.com/kildevaeld/scaffolt"
	"github.com/kildevaeld/scaffolt/vm/javascript"
)

func NewScript(desc scaffolt.ScriptDescription) scaffolt.Script {
	switch desc.Type {
	case scaffolt.Javascript:
		return javascript.New(desc.Path)
	}
	return nil
}
