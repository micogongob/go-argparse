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

var helpParameter = &Parameter{
	Code:        "help",
	Description: "Show help",
	IsOptional:  true,
	IsBoolean:   true,
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
		paddedRight := padRight(child.code, codePadRightLength)
		sb.WriteString(fmt.Sprintf("    %v -> %v\n", paddedRight, child.description))
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
	// TODO add unit tests
	var helpCode string
	if hp.Code == "" {
		helpCode = "{this}"
	} else {
		helpCode = hp.Code
	}

	var helpDescription string
	if hp.Description == "" {
		helpDescription = "Cli tool"
	} else {
		helpDescription = hp.Description
	}

	children := make([]helpInfo, len(hp.Commands))

	for k, v := range hp.Commands {
		children[k] = v.Help()
	}

	return helpInfo{
		code:         helpCode,
		description:  helpDescription,
		usageSuffix:  "[command] [subcommand] [...parameters]",
		childrenName: "commands",
		children:     children,
	}
}

func (hp *Command) Help() helpInfo {
	// TODO add unit tests
	var codeDescription string
	if hp.Description == "" {
		codeDescription = fmt.Sprintf("%v actions", hp.Code)
	} else {
		codeDescription = hp.Description
	}

	var helpDescription string
	if len(hp.aliases) <= 0 {
		helpDescription = codeDescription
	} else {
		helpDescription = fmt.Sprintf("%v. Alternatives: %v", codeDescription, strings.Join(hp.aliases, ", "))
	}

	children := make([]helpInfo, len(hp.Children))
	for k, v := range hp.Children {
		children[k] = v.Help()
	}

	return helpInfo{
		code:         hp.Code,
		description:  helpDescription,
		usageSuffix:  "[subcommand] [...parameters]",
		childrenName: "subcommands",
		children:     children,
	}
}

func (hp *ChildCommand) Help() helpInfo {
	// TODO add unit tests
	var codeDescription string
	if hp.Description == "" {
		codeDescription = fmt.Sprintf("Execute %v", hp.Code)
	} else {
		codeDescription = hp.Description
	}

	var helpDescription string
	if len(hp.aliases) <= 0 {
		helpDescription = codeDescription
	} else {
		helpDescription = fmt.Sprintf("%v. Alternatives: %v", codeDescription, strings.Join(hp.aliases, ", "))
	}

	children := make([]helpInfo, len(hp.Parameters))
	for k, v := range hp.Parameters {
		children[k] = v.Help()
	}

	return helpInfo{
		code:         hp.Code,
		description:  helpDescription,
		usageSuffix:  "[...parameters]",
		childrenName: "parameters",
		children:     children,
	}
}

func (hp *Parameter) Help() helpInfo {
	// TODO add unit tests
	var codeDescription string
	if hp.Description == "" {
		codeDescription = fmt.Sprintf("The %v", hp.Code)
	} else {
		codeDescription = hp.Description
	}

	var helpDescription strings.Builder
	if len(hp.aliases) > 0 {
		withoutParamCodeAliases := []string{}
		for _, alias := range hp.aliases {
			if alias == fmt.Sprintf("--%v", hp.Code) {
				withoutParamCodeAliases = append(withoutParamCodeAliases, hp.Code)
			} else {
				withoutParamCodeAliases = append(withoutParamCodeAliases, alias)
			}
		}
		helpDescription.WriteString(fmt.Sprintf("%v. Alternatives: %v", codeDescription, strings.Join(withoutParamCodeAliases, ", ")))
	} else {
		if hp.IsNumber {
			helpDescription.WriteString(fmt.Sprintf("%v. Number", codeDescription))
		} else if hp.IsBoolean {
			helpDescription.WriteString(fmt.Sprintf("%v. Boolean", codeDescription))
		} else {
			helpDescription.WriteString(fmt.Sprintf("%v. String", codeDescription))
		}
		if hp.IsOptional {
			helpDescription.WriteString(" (optional)")
		} else {
			helpDescription.WriteString(" (required)")
		}
	}

	return helpInfo{
		code:        fmt.Sprintf("--%v", hp.Code),
		description: helpDescription.String(),
	}
}
