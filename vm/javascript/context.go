package javascript

import (
	"encoding/json"
	"log"

	"github.com/kildevaeld/dict"
	"github.com/kildevaeld/motto"
	"github.com/kildevaeld/scaffolt"
	"github.com/robertkrimen/otto"
)

type context struct {
	ctx scaffolt.Context
	vm  *motto.Motto
}

func (self *context) Source(call otto.FunctionCall) otto.Value {
	val, err := self.vm.ToValue(self.ctx.Source())
	if err != nil {
		return val
	}
	return otto.Value{}
}

func (self *context) Target(call otto.FunctionCall) otto.Value {
	val, err := self.vm.ToValue(self.ctx.Target())
	if err != nil {
		return val
	}
	return otto.Value{}
}

func (self *context) Set(call otto.FunctionCall) otto.Value {

	if len(call.ArgumentList) != 2 {
		return otto.UndefinedValue()
	}

	key, _ := call.Argument(0).ToString()
	val, err := call.Argument(1).Export()
	if err != nil {
		return otto.UndefinedValue()
	}

	self.ctx.Set(key, val)

	return otto.UndefinedValue()

}

func toValue(vm *motto.Motto, value interface{}) otto.Value {
	var err error
	var val otto.Value
	switch n := value.(type) {
	case dict.Map, map[string]interface{}:
		bs, e := json.Marshal(n)
		if e != nil {
			err = e
			return otto.UndefinedValue()
		}
		object, oe := vm.Object("(" + string(bs) + ")")
		if oe != nil {
			err = oe
		} else {
			val = object.Value()
		}

	default:
		val, err = vm.ToValue(n)
	}

	if err != nil {
		log.Printf("Could not convert %v", err)
		return otto.UndefinedValue()
	}

	return val
}

func (self *context) Get(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) != 1 {
		return otto.UndefinedValue()
	}

	key, _ := call.Argument(0).ToString()
	val := self.ctx.Get(key)

	if val != nil {
		return toValue(self.vm, val)
	}

	return otto.UndefinedValue()
}

func (self *context) CreateFile(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) != 2 {
		return otto.UndefinedValue()
	}

	path, _ := call.Argument(0).ToString()
	val, _ := call.Argument(1).ToString()
	self.ctx.CreateFile(path, []byte(val))

	return otto.UndefinedValue()
}

func (self *context) Locals() map[string]interface{} {
	return self.ctx.Locals()
}

func (self *context) Move(call otto.FunctionCall) otto.Value {

	if len(call.ArgumentList) < 2 {
		return otto.UndefinedValue()
	}

	var source, target string
	var interpolate bool
	var err error

	if source, err = call.Argument(0).ToString(); err != nil {
		return otto.UndefinedValue()
	}

	if target, err = call.Argument(1).ToString(); err != nil {
		return otto.UndefinedValue()
	}

	if len(call.ArgumentList) == 3 {
		interpolate, err = call.Argument(2).ToBoolean()
	}
	self.ctx.Move(source, target, interpolate)
	return otto.UndefinedValue()
}

func (self *context) Exec(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) == 0 {
		return otto.UndefinedValue()
	}

	var vals []string
	for _, arg := range call.ArgumentList {
		v, e := arg.ToString()
		if e != nil {
			return otto.UndefinedValue()
		}
		vals = append(vals, v)
	}

	self.ctx.Exec(vals[0], vals[1:]...)

	return otto.UndefinedValue()

}
