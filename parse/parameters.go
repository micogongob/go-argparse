package parse

import (
	"fmt"
	"strings"
)

const (
	PARAMETER_VALUE_MAX_SIZE = 1000
)

func (command *ChildCommand) requiredParameters() []Parameter {
	params := []Parameter{}
	for _, param := range command.parameters {
		if !param.isOptional {
			params = append(params, param)
		}
	}
	return params
}

func (command *ChildCommand) parseParameterValues(args []string) error {
	requiredParameters := command.requiredParameters()
	if err := validateRequiredParameters(requiredParameters, args); err != nil {
		return err
	}
	parameterValues, err := filterParameterValues(command.parameters, args)
	if err != nil {
		return err
	}
	if err := validateParameterValues(parameterValues); err != nil {
		return err
	}

	// TODO return to on trigger
	fmt.Println(parameterValues)
	return nil
}

func (parameter *Parameter) matchesArg(rawArgValue string) (bool, bool) {
	usingEqualsAssignment, values := getEqualAssigntmentValues(rawArgValue)
	if usingEqualsAssignment {
		return fmt.Sprintf("--%v", parameter.code) == values[0], usingEqualsAssignment
	} else {
		return fmt.Sprintf("--%v", parameter.code) == rawArgValue, usingEqualsAssignment
	}
}

func toValidationMsgFormat(params []Parameter) string {
	s := []string{}
	for _, v := range params {
		s = append(s, fmt.Sprintf("--%v", v.code))
	}
	return strings.Join(s, ",")
}

func validateRequiredParameters(parameters []Parameter, args []string) error {
	notProvidedRequiredParameters := []Parameter{}

	for _, param := range parameters {
		exists := false

		for _, rawArgValue := range args {
			if matches, _ := param.matchesArg(rawArgValue); matches {
				exists = true
			}
		}

		if !exists {
			notProvidedRequiredParameters = append(notProvidedRequiredParameters, param)
		}
	}
	if len(notProvidedRequiredParameters) > 0 {
		return fmt.Errorf("missing required parameter/s: \"%v\" was not provided", toValidationMsgFormat(notProvidedRequiredParameters))
	}

	return nil
}

func validateUnknownValues(parameters []Parameter, args []string) error {
	return fmt.Errorf("Unimplemented validation")
}

func filterParameterValues(parameters []Parameter, args []string) (map[string]string, error) {
	parameterValues := map[string]string{}

	for i := 0; i < len(args); i++ {
		rawArgValue := args[i]

		argFoundParamMatch := false
		for _, param := range parameters {
			matches, usingEqualsAssignment := param.matchesArg(rawArgValue)
			if !matches {
				continue
			}

			if param.isFlag {
				if usingEqualsAssignment, _ := getEqualAssigntmentValues(rawArgValue); usingEqualsAssignment {
					return map[string]string{}, fmt.Errorf("invalid parameter value: \"--%v\" flag parameter cannot have value", param.code)
				}
				if err := validateIfNewParameterValue(&parameterValues, param); err != nil {
					return map[string]string{}, err
				}
				// TODO improve types of parameterValues
				parameterValues[param.code] = "1"
				argFoundParamMatch = true
				break 
			}
			if usingEqualsAssignment {
				if err := validateIfNewParameterValue(&parameterValues, param); err != nil {
					return map[string]string{}, err
				}
				_, values := getEqualAssigntmentValues(rawArgValue)
				parameterValues[param.code] = values[1]
				argFoundParamMatch = true
				break 
			} else {
				hasNextArg := len(args) > (i + 1)
				if !hasNextArg {
					return map[string]string{}, fmt.Errorf("missing parameter value: \"--%v\" was not provided", param.code)
				}

				nextArgValue := args[i+1]
				if isParameterFormat(nextArgValue) {
					return map[string]string{}, fmt.Errorf("missing parameter value: \"--%v\" was not provided", param.code)
				}

				if err := validateIfNewParameterValue(&parameterValues, param); err != nil {
					return map[string]string{}, err
				}
				parameterValues[param.code] = nextArgValue
				i++
				argFoundParamMatch = true
				break
			}
		}

		if !argFoundParamMatch {
			if isParameterFormat(rawArgValue) {
				usingEqualsAssignment, values := getEqualAssigntmentValues(rawArgValue)
				if usingEqualsAssignment {
					return map[string]string{}, fmt.Errorf("unknown parameter provided: \"%v\"", truncateForError(values[0]))
				} else {
					return map[string]string{}, fmt.Errorf("unknown parameter provided: \"%v\"", truncateForError(rawArgValue))
				}
			} else {
				return map[string]string{}, fmt.Errorf("unknown value provided: \"%v\"", truncateForError(rawArgValue))
			}
		}
	}
	return parameterValues, nil
}

func validateIfNewParameterValue(parameterValues *map[string]string, param Parameter) error {
	if _, ok := (*parameterValues)[param.code]; ok {
		return fmt.Errorf("invalid parameter: \"--%v\" was provided twice", param.code)
	}
	return nil
}

func isParameterFormat(rawValue string) bool {
	return len(rawValue) >= 2 && rawValue[:2] == "--"
}

func getEqualAssigntmentValues(rawArgValue string) (bool, []string) {
	parts := strings.Split(rawArgValue, "=")
	if len(parts) == 1 {
		return false, []string{}
	}
	return true, []string{parts[0], strings.Join(parts[1:], "")}
}

func truncateForError(longString string) string {
	TRUNCATE_LIMIT := 30
	if len(longString) <= TRUNCATE_LIMIT {
		return longString
	}
	return fmt.Sprintf("%s...", longString[:TRUNCATE_LIMIT-3])
}

func validateParameterValues(parameterValues map[string]string) error {
	for key, value := range parameterValues {
		if strings.ReplaceAll(value, " ", "") == "" {
			return fmt.Errorf("missing parameter value: \"--%v\" was not provided", key)
		}
		if len(value) > PARAMETER_VALUE_MAX_SIZE {
			return fmt.Errorf("invalid parameter value: \"--%v\" exceeds max of %d", key, PARAMETER_VALUE_MAX_SIZE)
		}
	}
	return nil
}
