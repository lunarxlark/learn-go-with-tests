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
		got, err := dictionary.Search("unknown")
		want := "could not find the word you were looing for"

		if err == nil {
			t.Fatal("expected to get an error")
		}
		assertStrings(t, err.Error(), want)
		assertError(t, err, ErrUnknownWord)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("there is no error")
	}
	if got != want {
		t.Error("got %q want %q", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error but didn't want one")
	}
}
