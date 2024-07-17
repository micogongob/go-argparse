package parse

type app struct {
	code        string
	description string
	commands    []Command
}

type Command struct {
	Code        string
	Description string
	Aliases     []string
	Children    []Command
}

type helpInfo struct {
	code         string
	description  string
	usageSuffix  string
	childrenName string
	children     []helpInfo
}

type ParseOutput struct {
	HelpMessage string
}
