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
	command := NewCommand(NewCommandInput{})
	command.AddChildCommand(AddChildCommandInput{
		Code: "s4",
	})

	// when
	_, err := NewApp(NewAppInput{
		Code: "App",
		Commands: []*Command{
			command,
		},
	})

	// then
	assertError(t, err, "invalid command setup: Code is not provided")
}

func TestCommandCodeExceedsMax(t *testing.T) {
	// given
	command := NewCommand(NewCommandInput{
		Code: VeryLongLongCode,
	})
	command.AddChildCommand(AddChildCommandInput{
		Code: "make-bucket",
		Parameters: []Parameter{
			NewCommandParameter(NewCommandParameterInput{
				Code: "bucket-name",
			}),
		},
	})

	// when
	_, err := NewApp(NewAppInput{
		Code: "App",
		Commands: []*Command{
			command,
		},
	})

	// then
	assertError(t, err, fmt.Sprintf("invalid command setup: \"%v...\" exceeds max size of 15", VeryLongLongCode[:15]))
}

func TestCommandsAddedExceedMax(t *testing.T) {
	// given
	commands := make([]*Command, ListsMaxSize)
	for i := 0; i < len(commands); i++ {
		commands[i] = NewCommand(NewCommandInput{
			Code: fmt.Sprintf("s%d", i),
		})
		commands[i].AddChildCommand(AddChildCommandInput{
			Code: "make-bucket",
			Parameters: []Parameter{
				NewCommandParameter(NewCommandParameterInput{
					Code: "bucket-name",
				}),
			},
		})
	}

	// when
	_, err := NewApp(NewAppInput{
		Code: "App",
		Commands: commands,
	})

	// then
	assertError(t, err, "invalid command setup: commands added exceeds max size of 100")
}

func TestCommandInvalidCodeInput(t *testing.T) {
	for _, code := range InvalidCodes {
		// given
		command := NewCommand(NewCommandInput{
			Code: code,
		})
		command.AddChildCommand(AddChildCommandInput{
			Code: "make-bucket",
			Parameters: []Parameter{
				NewCommandParameter(NewCommandParameterInput{
					Code: "bucket-name",
				}),
			},
		})

		// when
		_, err := NewApp(NewAppInput{
			Code: "App",
			Commands: []*Command{
				command,
			},
		})

		// then
		assertError(t, err, fmt.Sprintf("invalid command setup: \"%v\" has invalid characters [A-Za-z0-9_-]", code))
	}
}

func TestCommandDuplicateCode(t *testing.T) {
	// given
	command0 := NewCommand(NewCommandInput{
		Code: "s4",
	})
	command0.AddChildCommand(AddChildCommandInput{
		Code: "make-bucket",
	})
	command1 := NewCommand(NewCommandInput{
		Code: "s4",
	})
	command1.AddChildCommand(AddChildCommandInput{
		Code: "make-bucket",
	})

	// when
	_, err := NewApp(NewAppInput{
		Code: "App",
		Commands: []*Command{
			command0,
			command1,
		},
	})

	// then
	assertError(t, err, "invalid command setup: \"s4\" is provided more than once")
}

func TestChildCommandCodeNotProvided(t *testing.T) {
	// given
	command := NewCommand(NewCommandInput{
		Code: "s4",
	})
	command.AddChildCommand(AddChildCommandInput{})

	// when
	_, err := NewApp(NewAppInput{
		Code: "App",
		Commands: []*Command{
			command,
		},
	})

	// then
	assertError(t, err, "invalid child command setup: \"s4.childCommands[*].Code\" is not provided")
}

func TestChildCommandCodeExceedsMax(t *testing.T) {
	// given
	command := NewCommand(NewCommandInput{
		Code: "s4",
	})
	command.AddChildCommand(AddChildCommandInput{
		Code: VeryLongLongCode,
		Parameters: []Parameter{
			NewCommandParameter(NewCommandParameterInput{
				Code: "bucket-name",
			}),
		},
	})

	// when
	_, err := NewApp(NewAppInput{
		Code: "App",
		Commands: []*Command{
			command,
		},
	})

	// then
	assertError(t, err, fmt.Sprintf("invalid child command setup: \"s4.%v...\" exceeds max size of 15", VeryLongLongCode[:15]))
}

func TestChildCommandsAddedExceedMax(t *testing.T) {
	// given
	command := NewCommand(NewCommandInput{
		Code: "s4",
	})
	for i := 0; i < ListsMaxSize; i++ {
		command.AddChildCommand(AddChildCommandInput{
			Code: fmt.Sprintf("make-bucket-v%d", i),
			Parameters: []Parameter{
				NewCommandParameter(NewCommandParameterInput{
					Code: "bucket-name",
				}),
			},
		})
	}

	// when
	_, err := NewApp(NewAppInput{
		Code: "App",
		Commands: []*Command{
			command,
		},
	})

	// then
	assertError(t, err, "invalid child command setup: \"s4.childCommands\" added exceeds max size of 100")
}

func TestChildCommandInvalidCodeInput(t *testing.T) {
	for _, code := range InvalidCodes {
		// given
		command := NewCommand(NewCommandInput{
			Code: "s4",
		})
		command.AddChildCommand(AddChildCommandInput{
			Code: code,
			Parameters: []Parameter{
				NewCommandParameter(NewCommandParameterInput{
					Code: "bucket-name",
				}),
			},
		})

		// when
		_, err := NewApp(NewAppInput{
			Code: "App",
			Commands: []*Command{
				command,
			},
		})

		// then
		assertError(t, err, fmt.Sprintf("invalid child command setup: \"s4.%v\" has invalid characters [A-Za-z0-9_-]", code))
	}
}

func TestChildCommandDuplicateCode(t *testing.T) {
	// given
	command := NewCommand(NewCommandInput{
		Code: "s4",
	})
	command.AddChildCommand(AddChildCommandInput{
		Code: "make-bucket",
	})
	command.AddChildCommand(AddChildCommandInput{
		Code: "make-bucket",
	})

	// when
	_, err := NewApp(NewAppInput{
		Code: "App",
		Commands: []*Command{
			command,
		},
	})

	// then
	assertError(t, err, "invalid child command setup: \"s4.make-bucket\" is provided more than once")
}

func TestParameterCodeNotProvided(t *testing.T) {
	// given
	command := NewCommand(NewCommandInput{
		Code: "s4",
	})
	command.AddChildCommand(AddChildCommandInput{
		Code: "make-bucket",
		Parameters: []Parameter{
			NewCommandParameter(NewCommandParameterInput{}),
		},
	})

	// when
	_, err := NewApp(NewAppInput{
		Code: "App",
		Commands: []*Command{
			command,
		},
	})

	// then
	assertError(t, err, "invalid parameter setup: \"s4.make-bucket.parameters[*].Code\" is not provided")
}

func TestParameterCodeExceedsMax(t *testing.T) {
	// given
	command := NewCommand(NewCommandInput{
		Code: "s4",
	})
	command.AddChildCommand(AddChildCommandInput{
		Code: "make-bucket",
		Parameters: []Parameter{
			NewCommandParameter(NewCommandParameterInput{
				Code: VeryLongLongCode,
			}),
		},
	})

	// when
	_, err := NewApp(NewAppInput{
		Code: "App",
		Commands: []*Command{
			command,
		},
	})

	// then
	assertError(t, err, fmt.Sprintf("invalid parameter setup: \"s4.make-bucket.%v...\" exceeds max size of 15", VeryLongLongCode[:15]))
}

func TestParametersAddedExceedMax(t *testing.T) {
	// given
	command := NewCommand(NewCommandInput{
		Code: "s4",
	})
	parameters := make([]Parameter, ListsMaxSize)
	for i := 0; i < ListsMaxSize; i++ {
		parameters[i] = NewCommandParameter(NewCommandParameterInput{
			Code: fmt.Sprintf("bucket-name-v%d", i),
		})
	}
	command.AddChildCommand(AddChildCommandInput{
		Code: "make-bucket",
		Parameters: parameters,
	})

	// when
	_, err := NewApp(NewAppInput{
		Code: "App",
		Commands: []*Command{
			command,
		},
	})

	// then
	assertError(t, err, "invalid parameter setup: \"s4.make-bucket.parameters\" added exceeds max size of 100")
}

func TestParameterInvalidCodeInput(t *testing.T) {
	for _, code := range InvalidCodes {
		// given
		command := NewCommand(NewCommandInput{
			Code: "s4",
		})
		command.AddChildCommand(AddChildCommandInput{
			Code: "make-bucket",
			Parameters: []Parameter{
				NewCommandParameter(NewCommandParameterInput{
					Code: code,
				}),
			},
		})

		// when
		_, err := NewApp(NewAppInput{
			Code: "App",
			Commands: []*Command{
				command,
			},
		})

		// then
		assertError(t, err, fmt.Sprintf("invalid parameter setup: \"s4.make-bucket.%v\" has invalid characters [A-Za-z0-9_-]", code))
	}
}

func TestParameterDuplicateCode(t *testing.T) {
	// given
	command := NewCommand(NewCommandInput{
		Code: "s4",
	})
	command.AddChildCommand(AddChildCommandInput{
		Code: "make-bucket",
		Parameters: []Parameter{
			NewCommandParameter(NewCommandParameterInput{
				Code: "bucket-name",
			}),
			NewCommandParameter(NewCommandParameterInput{
				Code: "bucket-name",
			}),
		},
	})

	// when
	_, err := NewApp(NewAppInput{
		Code: "App",
		Commands: []*Command{
			command,
		},
	})

	// then
	assertError(t, err, "invalid parameter setup: \"s4.make-bucket.bucket-name\" is provided more than once")
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
				Code:        "make-bucket",
				Description: "Toggle make-bucket logging",
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
	assertError(t, err, "invalid parameter setup: \"make-bucket\" cannot be required and a flag at the same time")
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
	assertError(t, err, "invalid command setup: \"test\" should have atleast one child command")
}
