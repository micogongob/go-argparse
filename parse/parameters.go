package parse

import (
	"fmt"
	"strconv"
	"strings"
)

func (command *ChildCommand) requiredParameters() []*Parameter {
	params := []*Parameter{}
	for _, param := range command.Parameters {
		if !param.IsOptional {
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
	parameterValues, err := filterParameterValues(command.Parameters, args)
	if err != nil {
		return err
	}
	finalParameterValues, err := validateParameterValues(command.Parameters, parameterValues)
	if err != nil {
		return err
	}

	// TODO return to on trigger
	fmt.Println(finalParameterValues)
	return nil
}

func (parameter *Parameter) matchesArg(rawArgValue string) (bool, bool) {
	usingEqualsAssignment, values := getEqualAssigntmentValues(rawArgValue)
	if usingEqualsAssignment {
		return fmt.Sprintf("--%v", parameter.Code) == values[0], usingEqualsAssignment
	} else {
		return fmt.Sprintf("--%v", parameter.Code) == rawArgValue, usingEqualsAssignment
	}
}

func toValidationMsgFormat(params []*Parameter) string {
	s := []string{}
	for _, v := range params {
		s = append(s, fmt.Sprintf("--%v", v.Code))
	}
	return strings.Join(s, ",")
}

func validateRequiredParameters(parameters []*Parameter, args []string) error {
	notProvidedRequiredParameters := []*Parameter{}

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

func filterParameterValues(parameters []*Parameter, args []string) (map[string]string, error) {
	parameterValues := map[string]string{}

	for i := 0; i < len(args); i++ {
		rawArgValue := args[i]

		argFoundParamMatch := false
		for _, param := range parameters {
			matches, usingEqualsAssignment := param.matchesArg(rawArgValue)
			if !matches {
				continue
			}

			if param.IsBoolean {
				if usingEqualsAssignment, _ := getEqualAssigntmentValues(rawArgValue); usingEqualsAssignment {
					return map[string]string{}, fmt.Errorf("invalid parameter value: \"--%v\" boolean parameter cannot have value", param.Code)
				}
				if err := validateIfNewParameterValue(&parameterValues, *param); err != nil {
					return map[string]string{}, err
				}
				parameterValues[param.Code] = "any"
				argFoundParamMatch = true
				break
			}
			if usingEqualsAssignment {
				if err := validateIfNewParameterValue(&parameterValues, *param); err != nil {
					return map[string]string{}, err
				}
				_, values := getEqualAssigntmentValues(rawArgValue)
				parameterValues[param.Code] = values[1]
				argFoundParamMatch = true
				break
			} else {
				hasNextArg := len(args) > (i + 1)
				if !hasNextArg {
					return map[string]string{}, fmt.Errorf("missing parameter value: \"--%v\" was not provided", param.Code)
				}

				nextArgValue := args[i+1]
				if isParameterFormat(nextArgValue) {
					return map[string]string{}, fmt.Errorf("missing parameter value: \"--%v\" was not provided", param.Code)
				}

				if err := validateIfNewParameterValue(&parameterValues, *param); err != nil {
					return map[string]string{}, err
				}
				parameterValues[param.Code] = nextArgValue
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
	if _, ok := (*parameterValues)[param.Code]; ok {
		return fmt.Errorf("invalid parameter: \"--%v\" was provided twice", param.Code)
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

func validateParameterValues(parameters []*Parameter, parameterValues map[string]string) (map[string]ParameterValue, error) {
	PARAMETER_MAX_CHAR_LENGTH := 1000
	PARAMETER_MAX_NUMBER_VALUE := 2147483647

	finalParameterValues := map[string]ParameterValue{}
	for _, param := range parameters {
		value, ok := parameterValues[param.Code]
		if ok {
			if !param.IsBoolean {
				if strings.ReplaceAll(value, " ", "") == "" {
					return map[string]ParameterValue{}, fmt.Errorf("missing parameter value: \"--%v\" was not provided", param.Code)
				}
			}
			if param.IsNumber {
				asNumber, err := strconv.Atoi(value)
				if err != nil {
					return map[string]ParameterValue{}, fmt.Errorf("invalid parameter value: \"--%v\" expected numeric value", param.Code)
				}
				if asNumber > PARAMETER_MAX_NUMBER_VALUE {
					return map[string]ParameterValue{}, fmt.Errorf("invalid parameter value: \"--%v\" exceeds max number of %d", param.Code, PARAMETER_MAX_NUMBER_VALUE)
				}
				finalParameterValues[param.Code] = ParameterValue{
					NumberValue: asNumber,
				}
			} else if param.IsBoolean {
				finalParameterValues[param.Code] = ParameterValue{
					BooleanValue: true,
				}
			} else {
				if len(value) > PARAMETER_MAX_CHAR_LENGTH {
					return map[string]ParameterValue{}, fmt.Errorf("invalid parameter value: \"--%v\" exceeds max of %d", param.Code, PARAMETER_MAX_CHAR_LENGTH)
				}
				finalParameterValues[param.Code] = ParameterValue{
					StringValue: value,
				}
			}
		} else if !ok && param.IsBoolean {
			finalParameterValues[param.Code] = ParameterValue{
				BooleanValue: false,
			}
		}
	}
	return finalParameterValues, nil
}
