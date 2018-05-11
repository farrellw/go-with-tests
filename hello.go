package main

import (
	"fmt"
)

const helloPrefix = "Hello, "
const spanishPrefix = "Holda, "
const frenchPrefix = "Bonjour, "
const spanish = "Spanish"
const french = "French"

func greetingPrefix(language string) (prefix string){
	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = helloPrefix
	}
	return
}

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	
	return greetingPrefix(language) + name
}

func main() {
	name := ""
	language := ""

	fmt.Println(Hello(name, language))
}
