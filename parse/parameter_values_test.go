package parse

import (
	"reflect"
	"testing"
)

func TestSSSChildCommands(t *testing.T) {
	// given
	allArgs := [][]string{
		{SSS_COMMAND, SSS_LIST_QUEUES},
		{SSS_COMMAND, SSS_LIST_QUEUES, "--page-size", "420"},
		{SSS_COMMAND, SSS_LIST_QUEUES, "--page-size=420"},
		{SSS_COMMAND, SSS_LIST_QUEUES, "--page-size", "420", "--debug"},
		{SSS_COMMAND, SSS_LIST_QUEUES, "--page-size=420", "--debug"},
		{SSS_COMMAND, SSS_LIST_QUEUES, "--debug", "--page-size", "420"},
		{SSS_COMMAND, SSS_LIST_QUEUES, "--debug", "--page-size=420"},
		{SSS_COMMAND, SSS_SEND_MESSAGE, "--queue-url", "http://localhost:4566/000000"},
		{SSS_COMMAND, SSS_SEND_MESSAGE, "--queue-url=http://localhost:4566/000000"},
		{SSS_COMMAND, SSS_SEND_MESSAGE, "--queue-url", "http://localhost:4566/000000", "--debug"},
		{SSS_COMMAND, SSS_SEND_MESSAGE, "--queue-url=http://localhost:4566/000000", "--debug"},
		{SSS_COMMAND, SSS_SEND_MESSAGE, "--debug", "--queue-url", "http://localhost:4566/000000"},
		{SSS_COMMAND, SSS_SEND_MESSAGE, "--debug", "--queue-url=http://localhost:4566/000000"},
	}
	expectedParameterValues := []map[string]ParameterValue{
		{
			"debug": ParameterValue{
				BooleanValue: false,
			},
		},
		{
			"debug": ParameterValue{
				BooleanValue: false,
			},
			"page-size": ParameterValue{
				NumberValue: 420,
			},
		},
		{
			"debug": ParameterValue{
				BooleanValue: false,
			},
			"page-size": ParameterValue{
				NumberValue: 420,
			},
		},
		{
			"debug": ParameterValue{
				BooleanValue: true,
			},
			"page-size": ParameterValue{
				NumberValue: 420,
			},
		},
		{
			"debug": ParameterValue{
				BooleanValue: true,
			},
			"page-size": ParameterValue{
				NumberValue: 420,
			},
		},
		{
			"debug": ParameterValue{
				BooleanValue: true,
			},
			"page-size": ParameterValue{
				NumberValue: 420,
			},
		},
		{
			"debug": ParameterValue{
				BooleanValue: true,
			},
			"page-size": ParameterValue{
				NumberValue: 420,
			},
		},
		{
			"debug": ParameterValue{
				BooleanValue: false,
			},
			"queue-url": ParameterValue{
				StringValue: "http://localhost:4566/000000",
			},
		},
		{
			"debug": ParameterValue{
				BooleanValue: false,
			},
			"queue-url": ParameterValue{
				StringValue: "http://localhost:4566/000000",
			},
		},
		{
			"debug": ParameterValue{
				BooleanValue: true,
			},
			"queue-url": ParameterValue{
				StringValue: "http://localhost:4566/000000",
			},
		},
		{
			"debug": ParameterValue{
				BooleanValue: true,
			},
			"queue-url": ParameterValue{
				StringValue: "http://localhost:4566/000000",
			},
		},
		{
			"debug": ParameterValue{
				BooleanValue: true,
			},
			"queue-url": ParameterValue{
				StringValue: "http://localhost:4566/000000",
			},
		},
		{
			"debug": ParameterValue{
				BooleanValue: true,
			},
			"queue-url": ParameterValue{
				StringValue: "http://localhost:4566/000000",
			},
		},
	}
	assertNumberEquals(t, len(expectedParameterValues), len(allArgs))

	// then
	for i, args := range allArgs {
		assertParameterValues(t, args, expectedParameterValues[i])
	}
}

func assertParameterValues(t *testing.T, args []string, expectedParameterValues map[string]ParameterValue) {
	// given
	assertFunc := func(values map[string]ParameterValue) error {
		assertParameterValuesMapEquals(t, values, expectedParameterValues)
		return nil
	}
	t.Logf("Args %v", args)
	app := newOwsAppWithCommandHandler(t, args[1], assertFunc)

	// when
	parseOutput, err := app.parseStrings(args)

	// then
	assertNilError(t, err)
	assertStringEquals(t, parseOutput.helpMessage, "")
}

func assertParameterValuesMapEquals(t *testing.T, actual, expected map[string]ParameterValue) {
	if reflect.DeepEqual(actual, expected) {
		t.Errorf("ParameterValue map does not match\nactual:\n%v\nexpected:\n%v", actual, expected)
	}
}
