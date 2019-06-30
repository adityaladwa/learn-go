package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "This is just a test"}

	t.Run("Known Word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "This is just a test"
		assertStrings(t, got, want)
	})

	t.Run("Unknown Word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("Expected to get an error.")
		}
		assertStrings(t, err.Error(), ErrNotFound.Error())
	})
}

func TestAdd(t *testing.T) {
	t.Run("New Word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "This is a test")

		want := "This is a test"
		assertDefination(t, dictionary, "test", want)
	})

	t.Run("Existing Word", func(t *testing.T) {
		word := "test"
		defination := "this is just a test"
		dictionary := Dictionary{word: defination}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefination(t, dictionary, word, defination)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		defination := "this is just a text"
		dictionary := Dictionary{word: defination}
		newDefination := "This is new Defination"

		err := dictionary.Update(word, newDefination)

		assertError(t, err, nil)
		assertDefination(t, dictionary, word, newDefination)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		defination := "this is just a text"
		dictionary := Dictionary{word: defination}

		err := dictionary.Update(word, defination)

		assertError(t, err, nil)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test defination"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)

	if err != ErrNotFound {
		t.Errorf("Expected %s to be deleted", word)
	}
}

func assertDefination(t *testing.T, dictionary Dictionary, word, defination string) {
	t.Helper()
	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != defination {
		t.Errorf("Got '%s' want '%s'", got, defination)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got '%s' want '%s'", got, want)
	}
}

func assertError(t *testing.T, err, want error) {
	t.Helper()
	if err != want {
		t.Errorf("Got %s, want %s ", err.Error(), want.Error())
	}
}
