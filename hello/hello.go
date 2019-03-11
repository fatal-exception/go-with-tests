package main

import "fmt"

const englishPrefix = "Hello, "

// Hello is a useless function
func Hello(name string) string {
	if name == "" {
		return englishPrefix + "World"
	}
	return englishPrefix + name
}

func main() {
	fmt.Println(Hello("Matt"))
}
