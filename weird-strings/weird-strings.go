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
	chars := strings.Split(w, "")

	for i := 0; i < len(chars); i++ {
		chars[i] = transformChar(i)(chars[i])
	}

	return strings.Join(chars, "")
}

func transformChar(i int) func(string) string {
	if i%2 == 0 {
		return strings.ToUpper
	}
	return strings.ToLower
}
