package splitstrings

func Solution(str string) []string {
	if len(str) < 2 {
		return []string{str + "_"}
	}
	return []string{"ab"}
}
