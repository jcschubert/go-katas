package spinwords

import (
	"strings"
)

// SpinWords takes in a string of one or more words,
// and returns the same string, but with all five
// or more letter words reversed.
func SpinWords(sentence string) string {
	words := strings.Split(sentence, " ")

	for i := 0; i < len(words); i++ {
		if len(words[i]) >= 5 {
			words[i] = reverse(words[i])
		}
	}

	return strings.Join(words, " ")
}

func reverse(word string) string {
	runes := []rune(word)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
