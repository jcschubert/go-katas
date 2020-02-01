package weirdstrings

import (
	"strings"
)

// toWeirdCase accepts a string, and returns the same
// string with all even index characters in each word upper cased
// and all odd strings in each word lower cased
func toWeirdCase(str string) string {
	words := strings.Split(str, " ")
	for i := 0; i < len(words); i++ {

		words[i] = wordToWeirdCase(words[i])
	}

	return strings.Join(words, " ")
}

// wordToWeirdCase turns a single word to WeIrD CaSe
func wordToWeirdCase(w string) string {
	var chars []string = strings.Split(w, "")

	for i := 0; i < len(chars); i++ {
		if i%2 == 0 {
			chars[i] = strings.ToUpper(chars[i])
		} else {
			chars[i] = strings.ToLower(chars[i])
		}
	}

	return strings.Join(chars, "")
}
