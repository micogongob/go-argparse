package parse

import "testing"

func TestAppHelp(t *testing.T) {
	// given
	testApp := newTestApp()

	for i := 0; i <= len(HelpCommandAliases); i++ {
		var args []string
		if i == len(HelpCommandAliases) {
			args = []string{}
		} else {
			args = []string{HelpCommandAliases[i]}
		}
		t.Logf("Args: %v", args)

		// when
		parsedOutput, err := testApp.parseStrings(args)

		// then
		assertNonNil(t, err)
		assertStringEquals(t, parsedOutput.helpMessage, `Owsome cli

  usage: ows [command] [subcommand] [...parameters]

  commands:
    sss  -> SSS Queue Operations
    s4   -> S4 Bucket Operations
    help -> Show help. Alternatives: --help, -h
`)
	}
}

func TestSssHelp(t *testing.T) {
	// given
	testApp := newTestApp()

	for i := 0; i <= len(HelpCommandAliases); i++ {
		var args []string
		if i == len(HelpCommandAliases) {
			args = []string{SSS_CODE}
		} else {
			args = []string{SSS_CODE, HelpCommandAliases[i]}
		}
		t.Logf("Args: %v", args)

		// when
		parsedOutput, err := testApp.parseStrings(args)

		// then
		assertNonNil(t, err)
		assertStringEquals(t, parsedOutput.helpMessage, `SSS Queue Operations

  usage: sss [subcommand] [...parameters]

  subcommands:
    list-queues  -> Lists SSS queues
    send-message -> Send string message to SSS queue
    help         -> Show help. Alternatives: --help, -h
`)
	}
}

func TestS4Help(t *testing.T) {
	// tiven
	testApp := newTestApp()

	for i := 0; i <= len(HelpCommandAliases); i++ {
		var args []string
		if i == len(HelpCommandAliases) {
			args = []string{S4_CODE}
		} else {
			args = []string{S4_CODE, HelpCommandAliases[i]}
		}
		t.Logf("Args: %v", args)

		parsedOutput, err := testApp.parseStrings(args)

		// then
		assertNonNil(t, err)
		assertStringEquals(t, parsedOutput.helpMessage, `S4 Bucket Operations

  usage: s4 [subcommand] [...parameters]

  subcommands:
    make-bucket  -> Create S4 bucket
    copy-objects -> Copies object between s4 buckets
    help         -> Show help. Alternatives: --help, -h
`)
	}
}

func TestSssListQueuesHelp(t *testing.T) {
	// given
	testApp := newTestApp()

	for i := 0; i < len(HelpCommandAliases); i++ {
		args := []string{SSS_CODE, "list-queues", HelpCommandAliases[i]}
		t.Logf("Args: %v", args)

		// when
		parsedOutput, err := testApp.parseStrings(args)

		// then
		assertNonNil(t, err)
		assertStringEquals(t, parsedOutput.helpMessage, `Lists SSS queues

  usage: list-queues [...parameters]

  parameters:
    --region    -> the region of the SSS queues (required)
    --page-size -> pagination (optional)
    --debug     -> DEBUG logging. Flag (optional)
`)
	}
}
