package parse

import (
	"fmt"
	"strings"
)

var HelpCommand = Command{
	code:        "help",
	description: "Show help",
	aliases:     []string{"--help", "-h"},
}

var HelpChildCommand = ChildCommand{
	code:        "help",
	description: "Show help",
	aliases:     []string{"--help", "-h"},
}

func helpToString(help helpInfo) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("%v\n\n", help.description))
	if help.usageSuffix != "" {
		sb.WriteString(fmt.Sprintf("  usage: %v %v\n\n", help.code, help.usageSuffix))
	}

	if help.childrenName != "" {
		sb.WriteString(fmt.Sprintf("  %v:\n", help.childrenName))
	}

	codePadRightLength := greatestLengthCode(help.children)
	for _, child := range help.children {
		sb.WriteString(fmt.Sprintf("    %v -> %v\n", padRight(child.code, codePadRightLength), child.description))
	}

	return sb.String()
}

func greatestLengthCode(listOfHelp []helpInfo) int {
	length := 0
	for _, h := range listOfHelp {
		currentLength := len(h.code)
		if currentLength > length {
			length = currentLength
		}
	}
	return length
}

func padRight(source string, padLength int) string {
	if len(source) >= padLength {
		return source
	}

	var sb strings.Builder
	sb.WriteString(source)

	for i := len(source); i < padLength; i++ {
		sb.WriteString(" ")
	}
	return sb.String()
}

func (hp *App) Help() helpInfo {
	children := make([]helpInfo, len(hp.commands))

	for k, v := range hp.commands {
		children[k] = v.Help()
	}

	return helpInfo{
		code:         hp.code,
		description:  hp.description,
		usageSuffix:  "[command] [subcommand] [...parameters]",
		childrenName: "commands",
		children:     children,
	}
}

func (hp *Command) Help() helpInfo {
	var description string
	if len(hp.aliases) <= 0 {
		description = hp.description
	} else {
		description = fmt.Sprintf("%v. Alternatives: %v", hp.description, strings.Join(hp.aliases, ", "))
	}

	children := make([]helpInfo, len(hp.children))
	for k, v := range hp.children {
		children[k] = v.Help()
	}

	return helpInfo{
		code:         hp.code,
		description:  description,
		usageSuffix:  "[subcommand] [...parameters]",
		childrenName: "subcommands",
		children:     children,
	}
}

func (hp *ChildCommand) Help() helpInfo {
	var description string
	if len(hp.aliases) <= 0 {
		description = hp.description
	} else {
		description = fmt.Sprintf("%v. Alternatives: %v", hp.description, strings.Join(hp.aliases, ", "))
	}

	children := make([]helpInfo, len(hp.parameters))
	for k, v := range hp.parameters {
		children[k] = v.Help()
	}

	return helpInfo{
		code:         hp.code,
		description:  description,
		usageSuffix:  "[...parameters]",
		childrenName: "parameters",
		children:     children,
	}
}

func (hp *Parameter) Help() helpInfo {
	var description strings.Builder

	if hp.isFlag {
		description.WriteString(fmt.Sprintf("%v. Flag", hp.description))
	} else {
		description.WriteString(hp.description)
	}

	if hp.isOptional {
		description.WriteString(" (optional)")
	} else {
		description.WriteString(" (required)")
	}

	return helpInfo{
		code: fmt.Sprintf("--%v", hp.code),
		description: description.String(),
	}
}
