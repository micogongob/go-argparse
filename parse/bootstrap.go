package parse

func NewCommand(code, description string) Command {
	return Command{
		code: code,
		description: description,
		aliases: []string{}, // TODO don't support aliases outside for now
	}
}

func (command *Command) AddChildrenCommand(code, description string) {
	// TODO might be different struct
	childCommand := Command{
		code: code,
		description: description,
		aliases: []string{}, // TODO don't support aliases outside for now
		children: []Command{}, // TODO might be parameters instead of commands
	}
	command.children = append(command.children, childCommand)
}

func NewApp(code, description string, commands ...Command) App {
	for _, command := range commands {
		// TODO fix
		command.children = append(command.children, HelpChildCommand)
	}
	return App{
		code: code,
		description: description,
		commands: append(commands, HelpCommand),
	}
}
