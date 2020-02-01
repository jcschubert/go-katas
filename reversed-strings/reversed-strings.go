package reversedstrings

// Solution reverses strings.
// Example: "Hello!" -> "!olleH"
func Solution(word string) (reversed string) {
	for i := len(word) - 1; i >= 0; i-- {
		reversed += string(word[i])
	}
	return
}
