package main

import "C"
import (
	"sync"

	"github.com/kildevaeld/scaffolt"
)

func main() {}

//export GenD
type GenD int

type State struct {
	id         int
	generators map[*int]scaffolt.Generator
	lock       sync.Mutex
}

func (self *State) load(path string) int {

	/*gen, err := parser.LoadGeneratorFromPath(path)

	if err != nil {
		return -1
	}*/

	self.lock.Lock()
	self.id++
	id := self.id
	self.lock.Unlock()

	//self.generators[id] = gen

	return id
}

var state State

func init() {
	state = State{}

	state.generators = make(map[*int]scaffolt.Generator)
}

//export LoadGeneratorFromPath
func LoadGeneratorFromPath(path string) int {
	id := state.load(path)
	return id
}

//export Run
func Run(id int) error {
	return nil
}

//export DestroyGenerator
func DestroyGenerator(gen *GenD) {

}
