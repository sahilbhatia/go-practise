package greetings

import (
	"testing"
)

func TestGreetWithNameWithNameWithoutSpaces(t *testing.T) {
	name := "Sahil"
	want := "Hi Sahil, welcome to your first Go program!"
	got, _ := GreetWithName(name)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestGreetWithNameWithNameWithSpacesPrefixed(t *testing.T) {
	name := "      Sahil"
	want := "Hi Sahil, welcome to your first Go program!"
	got, _ := GreetWithName(name)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestGreetWithNameWithNameWithSpacesSuffixed(t *testing.T) {
	name := "Sahil      "
	want := "Hi Sahil, welcome to your first Go program!"
	got, _ := GreetWithName(name)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestGreetWithNameWithNameWithSpacesPrefixedAndSuffixed(t *testing.T) {
	name := "      Sahil      "
	want := "Hi Sahil, welcome to your first Go program!"
	got, _ := GreetWithName(name)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestGreetWithNameWithEmptyString(t *testing.T) {
	want := "name argument cannot be left blank"
	_, got := GreetWithName("")

	if got == nil {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestIsEmptyStringFalse(t *testing.T) {
	want := false
	got := isEmptyString("      Sahil      ")

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestIsEmptyStringTrue(t *testing.T) {
	want := true
	got := isEmptyString("            ")

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
