package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/ghodss/yaml"
	"github.com/kildevaeld/scaffolt"
	"github.com/kildevaeld/scaffolt/engine"
)

const (
	GeneratorDescriptionFile = "scaffolt"
)

func LoadGeneratorFromPath(path string) (scaffolt.Generator, error) {
	var e error
	path, e = filepath.Abs(path)
	if e != nil {
		return nil, nil
	}
	if !engine.IsDir(path) {
		return nil, fmt.Errorf("Path does not exists or isn't a directory: %s", path)
	}

	descFile, err := getDescriptionFile(path)

	if err != nil {
		return nil, err
	}

	gen := engine.NewGenerator(path, *descFile)

	if err = gen.Init(); err != nil {
		return nil, err
	}

	return gen, nil
}

func getDescriptionFile(path string) (*scaffolt.GeneratorDescription, error) {
	fullPath := filepath.Join(path, GeneratorDescriptionFile)
	var m scaffolt.GeneratorDescription
	var t string
	if engine.IsFile(fullPath + ".json") {
		t = "json"
	} else if engine.IsFile(fullPath + ".yaml") {
		t = "yaml"
	} else if engine.IsFile(fullPath + "yml") {
		t = "yml"
	} else {
		return nil, fmt.Errorf("Could not find a generator description in path: %s", path)
	}

	bs, err := ioutil.ReadFile(fullPath + "." + t)
	if err != nil {
		return nil, err
	}

	switch t {
	case "json":
		err = json.Unmarshal(bs, &m)
	case "yaml", "yml":
		err = yaml.Unmarshal(bs, &m)
	default:
		err = errors.New("ERROR")
	}

	if err != nil {
		return nil, err
	}
	return &m, err
}
