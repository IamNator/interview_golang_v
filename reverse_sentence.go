package revsentence

import (
	"regexp"
	"strings"
)

//ReverseSentence reverses the words in a given sentence
func ReverseSentence(sentence string) (reversedSentence string) {

	if sentence == "" { //return sentence if empty
		return sentence
	}

	words := regexp.MustCompile(`\w+`).FindAllString(sentence, -1) //split the sentence into a slice of words
	for _, word := range words {
		reversedSentence = word + " " + reversedSentence
	}

	return strings.TrimSpace(reversedSentence)
}
