package parse

type App struct {
	code        string
	description string
	commands    []Command
}

type Command struct {
	code        string
	description string
	aliases     []string
	children    []ChildCommand
}

type ChildCommand struct {
	code        string
	description string
	aliases     []string
	parameters  []Parameter
}

type Parameter struct {
	code        string
	description string
	isOptional  bool
	isFlag      bool
}

type helpInfo struct {
	code         string
	description  string
	usageSuffix  string
	childrenName string
	children     []helpInfo
}

type parseOutput struct {
	helpMessage string
}

type NewCommandInput struct {
	Code        string
	Description string
}

type AddChildCommandInput struct {
	Code        string
	Description string
	Parameters  []Parameter
}

type NewCommandParameterInput struct {
	Code        string
	Description string
	IsOptional  bool
	IsFlag      bool
}

type NewAppInput struct {
	Code        string
	Description string
	Commands    []*Command
}
