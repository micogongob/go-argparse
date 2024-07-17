package parse

type App struct {
	Description string
	Commands    []Command
	HelpCommand Command
}

type Command struct {
	Code        string
	Description string
	Triggers    []string
	SubCommands []SubCommand
	HelpCommand SubCommand
	OnCommand   func(paramValues map[string]string) error
}

type SubCommand struct {
	Code        string
	Description string
	Triggers    []string
	Parameters  []Parameter
	OnCommand   func(paramValues map[string]string) error
}

type Parameter struct {
	Code        string
	Description string
	Triggers    []string
	Optional    bool
	IsFlag      bool
}

type HelpInfo struct {
	code        *string
	description string
	subHelpInfo []HelpInfo
}

type HelpInfoProvider interface {
	helpInfo() HelpInfo
}
