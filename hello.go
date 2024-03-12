package main

import (
	"fmt"
)

const (
	spanish            = "Spanish"
	french             = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = greetingSuffix(language)
	}
	return greetingPrefix(language) + name
}

// the second parenthesis is the named return value
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func greetingSuffix(language string) (name string) {
	switch language {
	case spanish:
		name = "Mundo"
	case french:
		name = "Monde"
	default:
		name = "World"
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
