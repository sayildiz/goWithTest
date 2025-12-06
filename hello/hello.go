package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const spanish = "Spanish"
const french = "French"

func HelloWithLanguage(name, language string) string{
	if name == ""{
		name = "World"
	}

	return greetingPrefix(language) + name
}

// private functions start with lowercase letter
// named return value (prefix string) will create var prefix and return it whenever with return instead return prefix
func greetingPrefix(language string) (prefix string){
	switch language {
	case spanish:
		prefix= spanishHelloPrefix
	case french:
		prefix= frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func HelloName(name string) string{
	if name == ""{
		name = "World"
	}
	return englishHelloPrefix + name
}

func Hello() string {
	return "Hello, world"
}

func main() {
	fmt.Println(Hello())
}
