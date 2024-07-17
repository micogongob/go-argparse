package parse

import (
	"testing"
)

var TestApp App

func init() {
	sssCommand := NewCommand(SSS_CODE, "SSS Queue Operations")
	sssCommand.AddChildCommand("list-queues", "Lists SSS queues",
		NewCommandParameter("queue-name", "the name of SSS queue", false, false),
		NewCommandParameter("page-size", "pagination", true, false),
		NewCommandParameter("debug", "DEBUG logging", true, true))
	sssCommand.AddChildCommand("send-message", "Send string message to SSS queue")

	s4Command := NewCommand(S4_CODE, "S4 Bucket Operations")
	s4Command.AddChildCommand("make-bucket", "Create S4 bucket")
	s4Command.AddChildCommand("copy-objects", "Copies object between s4 buckets")

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

		t.Logf("Args: %v", args)
		parsedOutput, err := TestApp.parseStrings(args)

		actual := parsedOutput.helpMessage

		if err != nil {
			t.Errorf("Unexpected error. %v", err)
		}

		expected := `Owsome cli

  usage: ows [command] [subcommand] [...parameters]

  commands:
    sss  -> SSS Queue Operations
    s4   -> S4 Bucket Operations
    help -> Show help. Alternatives: --help, -h
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
			args = []string{SSS_CODE}
		} else {
			args = []string{SSS_CODE, HelpCommandAliases[i]}
		}

		t.Logf("Args: %v", args)
		parsedOutput, err := TestApp.parseStrings(args)

		actual := parsedOutput.helpMessage

		if err != nil {
			t.Errorf("Unexpected error. %v", err)
		}

		expected := `SSS Queue Operations

  usage: sss [subcommand] [...parameters]

  subcommands:
    list-queues  -> Lists SSS queues
    send-message -> Send string message to SSS queue
    help         -> Show help. Alternatives: --help, -h
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
			args = []string{S4_CODE}
		} else {
			args = []string{S4_CODE, HelpCommandAliases[i]}
		}

		t.Logf("Args: %v", args)
		parsedOutput, err := TestApp.parseStrings(args)

		actual := parsedOutput.helpMessage

		if err != nil {
			t.Errorf("Unexpected error. %v", err)
		}

		expected := `S4 Bucket Operations

  usage: s4 [subcommand] [...parameters]

  subcommands:
    make-bucket  -> Create S4 bucket
    copy-objects -> Copies object between s4 buckets
    help         -> Show help. Alternatives: --help, -h
`
		if actual != expected {
			t.Errorf("index: %v - \nactual:\n%v\nexpected:\n%v", i, actual, expected)
		}
	}
}

func TestSssListQueuesHelp(t *testing.T) {
	for i := 0; i <= len(HelpCommandAliases); i++ {
		var args []string
		if i == len(HelpCommandAliases) {
			args = []string{SSS_CODE, "list-queues"}
		} else {
			args = []string{SSS_CODE, "list-queues", HelpCommandAliases[i]}
		}

		t.Logf("Args: %v", args)
		parsedOutput, err := TestApp.parseStrings(args)

		actual := parsedOutput.helpMessage

		if err != nil {
			t.Errorf("Unexpected error. %v", err)
		}

		expected := `Lists SSS queues

  usage: list-queues [...parameters]

  parameters:
    --queue-name -> the name of SSS queue (required)
    --page-size  -> pagination (optional)
    --debug      -> DEBUG logging. Flag (optional)
`
		if actual != expected {
			t.Errorf("index: %v - \nactual:\n%v\nexpected:\n%v", i, actual, expected)
		}
	}
}
