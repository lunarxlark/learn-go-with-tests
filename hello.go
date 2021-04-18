package main

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	switch language {
	case "Spanish":
		return spanishHelloPrefix + name
	case "French":
		return frenchHelloPrefix + name
	default:
		return englishHelloPrefix + name
	}
}

//func main() {
//	name := ""
//	fmt.Println(Hello(name))
//}