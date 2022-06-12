package main

import "fmt"

const helloPrefix = "Hello, "
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
	if language == japanese {
		return janpanesePrefix + name
	}
	if language == french {
		return frenchPrefix + name
	}
	return helloPrefix + name
}