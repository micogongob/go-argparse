package parse

import (
	"fmt"
)

func (app *App) Parse() *Command {
	args := arguments()

	if len(args) <= 0 {
		app.showHelp()
	} else if argMatches(args, 0, app.helpCommand.Code, app.helpCommand.Triggers) {
		app.showHelp()
		return &app.helpCommand
	} else {
		isCommandMatched := false

		for _, command := range app.Commands {
			if argMatches(args, 0, command.Code, command.Triggers) {
				return &command
			}
		}

		if !isCommandMatched {
			app.showHelp()
		}
	}

	return nil
}

func (app *App) showHelp() {
	fmt.Println(app.Description)
	if len(app.Commands) > 0 {
		fmt.Printf("\nCommands:\n")
		// TODO format output
		for _, command := range app.Commands {
			fmt.Printf("    %v - %v\n", command.Code, command.Description)
		}
	} else {
		fmt.Println("No commands configured")
	}
}
