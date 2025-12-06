package main

import "testing"

func TestHello(t *testing.T){
	got:= Hello()
	want:= "Hello, world"

	assertCorrectMessage(t, got, want)
}

func TestHelloWithName(t *testing.T){
	t.Run("saying hello to people", func(t *testing.T){
		got:=HelloName("Chris")
		want:= "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T){
		got := HelloName("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

}

func TestHelloWithLanguage(t *testing.T){
	t.Run("in English", func(t *testing.T){
		got := HelloWithLanguage("Sara", "")
		want := "Hello, Sara"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T){
		got := HelloWithLanguage("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T){
		got := HelloWithLanguage("Anna", "French")
		want := "Bonjour, Anna"
		assertCorrectMessage(t, got, want)
	})
}

// "got, want string"" is shortform of "got string, want string"
func assertCorrectMessage(t testing.TB, got, want string){
	// tells test suite that it's a helper function
	// if fails, reported line number will be in function call not in helper function
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
