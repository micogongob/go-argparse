package parse

import (
	"fmt"
	"os"
)

func (app *App) Parse() {
	out, err := app.parseStrings(tail(os.Args))

	if err != nil {
		os.Exit(1)
	}

	if out.helpMessage != "" {
		fmt.Println(out.helpMessage)
	}
}

func (app *App) parseStrings(args []string) (parseOutput, error) {
	if len(args) > 0 {
		for _, command := range app.commands {
			if commandMatchesArg(command.code, command.aliases, args[0]) {
				if command.code == HelpCommand.code {
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
		for _, childCommand := range command.children {
			if commandMatchesArg(childCommand.code, childCommand.aliases, args[0]) {
				if childCommand.code == HelpChildCommand.code {
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
		// NOTE: outlier since parameter does not have help
		if commandMatchesArg(HelpCommand.code, HelpCommand.aliases, args[0]) {
			return parseOutput{
				helpMessage: helpToString(command.Help()),
			}, nil
		}
		for _, parameter := range command.parameters {
			if commandMatchesArg(parameter.code, []string{}, args[0]) {
				// TODO handle callback or callback
			}
		}
	}

	requiredParameters := command.requiredParameters()
	if len(requiredParameters) > 0 {
		return parseOutput{}, fmt.Errorf("missing required parameter/s: \"%v\" was not provided", toValidationMsgFormat(requiredParameters))
	}

	// TODO invoke child command function, maybe add validation in bootstrap?
	return parseOutput{}, nil
}
