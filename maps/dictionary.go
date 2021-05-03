package maps

import "errors"

type Dictionary map[string]string

var ErrUnknownWord = errors.New("could not find the word you were looing for")

func (d Dictionary) Search(word string) string {
	return d[word]
}
