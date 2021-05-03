package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
		assertNoError(t, err)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		assertError(t, err, ErrUnknownWord)
	})
}

func TestAdd(t *testing.T) {
	// var m map[string]string <- be able to runtime panic
	// var m = map[string]string{} better
	// varm = make(map[string]string) better

	t.Run("new  word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a word"

		err := dictionary.Add(word, definition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing  word", func(t *testing.T) {
		word := "test"
		definition := "this is just a word"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	newDef := "update word"
	dictionary := Dictionary{word: "old word"}

	err := dictionary.Update(word, newDef)

	assertDefinition(t, dictionary, word, newDef)
	assertNoError(t, err)
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test  def"}
	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrUnknownWord {
		t.Errorf("Expected %q to be deleted", word)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	if got == nil {
		if want == nil {
			return
		}
		t.Fatal("expecetd to get an error.")
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertDefinition(t *testing.T, dic Dictionary, word, def string) {
	t.Helper()

	got, err := dic.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if def != got {
		t.Errorf("got %q want %q", got, def)
	}
}
