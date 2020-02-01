package splitstrings

// Solution splits the string into pais of two characters. If the string contains
// an odd number of characters then it the missing second character of the pair
// is replaced with an underscore ("_")
//
// Examples:
// Solution("abc") // should return ["ab", "c_"]
// Solution("abcd") // should return ["ab", "cd"]
//
func Solution(str string) []string {
	result := []string{}

	item := str

	if len(str) < 2 {
		item = item + "_"
	}

	result = append(result, item)

	return result
}
