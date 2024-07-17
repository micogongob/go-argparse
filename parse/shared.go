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
