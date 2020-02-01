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
		var chars []string = strings.Split(words[i], "")

		for i := 0; i < len(chars); i++ {
			if i%2 == 0 {
				chars[i] = strings.ToUpper(chars[i])
			} else {
				chars[i] = strings.ToLower(chars[i])
			}
		}

		words[i] = strings.Join(chars, "")
	}

	return strings.Join(words, " ")
}
