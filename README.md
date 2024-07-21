# go-argparse

- A CLI parsing library inspired by AWS CLI v1
- There's already the **flag** go package and others. Why reinvent the wheel? obviously, for learning purposes :P


### Usage

```bash
go get github.com/micov6/go-argparse@v0.1.0
```

see [releases](https://github.com/micov6/go-argparse/releases) for latest version
<br>

```golang
package main

import (
	"fmt"
	"github.com/micov6/go-argparse"
)

func main() {
	app := parse.App{
		Code: "Cli",
		Description: "A cli tool",
		Commands: []*parse.Command{
			{
				Code: "greet",
				Description: "Contains various greetings",
				Children: []*parse.ChildCommand{
					{
						Code: "hello",
						CommandHandler: hello,
						Parameters: []*parse.Parameter{
							{
								Code: "name",
								Description: "Name of the person to greet",
							},
						},
					},
					{
						Code: "spanish-hello",
						CommandHandler: spanishHello,
						Parameters: []*parse.Parameter{
							{
								Code: "name",
								Description: "Name of the person to greet",
							},
						},
					},
				},
			},
		},
	}

	err := app.Parse()

	if err != nil {
		panic(err)
	}
}

func hello(parameterValues map[string]parse.ParameterValue) error {
	fmt.Printf("hello, %v\n", parameterValues["name"].StringValue)
	return nil
}

func spanishHello(parameterValues map[string]parse.ParameterValue) error {
	fmt.Printf("hola, como estas, %v\n", parameterValues["name"].StringValue)
	return nil
}
```

where **app.Parse()** will parse the **os.Args** based on commands, childCommands, and parameters you've defined
<br>

example:
```bash
go run main.go greet hello --name John
// output: hello, John
```
<br>

or:
```bash
go run main.go greet spanish-hello --name Juan
// output: hola, como estas, Juan
```
<br>

see the [*_test.go](./parse) for more examples

### Contributing

see [CONTRIBUTING.md](/CONTRIBUTING.md)
