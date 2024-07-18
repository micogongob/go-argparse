package parse

import "testing"

const (
	APP_CODE = "ows"
	APP_DESC = "Owsome cli"
	SSS_CODE = "sss"
	S4_CODE  = "s4"
)

var HelpCommandAliases = []string{"help", "--help", "-h"}

func newTestApp() App {
	sssCommand := NewCommand(SSS_CODE, "SSS Queue Operations")
	sssCommand.AddChildCommand("list-queues", "Lists SSS queues",
		NewCommandParameter("region", "the region of the SSS queues", false, false),
		NewCommandParameter("page-size", "pagination", true, false),
		NewCommandParameter("debug", "DEBUG logging", true, true))
	sssCommand.AddChildCommand("send-message", "Send string message to SSS queue")

	s4Command := NewCommand(S4_CODE, "S4 Bucket Operations")
	s4Command.AddChildCommand("make-bucket", "Create S4 bucket")
	s4Command.AddChildCommand("copy-objects", "Copies object between s4 buckets")

	return NewApp(APP_CODE, APP_DESC, sssCommand, s4Command)
}

func assertNonNil(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Unexpected error encountered: %v", err)
	}
}

func assertStringEquals(t *testing.T, actual, expected string) {
	if actual != expected {
		t.Errorf("Strings does not match\nactual:\n%v\nexpected:\n%v", actual, expected)
	}
}
