package parse

import "fmt"

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
		if len(command.children) <= 0 {
			return App{}, fmt.Errorf("invalid command setup: %v should have atleast one child command", command.code)
		}
		for _, childCommand := range command.children {
			for _, parameter := range childCommand.parameters {
				if !parameter.isOptional && parameter.isFlag {
					return App{}, fmt.Errorf("invalid parameter setup: %v cannot be required and a flag at the same time", parameter.code)
				}
			}
		}

		command.children = append(command.children, HelpChildCommand)
		appCommands = append(appCommands, *command)
	}

	return App{
		code:        input.Code,
		description: input.Description,
		commands:    append(appCommands, HelpCommand),
	}, nil
}
