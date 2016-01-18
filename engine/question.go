package engine

import "github.com/kildevaeld/scaffolt"
import "github.com/kildevaeld/go-widgets"

type Questions struct {
	desc      map[string]scaffolt.QuestionDescription
	questions map[string]widgets.Widget
}

func (self *Questions) Run(ctx scaffolt.Context) error {
	for n, q := range self.questions {
		val := q.Run()
		ctx.Set(n, val)
	}
	return nil
}

func (self *Questions) Init() error {
	for n, q := range self.desc {
		var widget widgets.Widget
		switch q.Type {
		case "input", "Input":
			widget = &widgets.Input{
				Message: n,
				Value:   q.Default,
			}
		case "list", "List":
			widget = &widgets.List{
				Message: n,
				Value:   q.Default,
				Choices: q.Choices,
			}
		case "checkbox", "Checkbox":
			widget = &widgets.Checkbox{
				Message: n,
				Value:   []string{q.Default},
				Choices: q.Choices,
			}
		case "confirm", "Confirm":
			widget = &widgets.Confirm{
				Message: n,
			}
		}

		if widget != nil {
			self.questions[n] = widget
		}
	}
	return nil
}

func NewQuestions(desc map[string]scaffolt.QuestionDescription) *Questions {
	return &Questions{desc, make(map[string]widgets.Widget)}
}
