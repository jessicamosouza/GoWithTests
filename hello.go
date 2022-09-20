package main

import "fmt"

const portuguese = "Portuguese"
const french = "French"
const japanese = "Japanese"
const englishHelloPrefix = "Hello, "
const portugueseHelloPrefix = "Olá, "
const frenchHelloPrefix = "Bonjour, "
const japaneseHelloPrefix = "こんにちは, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"

	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
	case japanese:
		prefix = japaneseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return

}

func main() {
	fmt.Println(Hello("Chloé", "French"))
}
