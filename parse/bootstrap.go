package parse

func NewCommand(input NewCommandInput) *Command {
	return &Command{
		code:        input.Code,
		description: input.Description,
		aliases:     []string{}, // TODO don't support aliases outside for now
	}
}

func (command *Command) AddChildCommand(input AddChildCommandInput) {
	childCommand := &ChildCommand{
		code:        input.Code,
		description: input.Description,
		parameters:  input.Parameters,
	}
	command.children = append(command.children, *childCommand)
}

func NewCommandParameter(input NewCommandParameterInput) Parameter {
	return Parameter{
		code: input.Code,
		description: input.Description,
		isOptional: input.IsOptional,
		isFlag: input.IsFlag,
	}
}

func NewApp(input NewAppInput) (App, error) {
	appCommands := []Command{}

	for _, command := range input.Commands {
		command.children = append(command.children, HelpChildCommand)
		appCommands = append(appCommands, *command)
	}

	return App{
		code:        input.Code,
		description: input.Description,
		commands:    append(appCommands, HelpCommand),
	}, nil
}
