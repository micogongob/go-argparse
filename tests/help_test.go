package tests

import (
	"github.com/micogongob/go-argparse/parse"
	"testing"
)

var HelpCommandAliases = []string{"help", "--help", "-h"}

func init() {
	parse.Setup(AppCode, AppDesc)
	parse.AddCommand(SssCommand)
	parse.AddCommand(S4Command)
}

func TestAppHelp(t *testing.T) {
	for i := 0; i <= len(HelpCommandAliases); i++ {
		var args []string
		if i == len(HelpCommandAliases) {
			args = []string{}
		} else {
			args = []string{HelpCommandAliases[i]}
		}

		parsedOutput, err := parse.ParseStrings(args)

		actual := parsedOutput.HelpMessage

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
