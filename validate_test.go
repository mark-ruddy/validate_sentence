package main

import (
	"testing"
)

func TestValidSentence(t *testing.T) {
	validSentence := "The quick brown fox said \"Hello Mr lazy dog\""
	valid, err := validateSentence(validSentence)
	if err != nil {
		t.Fatal(err)
	}
	if !valid {
		t.Fatalf("The sentence '%s' is not valid", validSentence)
	}
}

func TestInvalidSentence(t *testing.T) {
	invalidSentence := ""
	valid, err := validateSentence(invalidSentence)
	if err != nil {
		t.Fatal(err)
	}
	if valid {
		t.Fatalf("The sentence '%s' is valid", invalidSentence)
	}
}
