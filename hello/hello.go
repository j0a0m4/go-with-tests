package hello

const (
	englishPrefix    = "Hello, "
	spanishPrefix    = "Hola, "
	frenchPrefix     = "Bonjour, "
	portuguesePrefix = "Olá, "
	spanish          = "Spanish"
	french           = "French"
	portuguese       = "Portuguese"
)

func Hello(name, language string) string {
	if name == "" {
		return englishPrefix + "World"
	}

	return getPrefix(language) + name
}

func getPrefix(language string) string {
	switch language {
	case french:
		return frenchPrefix
	case spanish:
		return spanishPrefix
	case portuguese:
		return portuguesePrefix
	default:
		return englishPrefix
	}
}
