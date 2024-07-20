package parse

import "testing"

func TestS4ChildCommands(t *testing.T) {

}

func TestSSSChildCommands(t *testing.T) {
    // given
    allArgs := [][]string{
        {SSS_CODE, "list-queues"},
    }
    expectedParameterValues := []map[string]ParameterValue{
        {
            "debug": ParameterValue{
                BooleanValue: false,
            },
        },
    }
    assertNumberEquals(t, len(expectedParameterValues), len(allArgs))
    for _, args := range allArgs {
        t.Logf("Args %v", args)
        app := newTestApp(t)

        // when
        parseOutput, err := app.parseStrings(args)

        // then
        assertNilError(t, err)
        assertStringEquals(t, parseOutput.helpMessage, "")
    }
}
