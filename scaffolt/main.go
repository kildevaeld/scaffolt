package main

import (
	"fmt"
	"log"

	"github.com/kildevaeld/scaffolt"
)

func main() {
	gen, err := scaffolt.LoadGeneratorFromPath("example")
	if err != nil {
		log.Fatal(err)
	}

	if gen.Init(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", gen.Run("./output"))
}
