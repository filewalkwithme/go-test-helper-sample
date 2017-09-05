package hello

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	returned := HelloWorld()
	expected := "Hello, world!"

	compare(returned, expected, t)
}

func TestHelloWorldBroken(t *testing.T) {
	returned := HelloWorldBroken()
	expected := "Hello, world!"

	compare(returned, expected, t)
}

func compare(expected, returned string, t *testing.T) {
	// HERE! Mark as a HELPER function
	t.Helper()

	if returned != expected {
		t.Errorf("Returned: [%v], Expected: [%v]", returned, expected)
	}
}
