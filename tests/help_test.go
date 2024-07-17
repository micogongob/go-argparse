package tests

import (
	"github.com/micogongob/go-argparse/parse"
	"testing"
)

var HELP_COMMAND_ALIASES = []string{}

func init() {
	parse.Setup(APP_CODE, APP_DESC)
	parse.AddCommand(SSS_COMMAND)
	parse.AddCommand(S4_COMMAND)
}

func TestAppHelp(t *testing.T) {
	for i := 0; i <= len(HELP_COMMAND_ALIASES); i++ {
		var args []string
		if i == len(HELP_COMMAND_ALIASES) {
			args = []string{}
		} else {
			args = []string{HELP_COMMAND_ALIASES[i]}
		}

		actual, err := parse.ParseStrings(args)

		if err != nil {
			t.Errorf("Unexpected error. %v", err)
		}

		expected := `  Owsome cli

  usage: ows [command] [...arguments]

  commands:
    sss - SSS Operations
    s4 - S4 Operations
`
		if actual != expected {
			t.Errorf("Unexpected help string\nactual:\n%v\nexpected:\n%v", actual, expected)
		}
	}

}
