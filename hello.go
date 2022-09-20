package main

import "fmt"

const portuguese = "Portuguese"
const french = "French"
const englishHelloPrefix = "Hello, "
const portugueseHelloPrefix = "Olá, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"

	}
	if language == portuguese {
		return portugueseHelloPrefix + name

	}
	if language == french {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Chloé", "French"))
}
