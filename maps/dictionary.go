package maps

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	if dic, found := d[word]; found {
		return dic, nil
	} else {
		return "", ErrUnknownWord
	}
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrUnknownWord:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

type DictionaryErr string

const (
	ErrUnknownWord   = DictionaryErr("could not find the word you were looing for")
	ErrWordExists    = DictionaryErr("word is already exists")
	ErrWordNonExists = DictionaryErr("word didn't exists")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Update(word, newDef string) error {
	_, err := d.Search(word)
	switch err {
	case ErrUnknownWord:
		return ErrWordNonExists
	case nil:
		d[word] = newDef
	default:
		return err
	}
	return nil
}
