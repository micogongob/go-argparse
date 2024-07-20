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
	app := App{
		Code:        APP_CODE,
		Description: APP_DESC,
		Commands: []*Command{
			{
				Code:        SSS_CODE,
				Description: "SSS Queue Operations",
				Children: []*ChildCommand{
					{
						Code:        "version",
						Description: "Show SSS version",
					},
					{

						Code:        "list-queues",
						Description: "Lists SSS queues",
						Parameters: []*Parameter{
							{
								Code:        "page-size",
								Description: "pagination",
								IsOptional:  true,
								IsNumber:    true,
							},
							{
								Code:        "debug",
								Description: "DEBUG logging",
								IsOptional:  true,
								IsBoolean:   true,
							},
						},
					},
					{
						Code:        "send-message",
						Description: "Send string message to SSS queue",
						Parameters: []*Parameter{
							{
								Code:        "queue-url",
								Description: "the url of the SSS queue",
							},
							{
								Code:        "debug",
								Description: "DEBUG logging",
								IsOptional:  true,
								IsBoolean:   true,
							},
						},
					},
				},
			},
			{
				Code:        S4_CODE,
				Description: "S4 Bucket Operations",
				Children: []*ChildCommand{
					{
						Code:        "make-bucket",
						Description: "Create S4 bucket",
					},
					{
						Code:        "copy-objects",
						Description: "Copies object between s4 buckets",
					},
				},
			},
		},
	}

	err := app.validate()

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

func assertNumberEquals(t *testing.T, actual, expected int) {
	if actual != expected {
		t.Errorf("Numbers does not match\nactual:\n%v\nexpected:\n%v", actual, expected)
	}
}
