package main

import "fmt"

const (
	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
	spanish       = "Spanish"
)

func Hello(name, language string) string {
	if name == "" {
		return englishPrefix + "World"
	}

	if language == spanish {
		return spanishPrefix + name
	}
	return englishPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
