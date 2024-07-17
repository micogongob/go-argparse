package parse

func (instance *App) helpInfo() HelpInfo {
	subHelpInfo := make([]HelpInfo, len(instance.Commands))

	for i, s := range instance.Commands {
		subHelpInfo[i] = s.helpInfo()
	}

	return HelpInfo{
		code:        nil,
		description: instance.Description,
		subHelpInfo: subHelpInfo,
	}
}

func (instance *Command) helpInfo() HelpInfo {
	subHelpInfo := make([]HelpInfo, len(instance.SubCommands))

	for i, s := range instance.SubCommands {
		subHelpInfo[i] = s.helpInfo()
	}

	return HelpInfo{
		code:        &instance.Code,
		description: instance.Description,
		subHelpInfo: subHelpInfo,
	}
}

func (instance *SubCommand) helpInfo() HelpInfo {
	subHelpInfo := make([]HelpInfo, len(instance.Parameters))

	for i, s := range instance.Parameters {
		subHelpInfo[i] = s.helpInfo()
	}

	return HelpInfo{
		code:        &instance.Code,
		description: instance.Description,
		subHelpInfo: subHelpInfo,
	}
}

func (instance *Parameter) helpInfo() HelpInfo {
	return HelpInfo{
		code:        &instance.Code,
		description: instance.Description,
		subHelpInfo: []HelpInfo{},
	}
}
