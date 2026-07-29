package hello

import "fmt"

const (
	spanish = "Spanish"
	french = "French"
	jap = "Japanese"
	
	spanishHelloPrefix = "Hola, "
	englishHelloPrefix = "Hello, "
	frenchHelloPrefix = "Bonjour, "
	japHelloPrefix = "Ohayo, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case jap:
		prefix = japHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}
func main() {
	fmt.Println(Hello("Joe", ""))
}
