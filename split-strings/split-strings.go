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

	if len(str)%2 != 0 {
		str = str + "_"
	}

	for i := 0; i < len(str); i += 2 {
		result = append(result, str[i:i+2])
	}

	return result
}
