package main

import (
	"testing"
)

func TestHello(t *testing.T){

	assertCorrectMessage := func(t *testing.T, got string, want string) {
		t.Helper()
		
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("Say hello to an individual", func(t *testing.T){
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello world when an empty string is supplied", func(t *testing.T){
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In Spanish", func(t *testing.T){
		got := Hello("Elodie", "Spanish")
		want := "Holda, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In French", func(t *testing.T){
		got := Hello("Billiam", "French")
		want := "Bonjour, Billiam"
		assertCorrectMessage(t, got, want)
	})
}

