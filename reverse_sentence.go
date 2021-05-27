package revsentence

import (
	"strings"
)

//ReverseSentence reverses the words in a given sentence
func ReverseSentence(sentence string) (reversedSentence string) {

	if sentence == "" { //return sentence if empty
		return sentence
	}

	words := strings.Split(sentence, " ") //split the sentence into a slice of words
	for _, word := range words {
		reversedSentence = word + " " + reversedSentence
	}

	return strings.TrimSpace(reversedSentence)
}
