package tests

import (
	"github.com/micogongob/go-argparse/parse"
)

var (
	AppCode = "ows"
	AppDesc = "Owsome cli"
)

var (
	SssCommand = parse.Command{
		Code:        "sss",
		Description: "SSS Operations",
	}
	S4Command = parse.Command{
		Code:        "s4",
		Description: "S4 Operations",
	}
)
