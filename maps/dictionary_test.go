package main

import "testing"

func TestSearch(t *testing.T) {
	t.Run("Known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a test"}
		assertDefinition(t, dictionary, "test", "this is a test")
		_, err := dictionary.Search("test")
		assertError(t, err, nil)
	})
	t.Run("Unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a test"}

		_, err := dictionary.Search("bananas")
		assertError(t, err, ErrWordNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("Add new word to dict", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("man", "human male")
		_, err := dictionary.Search("man")
		assertDefinition(t, dictionary, "man", "human male")
		assertError(t, err, nil)
	})

	t.Run("Add existing word to dict", func(t *testing.T) {
		word := "earth"
		definition := "the ground"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, definition)
		assertError(t, err, ErrWordExistsAlready)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update with existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"
		dictionary.Update(word, newDefinition)

		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("Update with new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)
		assertError(t, err, ErrWordNotFound)

	})

}

func TestDelete(t *testing.T) {
	t.Run("delete word from dictionary", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		dictionary.Delete(word)
		_, err := dictionary.Search(word)
		assertError(t, err, ErrWordNotFound)

	})
}

func assertDefinition(t *testing.T, d Dictionary, word, definition string) {
	t.Helper()
	got, err := d.Search(word)
	if err != nil {
		t.Fatal("Should find word:", err)
	}

	if definition != got {
		t.Errorf("got '%s', want '%s'", got, definition)
	}
}

func assertError(t *testing.T, err error, want error) {
	t.Helper()
	if err != want {
		t.Errorf("Expected error to be '%v'. Was '%v'", want, err)
	}
}
