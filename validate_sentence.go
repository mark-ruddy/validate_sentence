package main

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	log "github.com/sirupsen/logrus"
)

const SPELLED_NUMBERS_UP_TO = 13

func validateSentence(sentence string) bool {
	if len(sentence) < 1 {
		log.Debugf("Sentence is empty: '%s'", sentence)
		return false
	}
	if !startsWithUppercase(sentence) {
		log.Debugf("Sentence does not start with an uppercase letter: '%s'", sentence)
		return false
	}

	if !hasEqualQuotes(sentence) {
		log.Debugf("Sentence has unequal quotes: '%s'", sentence)
		return false
	}

	if !strings.HasSuffix(sentence, ".") && !strings.HasSuffix(sentence, "?") && !strings.HasSuffix(sentence, "!") {
		log.Debugf("Sentence does not end with a punctuation mark: '%s'", sentence)
		return false
	}

	if strings.Count(sentence, ".") > 1 {
		log.Debugf("Sentence has more than one period: '%s'", sentence)
		return false
	}

	if !hasSpelledOutNumbers(sentence, SPELLED_NUMBERS_UP_TO) {
		log.Debugf("Sentence does not contain spelled out numbers: '%s'", sentence)
		return false
	}
	return true
}

func startsWithUppercase(sentence string) bool {
	offset := 0 // offset if sentence starts with a quote
	if sentence[0] == '"' || sentence[0] == '\'' {
		offset = 1
		if len(sentence) < 2 {
			log.Debugf("Sentence is too short: '%s'", sentence)
			return false
		}
	}
	firstRune, _ := utf8.DecodeRuneInString(sentence[offset:])
	if !unicode.IsUpper(firstRune) {
		return false
	}
	return true
}

func hasEqualQuotes(sentence string) bool {
	singleQuoteCount := strings.Count(sentence, `'`)
	doubleQuoteCount := strings.Count(sentence, `"`)
	if singleQuoteCount%2 != 0 || doubleQuoteCount%2 != 0 {
		return false
	}
	return true
}

func hasSpelledOutNumbers(sentence string, spelledUpTo int) bool {
	numRegex := regexp.MustCompile(`\d+`)
	digitWords := numRegex.FindAllString(sentence, -1) // match all digits
	for _, digitWord := range digitWords {
		num, err := strconv.Atoi(digitWord)
		if err != nil {
			continue
		}
		if num < spelledUpTo {
			return false
		}
	}
	return true
}
