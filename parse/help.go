package parse

func (hp *app) Help() helpInfo {
	children := make([]helpInfo, len(hp.commands))

	for k, v := range hp.commands {
		children[k] = v.Help()
	}

	return helpInfo{
		code:         hp.code,
		description:  hp.description,
		usageSuffix:  "[command] [...arguments]",
		childrenName: "commands",
		children:     children,
	}
}

func (hp *Command) Help() helpInfo {
	return helpInfo{
		code:        hp.Code,
		description: hp.Description,
	}
}
