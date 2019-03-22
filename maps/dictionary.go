package main

import (
	"errors"
	"fmt"
)

// ErrWordNotFound is for when word isn't in dict
var ErrWordNotFound = errors.New("Not Found")

// ErrWordExistsAlready used for when word is in dict already and insert attempted
var ErrWordExistsAlready = errors.New("Word already exists")

// Dictionary is
type Dictionary map[string]string

// Search is
func (d Dictionary) Search(word string) (string, error) {
	word, ok := d[word]
	if ok {
		return word, nil
	}
	return "", ErrWordNotFound
}

// Add adds word definition to dictionary
func (d Dictionary) Add(word, definition string) error {
	_, alreadyInserted := d[word]
	if alreadyInserted {
		return ErrWordExistsAlready
	}
	d[word] = definition
	return nil
}

func main() {
	myDict := Dictionary(map[string]string{"a": "b"})
	fmt.Println(myDict["a"])
}
