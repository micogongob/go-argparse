package parse

import (
	"fmt"
	"strings"
)

var APP = app{}

func init() {

}

func Setup(code, description string) {
	APP.code = code
	APP.description = description
}

func AddCommand(command Command) {
	APP.commands = append(APP.commands, command)
}

func Parse() (string, error) {
	return ParseStrings(arguments())
}

func ParseStrings(args []string) (string, error) {
	if len(args) > 0 {
		return "", nil
	} else {
		return helpToString(APP.Help()), nil
	}
}

func helpToString(help helpInfo) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("  %v\n\n", help.description))
	// TODO format
	sb.WriteString(fmt.Sprintf("  usage: %v %v\n\n", help.code, help.usageSuffix))

	sb.WriteString(fmt.Sprintf("  %v:\n", help.childrenName))
	for _, child := range help.children {
		sb.WriteString(fmt.Sprintf("    %v - %v\n", child.code, child.description))
	}

	return sb.String()
}
