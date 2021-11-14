package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{
		"test": "this is just a test",
	}

	t.Run("known word", func(t *testing.T) {
		given := "test"
		got, err := dictionary.Search(given)
		want := "this is just a test"
		assertNoError(t, err)
		assertStrings(t, given, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		given := "unknown"
		_, err := dictionary.Search(given)
		want := ErrNotFound

		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "this is just a test")

		given := "test"
		want := "this is just a test"
		got, err := dictionary.Search(given)

		assertNoError(t, err)
		assertStrings(t, given, got, want)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, definition)
		want := ErrWordExists

		assertError(t, err, want)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertNoError(t, err)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	assertError(t, err, ErrNotFound)
}

func assertStrings(t *testing.T, given, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, given)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wants an error but didnt get one")
	}
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("Didn't want an error but got %q", err)
	}
}
