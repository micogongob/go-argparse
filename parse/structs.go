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
}

type SubCommand struct {
	Code string
	Triggers []string
	Description string
	Arguments []Argument
}

type Argument struct {
	Code string
	Triggers []string
	Description string
	IsFlag bool
}

type ParseOutput struct {
	Code string
	ArgumentValues map[string]string
}
