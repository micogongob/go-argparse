package parse

type App struct {
	Description string
	Commands []Command
	helpCommand Command
}

type Command struct {
	Code string
	Triggers []string
	Description string
	SubCommands []SubCommand
	helpCommand SubCommand
}

type SubCommand struct {
	Code string
	Triggers []string
	Description string
	Parameters []Parameter
}

type Parameter struct {
	Code string
	Triggers []string
	Description string
	Required bool
	IsFlag bool
}

type ParseOutput struct {
	Code string
	ArgumentValues map[string]string
}
