package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHolaPrefix = "Hola, "
const frenchBonjourPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		if name == "World" {
			name = "Mundo"
		}
		return spanishHolaPrefix + name
	} else if language == french {
		if name == "World" {
			name = "Monde"
		}
		return frenchBonjourPrefix + name
	}

	return englishHelloPrefix + name

}

func main() {
	fmt.Println(Hello("world", ""))
}
