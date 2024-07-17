package tests

import (
	"testing"
)


func TestAppHelpEmpty(t *testing.T) {
	app := newTestApp()

	actual, err := app.ParseStrings([]string{})

	if err != nil {
		t.Errorf("Unexpected error. %v", err)
	}

	expected := `
Cli tool that helps you do stuff
	`

	if actual != expected {
		t.Errorf("Unexpected help string. %v", actual)
	}
}
