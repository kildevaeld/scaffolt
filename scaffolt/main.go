package main

import (
	"fmt"
	"log"

	"github.com/kildevaeld/scaffolt/parser"
)

func main() {
	gen, err := parser.LoadGeneratorFromPath("example")
	if err != nil {
		log.Fatal(err)
	}

	if gen.Init(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", gen.Run("./output"))
}
