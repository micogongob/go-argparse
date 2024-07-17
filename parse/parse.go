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
					// TODO handle callback
				}
			}
		}
	}

	return parseOutput{
		helpMessage: helpToString(app.Help()),
	}, nil
}

func (c *Command) matches(argValue string) bool {
	if argValue == c.code {
		return true
	}

	for _, alias := range c.aliases {
		if argValue == alias {
			return true
		}
	}

	return false
}
