package main

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestMain(m *testing.M) {
	log.SetLevel(log.DebugLevel)
	m.Run()
}

func TestValidSentences(t *testing.T) {
	validSentences := []string{
		"The quick brown fox said \"hello Mr lazy dog\".",
		"The quick brown fox said hello Mr lazy dog.",
		"One lazy dog is too few, 13 is too many.",
		"One lazy dog is too few, thirteen is too many.",
		"How many \"lazy dogs\" are there?",
	}

	for _, sentence := range validSentences {
		valid := validateSentence(sentence)
		if !valid {
			t.Errorf("Sentence '%s' should be valid, but got %v", sentence, valid)
		}
	}
}

func TestInvalidSentences(t *testing.T) {
	invalidSentences := []string{
		"The quick brown fox said \"hello Mr. lazy dog\".",
		"the quick brown fox said \"hello Mr lazy dog\".",
		"\"The quick brown fox said \"hello Mr lazy dog.\"",
		"One lazy dog is too few, 12 is too many.",
		"Are there 11, 12, or 13 lazy dogs?",
	}

	for _, sentence := range invalidSentences {
		valid := validateSentence(sentence)
		if valid {
			t.Errorf("Sentence '%s' should be invalid, but got %v", sentence, valid)
		}
	}
}

// Test startsWithUppercase
func TestValidStartsWithUppercase(t *testing.T) {
	validSentence := "\"The quick brown fox said 'hello Mr lazy dog'.\""
	valid := startsWithUppercase(validSentence)
	if !valid {
		t.Errorf("The sentence '%s' should be valid, but got %v", validSentence, valid)
	}
}

func TestInvalidStartsWithUppercase(t *testing.T) {
	invalidSentence := "\"the quick brown fox said 'hello Mr lazy dog'.\""
	valid := startsWithUppercase(invalidSentence)
	if valid {
		t.Errorf("The sentence '%s' should be invalid, but got %v", invalidSentence, valid)
	}
}

// Test hasEqualQuotes
func TestValidHasEqualQuotes(t *testing.T) {
	validSentence := "The quick brown fox had a \"few\" sets of 'quotes' \"Hello Mr lazy dog\"."
	valid := hasEqualQuotes(validSentence)
	if !valid {
		t.Errorf("The sentence '%s' should be valid, but got %v", validSentence, valid)
	}
}

func TestInvalidHasEqualQuotes(t *testing.T) {
	invalidSentence := "The quick brown fox had a \"few sets of 'quotes' \"Hello Mr lazy dog\""
	valid := hasEqualQuotes(invalidSentence)
	if valid {
		t.Errorf("The sentence '%s' should be invalid, but got %v", invalidSentence, valid)
	}
}

// Test hasSpelledOutNumbers
func TestValidHasSpelledOutNumbers(t *testing.T) {
	validSentence := "Our lazy dog is too few, 13 is too many, 14, 15, 16, is also too many."
	valid := hasSpelledOutNumbers(validSentence, SPELLED_NUMBERS_UP_TO)
	if !valid {
		t.Errorf("The sentence '%s' should be valid, but got %v", validSentence, valid)
	}
}

func TestInvalidHasSpelledOutNumbers(t *testing.T) {
	invalidSentence := "Our lazy dog is too few, 12 is too many, 4, 5, 6, is also too many"
	valid := hasSpelledOutNumbers(invalidSentence, SPELLED_NUMBERS_UP_TO)
	if valid {
		t.Errorf("The sentence '%s' should be invalid, but got %v", invalidSentence, valid)
	}
}
