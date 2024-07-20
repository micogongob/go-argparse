package parse

import "fmt"

const (
	CODE_WHITELIST_REGEX_PATTERN = "^[A-Za-z0-9_-]*$"
	CODE_MAX_CHAR_LENGTH         = 15
	LISTS_MAX_SIZE               = 100
)

func (app *App) validate() error {
	if len(app.Commands) > LISTS_MAX_SIZE {
		return fmt.Errorf("invalid command setup: commands added exceeds max size of %d", LISTS_MAX_SIZE)
	}

	commandCodes := map[string]bool{}
	for _, command := range app.Commands {
		if err := validateCommand(command); err != nil {
			return err
		}
		if _, exists := commandCodes[command.Code]; exists {
			return fmt.Errorf("invalid command setup: \"%v\" is provided more than once", command.Code)
		} else {
			commandCodes[command.Code] = true
		}

		childCommandCodes := map[string]bool{}
		for _, childCommand := range command.Children {
			if err := validateChildCommand(command, childCommand); err != nil {
				return err
			}
			if _, exists := childCommandCodes[childCommand.Code]; exists {
				return fmt.Errorf("invalid child command setup: \"%v.%v\" is provided more than once", command.Code, childCommand.Code)
			} else {
				childCommandCodes[childCommand.Code] = true
			}

			parameterCodes := map[string]bool{}
			for _, parameter := range childCommand.Parameters {
				if err := validateParameter(command, childCommand, parameter); err != nil {
					return err
				}
				if _, exists := parameterCodes[parameter.Code]; exists {
					return fmt.Errorf("invalid parameter setup: \"%v.%v.%v\" is provided more than once", command.Code, childCommand.Code, parameter.Code)
				} else {
					parameterCodes[parameter.Code] = true
				}
			}
		}

		command.Children = append(command.Children, helpChildCommand)
	}

	app.Commands = append(app.Commands, helpCommand)
	return nil
}

func validateCommand(command *Command) error {
	if command.Code == "" {
		return fmt.Errorf("invalid command setup: Code is not provided")
	}
	if len(command.Code) > CODE_MAX_CHAR_LENGTH {
		return fmt.Errorf("invalid command setup: \"%v...\" exceeds max size of %d", command.Code[:CODE_MAX_CHAR_LENGTH], CODE_MAX_CHAR_LENGTH)
	}
	if !matchesRegex(command.Code, CODE_WHITELIST_REGEX_PATTERN) {
		return fmt.Errorf("invalid command setup: \"%v\" has invalid characters [A-Za-z0-9_-]", command.Code)
	}
	if len(command.Children) <= 0 {
		return fmt.Errorf("invalid command setup: \"%v\" should have atleast one child command", command.Code)
	}
	if len(command.Children) > LISTS_MAX_SIZE {
		return fmt.Errorf("invalid child command setup: \"%v.childCommands\" added exceeds max size of %d", command.Code, LISTS_MAX_SIZE)
	}
	return nil
}

func validateChildCommand(command *Command, childCommand *ChildCommand) error {
	if childCommand.Code == "" {
		return fmt.Errorf("invalid child command setup: \"%v.childCommands[*].Code\" is not provided", command.Code)
	}
	if len(childCommand.Code) > CODE_MAX_CHAR_LENGTH {
		return fmt.Errorf("invalid child command setup: \"%v.%v...\" exceeds max size of %d", command.Code, childCommand.Code[:CODE_MAX_CHAR_LENGTH], CODE_MAX_CHAR_LENGTH)
	}
	if !matchesRegex(childCommand.Code, CODE_WHITELIST_REGEX_PATTERN) {
		return fmt.Errorf("invalid child command setup: \"%v.%v\" has invalid characters [A-Za-z0-9_-]", command.Code, childCommand.Code)
	}
	if len(childCommand.Parameters) > LISTS_MAX_SIZE {
		return fmt.Errorf("invalid parameter setup: \"%v.%v.Parameters\" added exceeds max size of %d", command.Code, childCommand.Code, LISTS_MAX_SIZE)
	}
	return nil
}

func validateParameter(command *Command, childCommand *ChildCommand, parameter *Parameter) error {
	if parameter.Code == "" {
		return fmt.Errorf("invalid parameter setup: \"%v.%v.Parameters[*].Code\" is not provided", command.Code, childCommand.Code)
	}
	if len(parameter.Code) > CODE_MAX_CHAR_LENGTH {
		return fmt.Errorf("invalid parameter setup: \"%v.%v.%v...\" exceeds max size of %d", command.Code, childCommand.Code, parameter.Code[:CODE_MAX_CHAR_LENGTH], CODE_MAX_CHAR_LENGTH)
	}
	if !matchesRegex(parameter.Code, CODE_WHITELIST_REGEX_PATTERN) {
		return fmt.Errorf("invalid parameter setup: \"%v.%v.%v\" has invalid characters [A-Za-z0-9_-]", command.Code, childCommand.Code, parameter.Code)
	}
	if !parameter.IsOptional && parameter.IsBoolean {
		return fmt.Errorf("invalid parameter setup: \"%v.%v.%v\" cannot be required and a boolean at the same time", command.Code, childCommand.Code, parameter.Code)
	}
	return nil
}
