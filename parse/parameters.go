package parse

import (
	"fmt"
	"strings"
)

func (command *ChildCommand) requiredParameters() []string {
	params := []string{}
	for _, param := range command.parameters {
		if !param.isOptional {
			params = append(params, param.code)
		}
	}
	return params
}

func parametersListToHelp(params []string) string {
	s := []string{}
	for _, v := range params {
		s = append(s, fmt.Sprintf("--%v", v))
	}
	return strings.Join(s, ",")
}
