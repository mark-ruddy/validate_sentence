package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func main() {
	debugLevelFlag := flag.Bool("debug", false, "Enable debug level logs")
	sentenceFlag := flag.String("sentence", "", "The sentence to validate")
	flag.Parse()

	log.SetLevel(log.InfoLevel)
	if *debugLevelFlag {
		log.SetLevel(log.DebugLevel)
	}

	sentence := *sentenceFlag
	if sentence == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a sentence: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		sentence = strings.TrimSpace(input)
	}

	valid := validateSentence(sentence)
	if valid {
		fmt.Printf("Valid sentence: '%s'\n", sentence)
	} else {
		fmt.Printf("Invalid sentence: '%s'\n", sentence)
	}
	return
}
