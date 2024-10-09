package main

func validateSentence(sentence string) (bool, error) {
	if len(sentence) < 1 {
		return false, nil
	}
	return true, nil
}
