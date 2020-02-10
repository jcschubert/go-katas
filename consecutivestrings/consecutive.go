package consecutivestrings

func LongestConsec(s []string, k int) string {
	if len(s) == 0 || k > len(s) {
		return ""
	}

	result := ""

	for i := 0; i <= len(s)-k; i++ {
		currString := ""
		for n := i; n < i+k; n++ {
			currString = currString + s[n]
		}
		if len(currString) > len(result) {
			result = currString
		}
	}

	return result
}
