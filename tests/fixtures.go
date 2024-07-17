package tests

import (
	"fmt"
	"github.com/micogongob/go-argparse/parse"
)

func newTestApp() parse.App {
	return parse.NewApp("Cli tool that helps you do stuff", []parse.Command{
		{
			Code:        "sqs",
			Triggers:    []string{"sqs"},
			Description: "Do SQS stuff",
			SubCommands: []parse.SubCommand{
				{
					Code:        "list-queues",
					Description: "Lists the SQS quues",
					Parameters: []parse.Parameter{
						{
							Code:        "queue-name",
							Description: "name of the queue",
						},
						{
							Code:        "json",
							Triggers:    []string{"-j", "-json"},
							Description: "Output in JSON",
							IsFlag:      true,
						},
					},
					OnCommand: func(pv map[string]string) error {
						return listQueues(pv)
					},
				},
			},
		},
		{
			Code:        "s3",
			Description: "Do S3 Bucket stuff",
			SubCommands: []parse.SubCommand{
				{
					Code:        "make-bucket",
					Description: "Creates S3 bucket",
					Parameters: []parse.Parameter{
						{
							Code:        "bucket-name",
							Description: "Name of the S3 bucket to create",
						},
						{
							Code:        "type",
							Description: "standard/infrequent_access",
						},
					},
					OnCommand: func(pv map[string]string) error {
						return makeBucket(pv)
					},
				},
			},
		},
	})
}

func listQueues(paramValues map[string]string) error {
	fmt.Println("Listing queues")
	return nil
}

func makeBucket(paramValues map[string]string) error {
	fmt.Println("Making bucket")
	return nil
}
