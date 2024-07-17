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

func (c *Command) matches(argValue string) bool {
	if argValue == c.code {
		return true
	}

	for _, alias := range c.aliases {
		if argValue == alias {
			return true
		}
	}

	return false
}
