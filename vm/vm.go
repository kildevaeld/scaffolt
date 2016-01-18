package vm

import (
	"github.com/kildevaeld/scaffolt"
	"github.com/kildevaeld/scaffolt/vm/javascript"
)

func NewScript(engine scaffolt.ScriptEngine, desc scaffolt.ScriptDescription) scaffolt.Script {
	switch desc.Type {
	case scaffolt.Javascript:
		return javascript.New(engine, desc.Path)
	}
	return nil
}
