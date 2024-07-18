package parse

import "testing"

const (
	APP_CODE = "ows"
	APP_DESC = "Owsome cli"
	SSS_CODE = "sss"
	S4_CODE  = "s4"
)

var HelpCommandAliases = []string{"help", "--help", "-h"}

func newTestApp(t *testing.T) App {
	sssCommand := NewCommand(NewCommandInput{
		Code:        SSS_CODE,
		Description: "SSS Queue Operations",
	})
	sssCommand.AddChildCommand(AddChildCommandInput{
		Code:        "version",
		Description: "Show SSS version",
	})
	sssCommand.AddChildCommand(AddChildCommandInput{
		Code:        "list-queues",
		Description: "Lists SSS queues",
		Parameters: []Parameter{
			NewCommandParameter(NewCommandParameterInput{
				Code:        "page-size",
				Description: "pagination",
				IsOptional:  true,
			}),
			NewCommandParameter(NewCommandParameterInput{
				Code:        "debug",
				Description: "DEBUG logging",
				IsOptional:  true,
				IsFlag:      true,
			}),
		},
	})
	sssCommand.AddChildCommand(AddChildCommandInput{
		Code:        "send-message",
		Description: "Send string message to SSS queue",
		Parameters: []Parameter{
			NewCommandParameter(NewCommandParameterInput{
				Code:        "queue-name",
				Description: "the name of the SSS queue",
			}),
		},
	})

	s4Command := NewCommand(NewCommandInput{
		Code:        S4_CODE,
		Description: "S4 Bucket Operations",
	})
	s4Command.AddChildCommand(AddChildCommandInput{
		Code:        "make-bucket",
		Description: "Create S4 bucket",
	})
	s4Command.AddChildCommand(AddChildCommandInput{
		Code:        "copy-objects",
		Description: "Copies object between s4 buckets",
	})

	app, err := NewApp(NewAppInput{
		Code:        APP_CODE,
		Description: APP_DESC,
		Commands: []*Command{
			sssCommand,
			s4Command,
		},
	})
	assertNilError(t, err)
	return app
}

func assertError(t *testing.T, err error, message string) {
	if err == nil {
		t.Errorf("Error did not happen: %v", message)
	} else {
		assertStringEquals(t, err.Error(), message)
	}
}

func assertNilError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Unexpected error encountered: %v", err)
	}
}

func assertStringEquals(t *testing.T, actual, expected string) {
	if actual != expected {
		t.Errorf("Strings does not match\nactual:\n%v\nexpected:\n%v", actual, expected)
	}
}
