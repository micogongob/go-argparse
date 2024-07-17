package parse

var App = app{}

var HelpCommand = Command{
	Code:        "help",
	Description: "Show help",
	Aliases:     []string{"--help", "-h"},
}

func init() {
	App.commands = append(App.commands, HelpCommand)
}

func Setup(code, description string) {
	App.code = code
	App.description = description
}

func AddCommand(command Command) {
	App.commands = append(App.commands, command)
}

func Parse() (ParseOutput, error) {
	return ParseStrings(arguments())
}

func ParseStrings(args []string) (ParseOutput, error) {
	if len(args) > 0 {
		for _, command := range App.commands {
			if command.matches(args[0]) {
				if command.Code == HelpCommand.Code {
					return ParseOutput{
						HelpMessage: helpToString(App.Help()),
					}, nil
				} else {
					// TODO handle callback
				}
			}
		}
	}

	return ParseOutput{
		HelpMessage: helpToString(App.Help()),
	}, nil
}

func (c *Command) matches(argValue string) bool {
	if argValue == c.Code {
		return true
	}

	for _, alias := range c.Aliases {
		if argValue == alias {
			return true
		}
	}

	return false
}
