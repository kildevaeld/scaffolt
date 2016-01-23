package engine

import (
	"fmt"

	"github.com/kildevaeld/scaffolt"
)
import "github.com/kildevaeld/go-widgets"

type Questions struct {
	desc      []scaffolt.QuestionDescription
	questions map[string]widgets.Widget
}

func getWidget(n string, q scaffolt.QuestionDescription, ctx scaffolt.Context) widgets.Widget {

	def := InterpolateOrDefault(randomString(20)+"-default", q.Default, ctx)

	switch q.Type {
	case "input", "Input":
		return &widgets.Input{
			Message: n,
			Value:   def,
		}
	case "list", "List":
		return &widgets.List{
			Message: n,
			Value:   def,
			Choices: q.Choices,
		}
	case "checkbox", "Checkbox":
		return &widgets.Checkbox{
			Message: n,
			Value:   []string{def},
			Choices: q.Choices,
		}
	case "confirm", "Confirm":
		return &widgets.Confirm{
			Message: n,
			//Default: stringToBoolean(def),
		}
	}
	return nil
}

func (self *Questions) Run(ctx scaffolt.Context) error {

	for _, q := range self.desc {
		w := getWidget(q.Name, q, ctx)
		if w == nil {
			return fmt.Errorf("No widget: %s", q.Type)
		}
		ctx.Set(q.Name, w.Run())

		for _, file := range q.Files {
			ctx.Generator().AddFile(file)
		}
	}

	return nil
}

func (self *Questions) Init() error {

	return nil
}

func NewQuestions(desc []scaffolt.QuestionDescription) *Questions {
	return &Questions{desc, make(map[string]widgets.Widget)}
}
