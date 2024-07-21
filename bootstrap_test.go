package parse

import (
	"fmt"
	"strings"
	"testing"
)

var InvalidCodes = []string{
	" ",
	"s ",
	"s!",
	"s@",
	"s#",
	"s$",
	"s%",
	"s^",
	"s&",
	"s*",
	"s(",
	"s)",
	"s=",
	"s+",
	"s{",
	"s}",
	"s[",
	"s]",
	"s]",
	"s:",
	"s;",
	"s>",
	"s.",
	"s,",
}

var VeryLongLongCode = strings.Repeat("s", 16)
var ListsMaxSize = 101

func TestCommandCodeNotProvided(t *testing.T) {
	// given
	app := App{
		Code: APP_CODE,
		Commands: []*Command{
			{
				Children: []*ChildCommand{
					{
						Code:           S4_COMMAND,
						CommandHandler: noopCommandHandler,
					},
				},
			},
		},
	}

	// when
	err := app.validate()

	// then
	assertError(t, err, "invalid command setup: Code is not provided")
}

func TestCommandCodeCharLengthExceedsMax(t *testing.T) {
	// given
	app := App{
		Code: APP_CODE,
		Commands: []*Command{
			{
				Code: VeryLongLongCode,
				Children: []*ChildCommand{
					{
						Code:           S4_MAKE_BUCKET,
						CommandHandler: noopCommandHandler,
						Parameters: []*Parameter{
							{
								Code: S4_BUCKET_NAME,
							},
						},
					},
				},
			},
		},
	}

	// when
	err := app.validate()

	// then
	assertError(t, err, fmt.Sprintf("invalid command setup: \"%v...\" exceeds max size of 15", VeryLongLongCode[:15]))
}

func TestCommandsAddedExceedMax(t *testing.T) {
	// given
	commands := make([]*Command, ListsMaxSize)
	for i := 0; i < len(commands); i++ {
		commands[i] = &Command{
			Code: fmt.Sprintf("s%d", i),
			Children: []*ChildCommand{
				{
					Code:           S4_MAKE_BUCKET,
					CommandHandler: noopCommandHandler,
					Parameters: []*Parameter{
						{
							Code: S4_BUCKET_NAME,
						},
					},
				},
			},
		}
	}
	app := App{
		Code:     APP_CODE,
		Commands: commands,
	}

	// when
	err := app.validate()

	// then
	assertError(t, err, "invalid command setup: commands added exceeds max size of 100")
}

func TestCommandInvalidCodeInput(t *testing.T) {
	for _, code := range InvalidCodes {
		// given
		app := App{
			Code: APP_CODE,
			Commands: []*Command{
				{
					Code: code,
					Children: []*ChildCommand{
						{
							Code:           S4_MAKE_BUCKET,
							CommandHandler: noopCommandHandler,
							Parameters: []*Parameter{
								{
									Code: S4_BUCKET_NAME,
								},
							},
						},
					},
				},
			},
		}

		// when
		err := app.validate()

		// then
		assertError(t, err, fmt.Sprintf("invalid command setup: \"%v\" has invalid characters [A-Za-z0-9_-]", code))
	}
}

func TestCommandDuplicateCode(t *testing.T) {
	// given
	app := App{
		Code: APP_CODE,
		Commands: []*Command{
			{
				Code: S4_COMMAND,
				Children: []*ChildCommand{
					{
						Code:           S4_MAKE_BUCKET,
						CommandHandler: noopCommandHandler,
					},
				},
			},
			{
				Code: S4_COMMAND,
				Children: []*ChildCommand{
					{
						Code: S4_MAKE_BUCKET,
					},
				},
			},
		},
	}

	// when
	err := app.validate()

	// then
	assertError(t, err, "invalid command setup: \"s4\" is provided more than once")
}

func TestChildCommandCodeNotProvided(t *testing.T) {
	// given
	app := App{
		Code: APP_CODE,
		Commands: []*Command{
			{
				Code: S4_COMMAND,
				Children: []*ChildCommand{
					{},
				},
			},
		},
	}

	// when
	err := app.validate()

	// then
	assertError(t, err, "invalid child command setup: \"s4.childCommands[*].Code\" is not provided")
}

func TestChildCommandCodeCharLengthExceedsMax(t *testing.T) {
	// given
	app := App{
		Code: APP_CODE,
		Commands: []*Command{
			{
				Code: S4_COMMAND,
				Children: []*ChildCommand{
					{
						Code:           VeryLongLongCode,
						CommandHandler: noopCommandHandler,
						Parameters: []*Parameter{
							{
								Code: S4_BUCKET_NAME,
							},
						},
					},
				},
			},
		},
	}

	// when
	err := app.validate()

	// then
	assertError(t, err, fmt.Sprintf("invalid child command setup: \"s4.%v...\" exceeds max size of 15", VeryLongLongCode[:15]))
}

func TestChildCommandsAddedExceedMax(t *testing.T) {
	// given
	children := make([]*ChildCommand, ListsMaxSize)
	for i := 0; i < ListsMaxSize; i++ {
		children[i] = &ChildCommand{
			Code:           fmt.Sprintf("make-bucket-v%d", i),
			CommandHandler: noopCommandHandler,
			Parameters: []*Parameter{
				{
					Code: S4_BUCKET_NAME,
				},
			},
		}
	}
	app := App{
		Code: APP_CODE,
		Commands: []*Command{
			{
				Code:     S4_COMMAND,
				Children: children,
			},
		},
	}

	// when
	err := app.validate()

	// then
	assertError(t, err, "invalid child command setup: \"s4.childCommands\" added exceeds max size of 100")
}

func TestChildCommandInvalidCodeInput(t *testing.T) {
	for _, code := range InvalidCodes {
		// given
		app := App{
			Code: APP_CODE,
			Commands: []*Command{
				{
					Code: S4_COMMAND,
					Children: []*ChildCommand{
						{
							Code:           code,
							CommandHandler: noopCommandHandler,
							Parameters: []*Parameter{
								{
									Code: S4_BUCKET_NAME,
								},
							},
						},
					},
				},
			},
		}

		// when
		err := app.validate()

		// then
		assertError(t, err, fmt.Sprintf("invalid child command setup: \"s4.%v\" has invalid characters [A-Za-z0-9_-]", code))
	}
}

func TestChildCommandDuplicateCode(t *testing.T) {
	// given
	app := App{
		Code: APP_CODE,
		Commands: []*Command{
			{
				Code: S4_COMMAND,
				Children: []*ChildCommand{
					{
						Code:           S4_MAKE_BUCKET,
						CommandHandler: noopCommandHandler,
					},
					{
						Code:           S4_MAKE_BUCKET,
						CommandHandler: noopCommandHandler,
					},
				},
			},
		},
	}

	// when
	err := app.validate()

	// then
	assertError(t, err, "invalid child command setup: \"s4.make-bucket\" is provided more than once")
}

func TestParameterCodeNotProvided(t *testing.T) {
	// given
	app := App{
		Code: APP_CODE,
		Commands: []*Command{
			{
				Code: S4_COMMAND,
				Children: []*ChildCommand{
					{
						Code:           S4_MAKE_BUCKET,
						CommandHandler: noopCommandHandler,
						Parameters: []*Parameter{
							{},
						},
					},
				},
			},
		},
	}

	// when
	err := app.validate()

	// then
	assertError(t, err, "invalid parameter setup: \"s4.make-bucket.Parameters[*].Code\" is not provided")
}

func TestParameterCodeCharLengthExceedsMax(t *testing.T) {
	// given
	app := App{
		Code: APP_CODE,
		Commands: []*Command{
			{
				Code: S4_COMMAND,
				Children: []*ChildCommand{
					{
						Code:           S4_MAKE_BUCKET,
						CommandHandler: noopCommandHandler,
						Parameters: []*Parameter{
							{
								Code: VeryLongLongCode,
							},
						},
					},
				},
			},
		},
	}

	// when
	err := app.validate()

	// then
	assertError(t, err, fmt.Sprintf("invalid parameter setup: \"s4.make-bucket.%v...\" exceeds max size of 15", VeryLongLongCode[:15]))
}

func TestParametersAddedExceedMax(t *testing.T) {
	// given
	parameters := make([]*Parameter, ListsMaxSize)
	for i := 0; i < ListsMaxSize; i++ {
		parameters[i] = &Parameter{
			Code: fmt.Sprintf("bucket-name-v%d", i),
		}
	}

	app := App{
		Code: APP_CODE,
		Commands: []*Command{
			{
				Code: S4_COMMAND,
				Children: []*ChildCommand{
					{
						Code:           S4_MAKE_BUCKET,
						CommandHandler: noopCommandHandler,
						Parameters:     parameters,
					},
				},
			},
		},
	}

	// when
	err := app.validate()

	// then
	assertError(t, err, "invalid parameter setup: \"s4.make-bucket.Parameters\" added exceeds max size of 100")
}

func TestParameterInvalidCodeInput(t *testing.T) {
	for _, code := range InvalidCodes {
		// given
		app := App{
			Code: APP_CODE,
			Commands: []*Command{
				{
					Code: S4_COMMAND,
					Children: []*ChildCommand{
						{
							Code:           S4_MAKE_BUCKET,
							CommandHandler: noopCommandHandler,
							Parameters: []*Parameter{
								{
									Code: code,
								},
							},
						},
					},
				},
			},
		}

		// when
		err := app.validate()

		// then
		assertError(t, err, fmt.Sprintf("invalid parameter setup: \"s4.make-bucket.%v\" has invalid characters [A-Za-z0-9_-]", code))
	}
}

func TestParameterDuplicateCode(t *testing.T) {
	// given
	app := App{
		Code: APP_CODE,
		Commands: []*Command{
			{
				Code: S4_COMMAND,
				Children: []*ChildCommand{
					{
						Code:           S4_MAKE_BUCKET,
						CommandHandler: noopCommandHandler,
						Parameters: []*Parameter{
							{
								Code: S4_BUCKET_NAME,
							},
							{
								Code: S4_BUCKET_NAME,
							},
						},
					},
				},
			},
		},
	}

	// when
	err := app.validate()

	// then
	assertError(t, err, "invalid parameter setup: \"s4.make-bucket.bucket-name\" is provided more than once")
}

func TestRequiredParameterButBoolean(t *testing.T) {
	// given
	app := App{
		Code: APP_CODE,
		Commands: []*Command{
			{
				Code: S4_COMMAND,
				Children: []*ChildCommand{
					{
						Code:           S4_MAKE_BUCKET,
						CommandHandler: noopCommandHandler,
						Parameters: []*Parameter{
							{
								Code:      S4_BUCKET_NAME,
								IsBoolean: true,
							},
						},
					},
				},
			},
		},
	}

	// when
	err := app.validate()

	// then
	assertError(t, err, "invalid parameter setup: \"s4.make-bucket.bucket-name\" cannot be required and a boolean at the same time")
}

func TestCommandWithoutChildCommand(t *testing.T) {
	// given
	app := App{
		Code:        "tester",
		Description: "Tester App",
		Commands: []*Command{
			{
				Code:        "test",
				Description: "Test Me",
			},
		},
	}

	// when
	err := app.validate()

	// then
	assertError(t, err, "invalid command setup: \"test\" should have atleast one child command")
}

func TestParameterMultipleTypeDefined(t *testing.T) {
	// given
	parameters := []*Parameter{
		{
			Code:       S4_BUCKET_NAME,
			IsOptional: true,
			IsBoolean:  true,
			IsNumber:   true,
		},
	}
	errorMsgs := []string{
		"invalid parameter setup: \"s4.make-bucket.bucket-name\" cannot be a number and a boolean at the same time",
	}
	for i := 0; i < len(parameters); i++ {
		app := App{
			Code: APP_CODE,
			Commands: []*Command{
				{
					Code: S4_COMMAND,
					Children: []*ChildCommand{
						{
							Code:           S4_MAKE_BUCKET,
							CommandHandler: noopCommandHandler,
							Parameters: []*Parameter{
								parameters[i],
							},
						},
					},
				},
			},
		}

		// when
		err := app.validate()

		// then
		assertError(t, err, errorMsgs[i])
	}
}

func TestÇhildCommandNoHandler(t *testing.T) {
	// given
	app := App{
		Code: APP_CODE,
		Commands: []*Command{
			{
				Code: S4_COMMAND,
				Children: []*ChildCommand{
					{
						Code: S4_MAKE_BUCKET,
					},
				},
			},
		},
	}

	// when
	err := app.validate()

	// then
	assertError(t, err, "invalid child command setup: \"s4.make-bucket.CommandHandler\" is not provided")
}
