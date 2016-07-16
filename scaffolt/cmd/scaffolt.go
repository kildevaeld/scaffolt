package cmd

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/kildevaeld/scaffolt/parser"
	"github.com/spf13/viper"
)

func scaffolt(template, output string) error {
	path := viper.Get("scaffolt.path").(string)
	tPath := filepath.Join(path, "templates", template)

	if _, err := os.Stat(tPath); err != nil {
		return err
	}

	gen, err := parser.LoadGeneratorFromPath(tPath)
	if err != nil {
		log.Fatal(err)
	}

	if gen.Init(); err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("%v", )
	return gen.Run(output)
}

func listTemplates() error {
	//path := viper.Get("scaffolt.path").(string)
	//tPath := filepath.Join(path, "templates")
	return nil
}

func addTemplate(template string) error {

	path := viper.Get("scaffolt.path").(string)
	tPath := filepath.Join(path, "templates")

	if stat, err := os.Stat(tPath); err != nil {
		os.MkdirAll(tPath, 0766)
	} else if stat.IsDir() {

	}

	if strings.HasPrefix(template, "file://") {
		return addTemplateFromPath(template, tPath)
	} else if strings.HasPrefix(template, "git+") {

	} else if strings.HasPrefix(template, "http") {

	}

	return nil
}

func addTemplateFromPath(path, target string) error {
	path = strings.Replace(path, "file://", "", 1)

	if stat, err := os.Stat(path); err != nil {
		return err
	} else if !stat.IsDir() {
		return errors.New("not a directory")
	}

	base := filepath.Base(path)

	return CopyDir(path, filepath.Join(target, base))

}
