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

func argsEmpty() bool {
	return len(arguments()) <= 0
}

func argsStartsWith(command Command) bool {
	args := arguments()

	if command.Triggers == nil || len(command.Triggers) <= 0 {
		return "--" + args[0] == command.Code
	}

	for _, trigger := range(command.Triggers) {
		if args[0] == trigger {
			return true
		}
	}

	return false
}
