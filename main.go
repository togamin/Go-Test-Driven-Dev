package main

import "fmt"

const englishPrefix = "Hello, "
const japanese = "Japanese"
const janpanesePrefix = "Konnichiwa, "
const french = "French"
const frenchPrefix = "Bonjour, "

func main() {
    fmt.Println(Hello("togamin",""))
}

func Hello(name string, language string) string {
	if name == ""{
		name = "World"
	}

	prefix := englishPrefix

	switch language {
	case japanese :
		prefix = janpanesePrefix
	case french :
		prefix = frenchPrefix
	}
	return prefix + name
}