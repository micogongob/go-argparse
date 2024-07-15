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

func argMatches(args []string, index int, pattern string, triggers []string) bool {
	if args[index] == pattern {
		return true
	}

	for _, trigger := range triggers {
		if args[index] == trigger {
			return true
		}
	}

	return false
}

func isValidParameter(args []string, parameter Parameter) bool {
	if parameter.Optional {
		return true
	}

	for i := 2; i < len(args); i++ {
		if "--"+args[i] == parameter.Code {
			return true
		}
		for _, trigger := range parameter.Triggers {
			if args[i] == trigger {
				return true
			}
		}
	}

	return false
}
