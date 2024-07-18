package parse

import "testing"

func TestAppHelp(t *testing.T) {
	// given
	testApp := newTestApp(t)

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
		assertNilError(t, err)
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
	testApp := newTestApp(t)

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
		assertNilError(t, err)
		assertStringEquals(t, parsedOutput.helpMessage, `SSS Queue Operations

  usage: sss [subcommand] [...parameters]

  subcommands:
    version      -> Show SSS version
    list-queues  -> Lists SSS queues
    send-message -> Send string message to SSS queue
    help         -> Show help. Alternatives: --help, -h
`)
	}
}

func TestS4Help(t *testing.T) {
	// tiven
	testApp := newTestApp(t)

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
		assertNilError(t, err)
		assertStringEquals(t, parsedOutput.helpMessage, `S4 Bucket Operations

  usage: s4 [subcommand] [...parameters]

  subcommands:
    make-bucket  -> Create S4 bucket
    copy-objects -> Copies object between s4 buckets
    help         -> Show help. Alternatives: --help, -h
`)
	}
}

func TestSssVersionHelp(t *testing.T) {
	// given
	testApp := newTestApp(t)

	for i := 0; i < len(HelpCommandAliases); i++ {
		args := []string{SSS_CODE, "version", HelpCommandAliases[i]}
		t.Logf("Args: %v", args)

		// when
		parsedOutput, err := testApp.parseStrings(args)

		// then
		assertNilError(t, err)
		assertStringEquals(t, parsedOutput.helpMessage, `Show SSS version

  usage: version
`)
	}
}

func TestSssListQueuesHelp(t *testing.T) {
	// given
	testApp := newTestApp(t)

	for i := 0; i < len(HelpCommandAliases); i++ {
		args := []string{SSS_CODE, "list-queues", HelpCommandAliases[i]}
		t.Logf("Args: %v", args)

		// when
		parsedOutput, err := testApp.parseStrings(args)

		// then
		assertNilError(t, err)
		assertStringEquals(t, parsedOutput.helpMessage, `Lists SSS queues

  usage: list-queues [...parameters]

  parameters:
    --page-size -> pagination (optional)
    --debug     -> DEBUG logging. Flag (optional)
`)
	}
}

func TestSssSendMessageHelp(t *testing.T) {
	// given
	testApp := newTestApp(t)

	for i := 0; i < len(HelpCommandAliases); i++ {
		args := []string{SSS_CODE, "send-message", HelpCommandAliases[i]}
		t.Logf("Args: %v", args)

		// when
		parsedOutput, err := testApp.parseStrings(args)

		// then
		assertNilError(t, err)
		assertStringEquals(t, parsedOutput.helpMessage, `Send string message to SSS queue

  usage: send-message [...parameters]

  parameters:
    --queue-name -> the name of the SSS queue (required)
`)
	}
}
