package parse

import (
	"fmt"
	"strings"
)

var helpCommand = &Command{
	Code:        "help",
	Description: "Show help",
	aliases:     []string{"--help", "-h"},
}

var helpChildCommand = &ChildCommand{
	Code:        "help",
	Description: "Show help",
	aliases:     []string{"--help", "-h"},
}

func helpToString(help helpInfo) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("%v\n\n", help.description))
	if len(help.children) > 0 {
		if help.usageSuffix != "" {
			sb.WriteString(fmt.Sprintf("  usage: %v %v\n\n", help.code, help.usageSuffix))
		}
	} else {
		sb.WriteString(fmt.Sprintf("  usage: %v\n", help.code))
	}

	if len(help.children) > 0 && help.childrenName != "" {
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
	children := make([]helpInfo, len(hp.Commands))

	for k, v := range hp.Commands {
		children[k] = v.Help()
	}

	return helpInfo{
		code:         hp.Code,
		description:  hp.Description,
		usageSuffix:  "[command] [subcommand] [...Parameters]",
		childrenName: "commands",
		children:     children,
	}
}

func (hp *Command) Help() helpInfo {
	var description string
	if len(hp.aliases) <= 0 {
		description = hp.Description
	} else {
		description = fmt.Sprintf("%v. Alternatives: %v", hp.Description, strings.Join(hp.aliases, ", "))
	}

	children := make([]helpInfo, len(hp.Children))
	for k, v := range hp.Children {
		children[k] = v.Help()
	}

	return helpInfo{
		code:         hp.Code,
		description:  description,
		usageSuffix:  "[subcommand] [...Parameters]",
		childrenName: "subcommands",
		children:     children,
	}
}

func (hp *ChildCommand) Help() helpInfo {
	var description string
	if len(hp.aliases) <= 0 {
		description = hp.Description
	} else {
		description = fmt.Sprintf("%v. Alternatives: %v", hp.Description, strings.Join(hp.aliases, ", "))
	}

	children := make([]helpInfo, len(hp.Parameters))
	for k, v := range hp.Parameters {
		children[k] = v.Help()
	}

	return helpInfo{
		code:         hp.Code,
		description:  description,
		usageSuffix:  "[...Parameters]",
		childrenName: "parameters",
		children:     children,
	}
}

func (hp *Parameter) Help() helpInfo {
	var description strings.Builder

	if hp.Flag {
		description.WriteString(fmt.Sprintf("%v. Flag", hp.Description))
	} else {
		description.WriteString(hp.Description)
	}

	if hp.Optional {
		description.WriteString(" (optional)")
	} else {
		description.WriteString(" (required)")
	}

	return helpInfo{
		code:        fmt.Sprintf("--%v", hp.Code),
		description: description.String(),
	}
}
