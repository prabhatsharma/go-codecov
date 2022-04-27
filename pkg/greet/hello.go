package greet

func Hello(user string, lang string) string {
	greeting := "Hello, "

	greeting = greetingPrefix(lang)

	if user == "" {
		return greeting + "World"
	}

	return greeting + user
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case "English":
		prefix = "Hello, "
	case "Spanish":
		prefix = "Hola, "
	case "French":
		prefix = "Bonjour, "
	case "German":
		prefix = "Guten Tag, "
	default:
		prefix = "Hello, "
	}
	return prefix
}
