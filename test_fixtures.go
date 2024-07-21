package parse

import "testing"

const (
	APP_CODE         = "ows"
	APP_DESC         = "Owsome cli"
	SSS_COMMAND      = "sss"
	SSS_LIST_QUEUES  = "list-queues"
	SSS_SEND_MESSAGE = "send-message"
	S4_COMMAND       = "s4"
	S4_MAKE_BUCKET   = "make-bucket"
	S4_BUCKET_NAME   = "bucket-name"
	S4_COPY_OBJECTS  = "copy-objects"
)

var HelpCommandAliases = []string{"help", "--help", "-h"}

func newOwsApp(t *testing.T) App {
	app := App{
		Code:        APP_CODE,
		Description: APP_DESC,
		Commands: []*Command{
			{
				Code:        SSS_COMMAND,
				Description: "SSS Queue Operations",
				Children: []*ChildCommand{
					{
						Code:           "version",
						Description:    "Show SSS version",
						CommandHandler: noopCommandHandler,
					},
					{

						Code:           SSS_LIST_QUEUES,
						Description:    "Lists SSS queues",
						CommandHandler: noopCommandHandler,
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
						Code:           SSS_SEND_MESSAGE,
						Description:    "Send string message to SSS queue",
						CommandHandler: noopCommandHandler,
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
				Code:        S4_COMMAND,
				Description: "S4 Bucket Operations",
				Children: []*ChildCommand{
					{
						Code:           S4_MAKE_BUCKET,
						Description:    "Create S4 bucket",
						CommandHandler: noopCommandHandler,
					},
					{
						Code:           S4_COPY_OBJECTS,
						Description:    "Copies object between s4 buckets",
						CommandHandler: noopCommandHandler,
					},
				},
			},
		},
	}

	err := app.validate()

	assertNilError(t, err)
	return app
}

func noopCommandHandler(values map[string]ParameterValue) error {
	return nil
}

func newOwsAppWithCommandHandler(t *testing.T, childCommandCode string, commandHandler func(map[string]ParameterValue) error) App {
	app := newOwsApp(t)
	for i, command := range app.Commands {

		for j, childCommand := range command.Children {
			if childCommand.Code == childCommandCode {
				app.Commands[i].Children[j].CommandHandler = commandHandler
				return app
			}
		}
	}
	t.Errorf("childCommandCode=%v is not found", childCommandCode)
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
