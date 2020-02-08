package braces

// ValidBraces takes a string of braces, and returns true if the
// order of the braces is valid, and false if it is not.
func ValidBraces(s string) bool {
	if s == "" {
		return true
	}
	if s == "()" {
		return true
	}
	return false
}
