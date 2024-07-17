package parse

func tail(source []string) []string {
	if len(source) <= 0 {
		return []string{}
	}
	return source[1:]
}

func commandMatchesArg(code string, aliases []string, argValue string) bool {
	if argValue == code {
		return true
	}

	for _, alias := range aliases {
		if argValue == alias {
			return true
		}
	}

	return false
}
