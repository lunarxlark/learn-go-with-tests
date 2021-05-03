package maps

import "errors"

type Dictionary map[string]string

var ErrUnknownWord = errors.New("could not find the word you were looing for")

func (d Dictionary) Search(word string) (string, error) {
	if dic, found := d[word]; found {
		return dic, nil
	} else {
		return "", ErrUnknownWord
	}
}
