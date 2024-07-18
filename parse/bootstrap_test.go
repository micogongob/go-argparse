package parse

import "testing"

func TestCodesInvalidInput(t *testing.T) {

}

func TestDuplicateCodes(t *testing.T) {

}

func TestRequiredParameterButFlag(t *testing.T) {
	// given
	command := NewCommand(NewCommandInput{
		Code:        "test",
		Description: "Test Me",
	})
	command.AddChildCommand(AddChildCommandInput{
		Code:        "file-name",
		Description: "Name of file to test",
		Parameters: []Parameter{
			NewCommandParameter(NewCommandParameterInput{
				Code:        "verbose",
				Description: "Toggle verbose logging",
				IsFlag:      true,
			}),
		},
	})

	// when
	_, err := NewApp(NewAppInput{
		Code:        "tester",
		Description: "Tester App",
		Commands: []*Command{
			command,
		},
	})

	// then
	assertError(t, err, "invalid parameter setup: verbose cannot be required and a flag at the same time")
}

func TestCommandWithoutChildCommand(t *testing.T) {
	// given
	command := NewCommand(NewCommandInput{
		Code:        "test",
		Description: "Test Me",
	})

	// when
	_, err := NewApp(NewAppInput{
		Code:        "tester",
		Description: "Tester App",
		Commands: []*Command{
			command,
		},
	})

	// then
	assertError(t, err, "invalid command setup: test should have atleast one child command")
}
