package parse

import (
	"fmt"
	"strings"
)

func NewApp(description string, commands []Command) App {
	// for index, command := range commands {
	// 	subCommandhelpCommand := SubCommand{
	// 		Code:        "help",
	// 		Triggers:    []string{"--help", "-h"},
	// 		Description: "Show Help",
	// 	}
	// 	commands[index].helpCommand = subCommandhelpCommand
	// 	commands[index].SubCommands = append(command.SubCommands, subCommandhelpCommand)
	// }
	appHelpCommand := Command{
		Code:        "help",
		Triggers:    []string{"--help", "-h"},
		Description: "Show Help",
	}

	app := App{
		Description: description,
		Commands:    append(commands, appHelpCommand),
		HelpCommand: appHelpCommand,
	}

	appHelpCommand.OnCommand = func(paramValues map[string]string) error {
		return nil
	}

	return app
}

func (app *App) Parse() (string, error) {
	return app.ParseStrings(arguments())
}

func (app *App) ParseStrings(args []string) (string, error) {
	if len(args) > 0 {
		return "", nil
	} else {
		return helpInfoToString(app.helpInfo()), nil
		// } else if argMatches(args, 0, app.helpCommand.Code, app.helpCommand.Triggers) {
		// 	return showHelp(&app.HelpCommand)
		// } else {
		// 	isCommandMatched := false

		// 	for _, command := range app.Commands {
		// 		if argMatches(args, 0, command.Code, command.Triggers) {
		// 			return &command
		// 		}
		// 	}

		// 	if !isCommandMatched {
		// 		return showHelp(&app.HelpCommand)
		// 	}
	}
}

func helpInfoToString(help HelpInfo) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("%v\n", help.description))
	for _, subHelp := range help.subHelpInfo {
		code := subHelp.code
		if code == nil {
			sb.WriteString(fmt.Sprintf("%v\n", helpInfoToString(subHelp)))
		}
		// TODO format
		sb.WriteString(fmt.Sprintf("   %v - %v\n", code, subHelp.description))
	}
	return sb.String()
}
