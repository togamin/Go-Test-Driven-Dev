package main

import "fmt"

const helloPrefix = "Hello, "

func main() {
    fmt.Println(Hello("togamin"))
}

func Hello(name string) string {
	if name == ""{
		name = "World"
	}
	return helloPrefix + name
}