package parse

import (
	"fmt"
)

func NewApp(description string, commands []Command) App {
	helpCommand := Command{
		Code: "help",
		Triggers: []string{"help","--help","-h"},
		Description: "Show Help",
	}

	commands = append(commands, helpCommand)

	return App{
		Description: description,
		Commands: commands,
		helpCommand: helpCommand,
	}
}

func (app *App) Parse() *Command {
	if argsEmpty() {
		app.showHelp()
	} else if argsStartsWith(app.helpCommand) {
		app.showHelp()
		return &app.helpCommand
	} else {
		isCommandMatched := false

		for _, command := range(app.Commands) {
			if argsStartsWith(command) {
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
		for _, command := range(app.Commands) {
			fmt.Printf("  %v - %v\n", command.Code, command.Description)
		}
	} else {
		fmt.Println("No commands configured")
	}
}
