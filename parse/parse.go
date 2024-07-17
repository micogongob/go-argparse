package parse

import (
	"fmt"
	"os"
)

func (app *App) Parse() {
	out, err := app.parseStrings(arguments())

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
			if command.matches(args[0]) {
				if command.code == HelpCommand.code {
					return parseOutput{
						helpMessage: helpToString(app.Help()),
					}, nil
				} else {
					return command.parseStrings(args)
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
		for _, command := range command.children {
			if command.matches(args[0]) {
				if command.code == HelpChildCommand.code {
					return parseOutput{
						helpMessage: helpToString(command.Help()),
					}, nil
				} else {
					// TODO handle callback or callback
				}
			}
		}
	}

	return parseOutput{
		helpMessage: helpToString(command.Help()),
	}, nil
}
