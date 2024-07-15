package main

import (
	"github.com/micogongob/go-argparse/internal"
	"github.com/micogongob/go-argparse/parse"
)

func main() {
	app := parse.NewApp("Cli tool that helps you do stuff", []parse.Command{
		{
			Code: "sqs",
			Triggers: []string{"sqs"},
			Description: "Do SQS stuff",
			SubCommands: []parse.SubCommand{
				{
					Code: "list-queues",
					Description: "Lists the SQS quues",
					Arguments: []parse.Argument{
						{
							Code: "queue-name",
							Description: "name of the queue",
						},
						{
							Code: "json",
							Triggers: []string{"-j", "--json"},
							Description: "Output in JSON",
							IsFlag: true,
						},
					},
				},
			},
		},
		{
			Code: "s3",
			Description: "Do S3 Bucket stuff",
			SubCommands: []parse.SubCommand{
				{
					Code: "make-bucket",
					Description: "Creates S3 bucket",
					Arguments: []parse.Argument{
						{
							Code: "bucket-name",
							Description: "Code of the S3 bucket to create",
						},
						{
							Code: "type",
							Description: "standard/infrequent_access",
						},
					},
				},
			},
		},
	})

	parsedCommand := app.Parse()
	switch true {
	case parsedCommand == nil:
		internal.Fail("Unknown command provided")
	case parsedCommand.Code == "sqs":
		handleSqs(parsedCommand)
	case parsedCommand.Code == "s3":
		handleS3(parsedCommand)
	}
}

func handleSqs(command *parse.Command) {

}

func handleS3(command *parse.Command) {

}
