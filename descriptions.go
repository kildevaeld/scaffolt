package scaffolt

type GeneratorDescription struct {
	Name        string
	Description string
	Tasks       []TaskDescription
}

type TaskDescription struct {
	Name      string
	Before    ScriptDescription
	After     ScriptDescription
	Run       ScriptDescription
	Questions map[string]QuestionDescription
	Files     []FileDescription
}

type ScriptDescription struct {
	Type IntepreterType
	Path string
}

type QuestionDescription struct {
	Type    string
	Default string
	Before  ScriptDescription
	After   ScriptDescription
	Choices []string
}

type FileDescription struct {
	Source string
	Target string
	Before ScriptDescription
	After  ScriptDescription
}
