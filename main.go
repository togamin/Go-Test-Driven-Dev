package main

import "fmt"

const helloPrefix = "Hello, "
const japanese = "Japanese"
const janpanesehelloPrefix = "Konnichiwa, "

func main() {
    fmt.Println(Hello("togamin",""))
}

func Hello(name string, language string) string {
	if name == ""{
		name = "World"
	}
	if language == japanese {
		return janpanesehelloPrefix + name
	}
	return helloPrefix + name
}