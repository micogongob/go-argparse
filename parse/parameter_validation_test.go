package parse

import "testing"

func TestMissingParameterRequired(t *testing.T) {
	// given
	args := []string{SSS_CODE, "send-message"}
	t.Logf("Args: %v", args)
	testApp := newTestApp(t)

	// when
	parsedOutput, err := testApp.parseStrings(args)

	// then
	assertError(t, err, "missing required parameter/s: --queue-name")
	assertStringEquals(t, parsedOutput.helpMessage, "")
}
