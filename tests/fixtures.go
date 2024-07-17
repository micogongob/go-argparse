package tests

import (
	"github.com/micogongob/go-argparse/parse"
)

var (
	APP_CODE = "ows"
	APP_DESC = "Owsome cli"
)

var (
	SSS_COMMAND = parse.Command{
		Code:        "sss",
		Description: "SSS Operations",
	}
	S4_COMMAND = parse.Command{
		Code:        "s4",
		Description: "S4 Operations",
	}
)
