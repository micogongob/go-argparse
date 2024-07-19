package parse

import "testing"

func TestAppHelp(t *testing.T) {
	// given
	testApp := newTestApp(t)

	for i := 0; i <= len(HelpCommandaliases); i++ {
		var args []string
		if i == len(HelpCommandaliases) {
			args = []string{}
		} else {
			args = []string{HelpCommandaliases[i]}
		}
		t.Logf("Args: %v", args)

		// when
		parsedOutput, err := testApp.parseStrings(args)

		// then
		assertNilError(t, err)
		assertStringEquals(t, parsedOutput.helpMessage, `Owsome cli

  usage: ows [command] [subcommand] [...Parameters]

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

	for i := 0; i <= len(HelpCommandaliases); i++ {
		var args []string
		if i == len(HelpCommandaliases) {
			args = []string{SSS_CODE}
		} else {
			args = []string{SSS_CODE, HelpCommandaliases[i]}
		}
		t.Logf("Args: %v", args)

		// when
		parsedOutput, err := testApp.parseStrings(args)

		// then
		assertNilError(t, err)
		assertStringEquals(t, parsedOutput.helpMessage, `SSS Queue Operations

  usage: sss [subcommand] [...Parameters]

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

	for i := 0; i <= len(HelpCommandaliases); i++ {
		var args []string
		if i == len(HelpCommandaliases) {
			args = []string{S4_CODE}
		} else {
			args = []string{S4_CODE, HelpCommandaliases[i]}
		}
		t.Logf("Args: %v", args)

		parsedOutput, err := testApp.parseStrings(args)

		// then
		assertNilError(t, err)
		assertStringEquals(t, parsedOutput.helpMessage, `S4 Bucket Operations

  usage: s4 [subcommand] [...Parameters]

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

	for i := 0; i < len(HelpCommandaliases); i++ {
		args := []string{SSS_CODE, "version", HelpCommandaliases[i]}
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

	for i := 0; i < len(HelpCommandaliases); i++ {
		args := []string{SSS_CODE, "list-queues", HelpCommandaliases[i]}
		t.Logf("Args: %v", args)

		// when
		parsedOutput, err := testApp.parseStrings(args)

		// then
		assertNilError(t, err)
		assertStringEquals(t, parsedOutput.helpMessage, `Lists SSS queues

  usage: list-queues [...Parameters]

  parameters:
    --page-size -> pagination (optional)
    --debug     -> DEBUG logging. Flag (optional)
`)
	}
}

func TestSssSendMessageHelp(t *testing.T) {
	// given
	testApp := newTestApp(t)

	for i := 0; i < len(HelpCommandaliases); i++ {
		args := []string{SSS_CODE, "send-message", HelpCommandaliases[i]}
		t.Logf("Args: %v", args)

		// when
		parsedOutput, err := testApp.parseStrings(args)

		// then
		assertNilError(t, err)
		assertStringEquals(t, parsedOutput.helpMessage, `Send string message to SSS queue

  usage: send-message [...Parameters]

  parameters:
    --queue-url -> the url of the SSS queue (required)
    --debug     -> DEBUG logging. Flag (optional)
`)
	}
}
