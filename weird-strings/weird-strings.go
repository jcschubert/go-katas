package weirdstrings

import "strings"

// toWeirdCase accepts a string, and returns the same
// string with all even index characters in each word upper cased
// and all odd strings in each word lower cased
func toWeirdCase(str string) string {
	return strings.ToUpper(str)
}
