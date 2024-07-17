package parse

import (
	"testing"
)

var TestApp App

func init() {
	sssCommand := NewCommand(SSS_CODE, "SSS Queue Operations")
	sssCommand.AddChildrenCommand("list-queues", "Lists SSS queues")
	sssCommand.AddChildrenCommand("send-message", "Send string message to SSS queue")

	s4Command := NewCommand(S4_CODE, "S4 Bucket Operations")
	s4Command.AddChildrenCommand("make-bucket", "Create S4 bucket")
	s4Command.AddChildrenCommand("copy-objects", "Copies object between s4 buckets")

	TestApp = NewApp(APP_CODE, APP_DESC, sssCommand, s4Command)
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

  usage: ows [command] [subcommand] [...parameters]

  commands:
    sss - SSS Queue Operations
    s4 - S4 Bucket Operations
    help - Show help. Alternatives: --help, -h
`
		if actual != expected {
			t.Errorf("index: %v - \nactual:\n%v\nexpected:\n%v", i, actual, expected)
		}
	}
}

func TestSssHelp(t *testing.T) {
	for i := 0; i <= len(HelpCommandAliases); i++ {
		var args []string
		if i == len(HelpCommandAliases) {
			args = []string{"sss"}
		} else {
			args = []string{"sss", HelpCommandAliases[i]}
		}

		parsedOutput, err := TestApp.parseStrings(args)

		actual := parsedOutput.helpMessage

		if err != nil {
			t.Errorf("Unexpected error. %v", err)
		}

		expected := `SSS Queue Operations

  usage: sss [subcommand] [...parameters]

  subcommands:
    list-queues - Lists SSS queues
    send-message - Send string message to SSS queue
`
		if actual != expected {
			t.Errorf("index: %v - \nactual:\n%v\nexpected:\n%v", i, actual, expected)
		}
	}
}

func TestS4Help(t *testing.T) {
	for i := 0; i <= len(HelpCommandAliases); i++ {
		var args []string
		if i == len(HelpCommandAliases) {
			args = []string{"s4"}
		} else {
			args = []string{"s4", HelpCommandAliases[i]}
		}

		parsedOutput, err := TestApp.parseStrings(args)

		actual := parsedOutput.helpMessage

		if err != nil {
			t.Errorf("Unexpected error. %v", err)
		}

		expected := `S4 Bucket Operations

  usage: s4 [subcommand] [...parameters]

  subcommands:
    make-bucket - Create S4 bucket
    copy-objects - Copies object between s4 buckets
`
		if actual != expected {
			t.Errorf("index: %v - \nactual:\n%v\nexpected:\n%v", i, actual, expected)
		}
	}
}
