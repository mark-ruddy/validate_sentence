package main

import (
	"fmt"
)

func main() {
	testSentence := "The quick brown fox said \"Hello Mr lazy dog\""
	valid, err := validateSentence(testSentence)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The sentence '%s' is %v\n", testSentence, valid)
}
