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

// Update replaces the definition of a word already in dictionary
func (d Dictionary) Update(word, newDefinition string) error {
	_, alreadyInserted := d[word]
	if alreadyInserted {
		d[word] = newDefinition
		return nil
	}
	return ErrWordNotFound
}

// Delete removes word from dict
func (d Dictionary) Delete(word string) error {
	_, alreadyInserted := d[word]
	if alreadyInserted {
		delete(d, word)
		return nil
	}
	return ErrWordNotFound
}

func main() {
	myDict := Dictionary(map[string]string{"a": "b"})
	fmt.Println(myDict["a"])
}
