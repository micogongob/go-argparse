package parse

func NewCommand(code, description string) *Command {
	return &Command{
		code:        code,
		description: description,
		aliases:     []string{}, // TODO don't support aliases outside for now
	}
}

func (command *Command) AddChildCommand(code, description string, parameters ...Parameter) {
	childCommand := &ChildCommand{
		code:        code,
		description: description,
		parameters:  parameters,
	}
	command.children = append(command.children, *childCommand)
}

func NewCommandParameter(code, description string, isOptional, isFlag bool) Parameter {
	return Parameter{
		code:        code,
		description: description,
		isOptional:  isOptional,
		isFlag:      isFlag,
	}
}

func NewApp(code, description string, commands ...*Command) App {
	appCommands := []Command{}

	for _, command := range commands {
		command.children = append(command.children, HelpChildCommand)
		appCommands = append(appCommands, *command)
	}

	return App{
		code:        code,
		description: description,
		commands:    append(appCommands, HelpCommand),
	}
}
