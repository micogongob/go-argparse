package parse

import (
	"fmt"
	"os"
)

func (app *App) Parse() error {
	err := app.validate()

	out, err := app.parseStrings(tail(os.Args))

	if err != nil {
		return err
	}

	if out.helpMessage != "" {
		fmt.Println(out.helpMessage)
	}

	return nil
}

func (app *App) parseStrings(args []string) (parseOutput, error) {
	if len(args) > 0 {
		for _, command := range app.Commands {
			if commandMatchesArg(command.Code, command.aliases, args[0]) {
				if command.Code == helpCommand.Code {
					return parseOutput{
						helpMessage: helpToString(app.Help()),
					}, nil
				} else {
					return command.parseStrings(tail(args))
				}
			}
		}
	}

	return parseOutput{
		helpMessage: helpToString(app.Help()),
	}, nil
}

func (command *Command) parseStrings(args []string) (parseOutput, error) {
	if len(args) > 0 {
		for _, childCommand := range command.Children {
			if commandMatchesArg(childCommand.Code, childCommand.aliases, args[0]) {
				if childCommand.Code == helpChildCommand.Code {
					return parseOutput{
						helpMessage: helpToString(command.Help()),
					}, nil
				} else {
					return childCommand.parseStrings(tail(args))
				}
			}
		}
	}

	return parseOutput{
		helpMessage: helpToString(command.Help()),
	}, nil
}

func (command *ChildCommand) parseStrings(args []string) (parseOutput, error) {
	if len(args) > 0 {
		if len(args) == 1 && commandMatchesArg(helpParameter.Code, helpParameter.aliases, args[0]) {
			return parseOutput{
				helpMessage: helpToString(command.Help()),
			}, nil
		}
		err := command.extractParameterValues(args)
		return parseOutput{}, err
	}

	requiredParameters := command.requiredParameters()
	if len(requiredParameters) > 0 {
		return parseOutput{}, fmt.Errorf("missing required parameter/s: \"%v\" was not provided", toValidationMsgFormat(requiredParameters))
	}

	// TODO invoke child command function
	return parseOutput{}, nil
}
