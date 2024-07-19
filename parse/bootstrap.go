package parse

import "fmt"

const (
	CODE_WHITELIST_REGEX_PATTERN = "^[A-Za-z0-9_-]*$"
	CODE_MAX_CHAR_LENGTH = 15
	LISTS_MAX_SIZE       = 100
)

func NewCommand(input NewCommandInput) *Command {
	return &Command{
		code:        input.Code,
		description: input.Description,
		aliases:     []string{}, // NOTE: might not support aliasess for now aside from "help" command
	}
}

func (command *Command) AddChildCommand(input AddChildCommandInput) {
	childCommand := &ChildCommand{
		code:        input.Code,
		description: input.Description,
		parameters:  input.Parameters,
		aliases:     []string{}, // NOTE: might not support aliasess for now aside from "help" command
	}
	command.children = append(command.children, *childCommand)
}

func NewCommandParameter(input NewCommandParameterInput) Parameter {
	return Parameter{
		code:        input.Code,
		description: input.Description,
		isOptional:  input.IsOptional,
		isFlag:      input.IsFlag,
	}
}

func NewApp(input NewAppInput) (App, error) {
	if len(input.Commands) > LISTS_MAX_SIZE {
		return App{}, fmt.Errorf("invalid command setup: commands added exceeds max size of %d", LISTS_MAX_SIZE)
	}

	appCommands := []Command{}

	commandCodes := map[string]bool{}
	for _, command := range input.Commands {
		if err := validateCommand(command); err != nil {
			return App{}, err
		}
		if _, exists := commandCodes[command.code]; exists {
			return App{}, fmt.Errorf("invalid command setup: \"%v\" is provided more than once", command.code)
		} else {
			commandCodes[command.code] = true
		}

		childCommandCodes := map[string]bool{}
		for _, childCommand := range command.children {
			if err := validateChildCommand(command, childCommand); err != nil {
				return App{}, err
			}
			if _, exists := childCommandCodes[childCommand.code]; exists {
				return App{}, fmt.Errorf("invalid child command setup: \"%v.%v\" is provided more than once", command.code, childCommand.code)
			} else {
				childCommandCodes[childCommand.code] = true
			}

			parameterCodes := map[string]bool{}
			for _, parameter := range childCommand.parameters {
				if err := validateParameter(command, childCommand, parameter); err != nil {
					return App{}, err
				}
				if _, exists := parameterCodes[parameter.code]; exists {
					return App{}, fmt.Errorf("invalid parameter setup: \"%v.%v.%v\" is provided more than once", command.code, childCommand.code, parameter.code)
				} else {
					parameterCodes[parameter.code] = true
				}
			}
		}

		command.children = append(command.children, HelpChildCommand)
		appCommands = append(appCommands, *command)
	}

	return App{
		code:        input.Code,
		description: input.Description,
		commands:    append(appCommands, HelpCommand),
	}, nil
}

func validateCommand(command *Command) error {
	if command.code == "" {
		return fmt.Errorf("invalid command setup: Code is not provided")
	}
	if len(command.code) > CODE_MAX_CHAR_LENGTH {
		return fmt.Errorf("invalid command setup: \"%v...\" exceeds max size of %d", command.code[:CODE_MAX_CHAR_LENGTH], CODE_MAX_CHAR_LENGTH)
	}
	if !matchesRegex(command.code, CODE_WHITELIST_REGEX_PATTERN) {
		return fmt.Errorf("invalid command setup: \"%v\" has invalid characters [A-Za-z0-9_-]", command.code)
	}
	if len(command.children) <= 0 {
		return fmt.Errorf("invalid command setup: \"%v\" should have atleast one child command", command.code)
	}
	if len(command.children) > LISTS_MAX_SIZE {
		return fmt.Errorf("invalid child command setup: \"%v.childCommands\" added exceeds max size of %d", command.code, LISTS_MAX_SIZE)
	}
	return nil
}

func validateChildCommand(command *Command, childCommand ChildCommand) error {
	if childCommand.code == "" {
		return fmt.Errorf("invalid child command setup: \"%v.childCommands[*].Code\" is not provided", command.code)
	}
	if len(childCommand.code) > CODE_MAX_CHAR_LENGTH {
		return fmt.Errorf("invalid child command setup: \"%v.%v...\" exceeds max size of %d", command.code, childCommand.code[:CODE_MAX_CHAR_LENGTH], CODE_MAX_CHAR_LENGTH)
	}
	if !matchesRegex(childCommand.code, CODE_WHITELIST_REGEX_PATTERN) {
		return fmt.Errorf("invalid child command setup: \"%v.%v\" has invalid characters [A-Za-z0-9_-]", command.code, childCommand.code)
	}
	if len(childCommand.parameters) > LISTS_MAX_SIZE {
		return fmt.Errorf("invalid parameter setup: \"%v.%v.parameters\" added exceeds max size of %d", command.code, childCommand.code, LISTS_MAX_SIZE)
	}
	return nil
}

func validateParameter(command *Command, childCommand ChildCommand, parameter Parameter) error {
	if parameter.code == "" {
		return fmt.Errorf("invalid parameter setup: \"%v.%v.parameters[*].Code\" is not provided", command.code, childCommand.code)
	}
	if len(parameter.code) > CODE_MAX_CHAR_LENGTH {
		return fmt.Errorf("invalid parameter setup: \"%v.%v.%v...\" exceeds max size of %d", command.code, childCommand.code, parameter.code[:CODE_MAX_CHAR_LENGTH], CODE_MAX_CHAR_LENGTH)
	}
	if !matchesRegex(parameter.code, CODE_WHITELIST_REGEX_PATTERN) {
		return fmt.Errorf("invalid parameter setup: \"%v.%v.%v\" has invalid characters [A-Za-z0-9_-]", command.code, childCommand.code, parameter.code)
	}
	if !parameter.isOptional && parameter.isFlag {
		return fmt.Errorf("invalid parameter setup: \"%v.%v.%v\" cannot be required and a flag at the same time", command.code, childCommand.code, parameter.code)
	}
	return nil
}
