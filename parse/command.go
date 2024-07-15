package parse

import (
	"fmt"
	"strings"

	"github.com/micogongob/go-argparse/internal"
)

func (command *Command) Parse() *ParseOutput {
	args := arguments()

	if len(args) <= 1 {
		command.showHelp()
	} else if argMatches(args, 1, command.helpCommand.Code, command.helpCommand.Triggers) {
		command.showHelp()
		return &ParseOutput{
			Code: command.Code,
			ArgumentValues: map[string]string{},
		}
	} else {
		isCommandMatched := false

		for _, subCommand := range(command.SubCommands) {
			if argMatches(args, 1, subCommand.Code, subCommand.Triggers) {
				return &ParseOutput{
					Code: subCommand.Code,
					ArgumentValues: subCommand.toParsedMap(),
				}
			}
		}

		if !isCommandMatched {
			command.showHelp()
		}
	}

	return nil
}

func (command *Command) showHelp() {
	fmt.Println(command.Description)
	if len(command.SubCommands) > 0 {
		fmt.Printf("\nSub-Commands:\n")
		// TODO format output
		for _, subCommand := range(command.SubCommands) {
			subCommand.showHelp()
		}
	} else {
		fmt.Println("No sub-commands configured")
	}
}

func (subCommand *SubCommand) showHelp() {
	fmt.Printf("    %v - %v\n", subCommand.Code, subCommand.Description)
	for _, parameter := range(subCommand.Parameters) {
		alternativeTriggers := ""
		if len(parameter.Triggers) > 0 {
			alternativeTriggers = fmt.Sprintf(" (alt. %v)", strings.Join(parameter.Triggers, ","))
		}
		fmt.Printf("        --%v - %v%v\n", parameter.Code, parameter.Description, alternativeTriggers)
	}
}

func (subCommand *SubCommand) toParsedMap() map[string]string {
	values := map[string]string{}

	args := arguments()

	if len(args) <= 2 {
		subCommand.showHelp()
		internal.Fail("No parameter values was provided")
	}

	unknownParameterValues := []string{}
	for i := 2; i < len(args); i++ {
		matchedArg := false
		for _, parameter := range(subCommand.Parameters) {
			if argMatches(args, i, "--" + parameter.Code, parameter.Triggers) {
				matchedArg = true
			}
		}
		if !matchedArg {
			unknownParameterValues = append(unknownParameterValues, args[i])
		}
	}
	if len(unknownParameterValues) > 0 {
		internal.Fail(fmt.Sprintf("Unknown parameters was provided: %v", strings.Join(unknownParameterValues, ",")))
	}

	for _, parameter := range(subCommand.Parameters) {
		if !isValidParameter(args, parameter) {
			subCommand.showHelp()
			internal.Fail(fmt.Sprintf("--%v was not provided", parameter.Code))
		}
	}

	return values
}