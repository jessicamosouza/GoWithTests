package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"

	}
	if language == "Portuguese" {
		return "Olá, " + name

	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Jéssica", "Portuguese"))
}
