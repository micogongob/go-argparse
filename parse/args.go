package parse

import (
	"os"
)

func arguments() []string {
	if len(os.Args) <= 0 {
		return []string{}
	}
	return os.Args[1:len(os.Args)]
}

func argMatchesCommand(args []string, index int, command Command) bool {
	if args[index] == command.Code {
		return true
	}

	for _, trigger := range(command.Triggers) {
		if args[index] == trigger {
			return true
		}
	}

	return false
}

func argMatchesSubCommand(args []string, index int, command SubCommand) bool {
	if args[index] == command.Code {
		return true
	}

	for _, trigger := range(command.Triggers) {
		if args[index] == trigger {
			return true
		}
	}

	return false
}
