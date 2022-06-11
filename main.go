package main

import "fmt"

func main() {
    fmt.Println(Hello("togamin"))
}

func Hello(name string) string {
	return "Hello, " + name
}