package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

const spanish = "Spanish"
const french = "French"

func Hello(name string, language string) string {

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	}

	if name == "" {
		name = "world"

	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
