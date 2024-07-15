package parse

import (
	"fmt"
	"strings"
)

func NewApp(description string, commands []Command) App {
	for index, command := range(commands) {
		subCommandhelpCommand := SubCommand{
			Code: "help",
			Triggers: []string{"--help", "-h"},
			Description: "Show Help",
		} 
		commands[index].helpCommand = subCommandhelpCommand
		commands[index].SubCommands = append(command.SubCommands, subCommandhelpCommand)
	}

	helpCommand := Command{
		Code: "help",
		Triggers: []string{"--help", "-h"},
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
	args := arguments()

	if len(args) <= 0 {
		app.showHelp()
	} else if argMatchesCommand(args, 0, app.helpCommand) {
		app.showHelp()
		return &app.helpCommand
	} else {
		isCommandMatched := false

		for _, command := range(app.Commands) {
			if argMatchesCommand(args, 0, command) {
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
			fmt.Printf("    %v - %v\n", command.Code, command.Description)
		}
	} else {
		fmt.Println("No commands configured")
	}
}

func (command *Command) Parse() *ParseOutput {
	args := arguments()

	if len(args) <= 1 {
		command.showHelp()
	} else if argMatchesSubCommand(args, 1, command.helpCommand) {
		command.showHelp()
		return &ParseOutput{
			Code: command.Code,
			ArgumentValues: map[string]string{},
		}
	} else {
		isCommandMatched := false

		for _, subCommand := range(command.SubCommands) {
			if argMatchesSubCommand(args, 1, subCommand) {
				return &ParseOutput{
					Code: subCommand.Code,
					// TODO parse
					ArgumentValues: map[string]string{},
				}
			}
		}

		if !isCommandMatched {
			command.showHelp()
		}
	}

	return nil
}

func (command *Command) showHelp() {
	fmt.Println(command.Description)
	if len(command.SubCommands) > 0 {
		fmt.Printf("\nSub-Commands:\n")
		// TODO format output
		for _, subCommand := range(command.SubCommands) {
			fmt.Printf("    %v - %v\n", subCommand.Code, subCommand.Description)
			for _, parameter := range(subCommand.Parameters) {
				alternativeTriggers := ""
				if len(parameter.Triggers) > 0 {
					alternativeTriggers = fmt.Sprintf(" (alt. %v)", strings.Join(parameter.Triggers, ","))
				}
				fmt.Printf("        --%v - %v%v\n", parameter.Code, parameter.Description, alternativeTriggers)
			}
		}
	} else {
		fmt.Println("No sub-commands configured")
	}
}
