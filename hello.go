package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	// declare variable :=
	prefix := englishHelloPrefix

	switch language {
	case "Spanish":
		if name == "World" {
			name = "Mundo"
		}
		prefix = spanishHelloPrefix
	case "French":
		if name == "World" {
			name = "Monde"
		}
		prefix = frenchHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
