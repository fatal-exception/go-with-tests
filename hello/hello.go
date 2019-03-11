package main

import (
	"fmt"
	"strings"
)

// Hello is a useless function
func Hello(name string, language string) string {
	helloPrefixes := map[string]string{}
	helloPrefixes["english"] = "Hello"
	helloPrefixes["spanish"] = "Hola"
	helloPrefixes["french"] = "Bonjour"

	if name == "" {
		name = "World"
	}

	language = strings.ToLower(language)
	_, ok := helloPrefixes[language]

	if !ok || language == "" {
		language = "english"
	}
	return helloPrefixes[language] + ", " + name
}

func main() {
	fmt.Println(Hello("Matt", ""))
}
