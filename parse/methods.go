package parse

func NewApp(code, description string) App {
	return App{
		code: code,
		description: description,
		commands: []Command{HelpCommand},
	}
}

func (app *App) AddCommand(code, description string, aliases []string, children []Command) {
	app.commands = append(app.commands, Command{
		code: code,
		description: description,
		aliases: aliases,
		children: children,
	})
}