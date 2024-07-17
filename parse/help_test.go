package parse

import (
	"testing"
)

var TestApp = NewApp(APP_CODE, APP_DESC)

func init() {
	TestApp.AddCommand(SSS_CODE, "SSS Operations", []string{}, []Command{})
	TestApp.AddCommand(S4_CODE, "S4 Operations", []string{}, []Command{})
}

func TestAppHelp(t *testing.T) {
	for i := 0; i <= len(HelpCommandAliases); i++ {
		var args []string
		if i == len(HelpCommandAliases) {
			args = []string{}
		} else {
			args = []string{HelpCommandAliases[i]}
		}

		parsedOutput, err := TestApp.parseStrings(args)

		actual := parsedOutput.helpMessage

		if err != nil {
			t.Errorf("Unexpected error. %v", err)
		}

		expected := `Owsome cli

  usage: ows [command] [...arguments]

  commands:
    help - Show help. Alternatives: --help, -h
    sss - SSS Operations
    s4 - S4 Operations
`
		if actual != expected {
			t.Errorf("index: %v - \nactual:\n%v\nexpected:\n%v", i, actual, expected)
		}
	}

}
