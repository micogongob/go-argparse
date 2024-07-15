package parse

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
