package scaffolt

type GeneratorDescription struct {
	Name        string
	Description string
	Tasks       []TaskDescription
	Files       []FileDescription
}

type TaskDescription struct {
	Name      string
	Before    ScriptDescription
	After     ScriptDescription
	Run       ScriptDescription
	Questions []QuestionDescription
	Files     []FileDescription
}

type ScriptDescription struct {
	Type IntepreterType
	Path string
}

type QuestionDescription struct {
	Type    string
	Default string
	Name    string
	Message string
	Before  ScriptDescription
	After   ScriptDescription
	Choices []string
	Files   []FileDescription
}

type FileDescription struct {
	Source      string
	Target      string
	Interpolate bool
	Before      ScriptDescription
	After       ScriptDescription
	Content     string
}
