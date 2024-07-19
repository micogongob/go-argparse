package parse

import (
	"fmt"
	"strings"
	"testing"
)

var VeryVeryLongParamValue = strings.Repeat("X", 101)
var InvalidParamValues = []string{
	"*",
	"\\",
	"/",
	"<",
	">",
	"%",
	"$",
	"^",
	"#",
	"!",
	"?",
	"@",
	"&",
	"--",
}

func TestRequiredParameterKeyNotProvided(t *testing.T) {
	// given
	args := []string{SSS_CODE, "send-message"}
	t.Logf("Args: %v", args)
	testApp := newTestApp(t)

	// when
	parsedOutput, err := testApp.parseStrings(args)

	// then
	assertError(t, err, "missing required parameter/s: \"--queue-url\" was not provided")
	assertStringEquals(t, parsedOutput.helpMessage, "")
}

func TestRequiredParameterValueNotProvided(t *testing.T) {
	for i := 0; i < 2; i++ {
		// given
		var args []string
		if i == 0 {
			args = []string{SSS_CODE, "send-message", "--queue-url"}
		} else {
			args = []string{SSS_CODE, "send-message", "--queue-url="}
		}
		t.Logf("Args: %v", args)
		testApp := newTestApp(t)

		// when
		parsedOutput, err := testApp.parseStrings(args)

		// then
		assertError(t, err, "missing parameter value: \"--queue-url\" was not provided with a value")
		assertStringEquals(t, parsedOutput.helpMessage, "")
	}
}

func TestOptionalParameterValueNotProvided(t *testing.T) {
	for i := 0; i < 2; i++ {
		// given
		var args []string
		if i == 0 {
			args = []string{SSS_CODE, "list-queues", "--page-size"}
		} else {
			args = []string{SSS_CODE, "list-queues", "--page-size="}
		}
		t.Logf("Args: %v", args)
		testApp := newTestApp(t)

		// when
		parsedOutput, err := testApp.parseStrings(args)

		// then
		assertError(t, err, "missing parameter value: \"--page-size\" was not provided with a value")
		assertStringEquals(t, parsedOutput.helpMessage, "")
	}
}

func TestParameterValueIsBlank(t *testing.T) {
	for i := 0; i < 2; i++ {
		// given
		var args []string
		if i == 0 {
			args = []string{SSS_CODE, "send-message", "--queue-url", "   "}
		} else {
			args = []string{SSS_CODE, "send-message", "--queue-url=       "}
		}
		t.Logf("Args: %v", args)
		testApp := newTestApp(t)

		// when
		parsedOutput, err := testApp.parseStrings(args)

		// then
		assertError(t, err, "invalid parameter value: \"--queue-url\" cannot be blank")
		assertStringEquals(t, parsedOutput.helpMessage, "")
	}
}

func TestParameterValueCharLengthExceedsMax(t *testing.T) {
	// given
	for i := 0; i < 2; i++ {
		// given
		var args []string
		if i == 0 {
			args = []string{
				SSS_CODE,
				"send-message",
				"--queue-url",
				VeryVeryLongParamValue,
			}
		} else {
			args = []string{
				SSS_CODE,
				"send-message",
				fmt.Sprintf("--queue-url=%v", VeryVeryLongParamValue),
			}
		}
		t.Logf("Args: %v", args)
		testApp := newTestApp(t)

		// when
		parsedOutput, err := testApp.parseStrings(args)

		// then
		assertError(t, err, "invalid parameter value: \"--queue-url\" exceeds max of 1000")
		assertStringEquals(t, parsedOutput.helpMessage, "")
	}
}

func TestParameterInvalidValueInput(t *testing.T) {
	for _, paramValue := range InvalidParamValues {
		for i := 0; i < 2; i++ {
			// given
			var args []string
			if i == 0 {
				args = []string{
					SSS_CODE,
					"send-message",
					"--queue-url",
					paramValue,
				}
			} else {
				args = []string{
					SSS_CODE,
					"send-message",
					fmt.Sprintf("--queue-url=this-is-the-prefix %v", paramValue),
				}
			}
			t.Logf("Args: %v", args)
			testApp := newTestApp(t)

			// when
			parsedOutput, err := testApp.parseStrings(args)

			// then
			assertError(t, err, "invalid parameter value: \"--queue-url\" contains not allowed characters (*\\\\<>%$^#!?@&|--)")
			assertStringEquals(t, parsedOutput.helpMessage, "")
		}
	}
}

// TODO support list input and map?
// func TestNonListParameterDuplicateKey(t *testing.T) {
// 	for i := 0; i < 2; i++ {
// 		// given
// 		var args []string
// 		if i == 0 {
// 			args = []string{
// 				SSS_CODE,
// 				"send-message",
// 				"--queue-url",
// 				"https://sqs.us-east-1.amazonaws.com/00000000/TEST-account-created-queue",
// 				"--queue-url",
// 				"https://sqs.us-east-1.amazonaws.com/00000000/TEST-account-created-queue",
// 			}
// 		} else {
// 			args = []string{
// 				SSS_CODE,
// 				"send-message",
// 				"--queue-url=https://sqs.us-east-1.amazonaws.com/00000000/TEST-account-created-queue",
// 				"--queue-url=https://sqs.us-east-1.amazonaws.com/00000000/TEST-account-created-queue",
// 			}
// 		}
// 		t.Logf("Args: %v", args)
// 		testApp := newTestApp(t)

// 		// when
// 		parsedOutput, err := testApp.parseStrings(args)

// 		// then
// 		assertError(t, err, "invalid none collection type parameter: \"--queue-url\" was provided twice")
// 		assertStringEquals(t, parsedOutput.helpMessage, "")
// 	}
// }

func TestUnknownParameter(t *testing.T) {
	for i := 0; i < 2; i++ {
		// given
		var args []string
		if i == 0 {
			args = []string{
				SSS_CODE,
				"send-message",
				"--queue-url",
				"https://sqs.us-east-1.amazonaws.com/00000000/TEST-account-created-queue",
				"--another-url",
				"https://sqs.us-east-1.amazonaws.com/00000000/TEST-account-created-queue",
			}
		} else {
			args = []string{
				SSS_CODE,
				"send-message",
				"--queue-url=https://sqs.us-east-1.amazonaws.com/00000000/TEST-account-created-queue",
				"--another-url=https://sqs.us-east-1.amazonaws.com/00000000/TEST-account-created-queue",
			}
		}
		t.Logf("Args: %v", args)
		testApp := newTestApp(t)

		// when
		parsedOutput, err := testApp.parseStrings(args)

		// then
		assertError(t, err, "unknown parameter provided: \"--another-url\"")
		assertStringEquals(t, parsedOutput.helpMessage, "")
	}
}

func TestParameterUnknownValue(t *testing.T) {
	for i := 0; i < 2; i++ {
		// given
		var args []string
		if i == 0 {
			args = []string{
				SSS_CODE,
				"send-message",
				"--queue-url",
				"https://sqs.us-east-1.amazonaws.com/00000000/TEST-account-created-queue",
				"--debug",
				"out_of_place",
			}
		} else {
			args = []string{
				SSS_CODE,
				"send-message",
				"--queue-url=https://sqs.us-east-1.amazonaws.com/00000000/TEST-account-created-queue",
				"--debug",
				"out_of_place",
			}
		}
		t.Logf("Args: %v", args)
		testApp := newTestApp(t)

		// when
		parsedOutput, err := testApp.parseStrings(args)

		// then
		assertError(t, err, "unknown value provided: \"out_of_place\"")
		assertStringEquals(t, parsedOutput.helpMessage, "")
	}
}

func TestFlagParameterProvidedWithValue(t *testing.T) {
	for i := 0; i < 2; i++ {
		// given
		var args []string
		if i == 0 {
			args = []string{
				SSS_CODE,
				"list-queues",
				"-page-size",
				"10",
				"--debug",
				"some_value",
			}
		} else {
			args = []string{
				SSS_CODE,
				"list-queues",
				"-page-size",
				"--debug=some_value",
			}
		}
		t.Logf("Args: %v", args)
		testApp := newTestApp(t)

		// when
		parsedOutput, err := testApp.parseStrings(args)

		// then
		assertError(t, err, "invalid parameter value: \"--debug\" flag parameter cannot have value")
		assertStringEquals(t, parsedOutput.helpMessage, "")
	}
}
