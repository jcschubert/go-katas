package braces

import (
	"strings"
)

type stack []string

func (s stack) Push(value string) stack {
	return append(s, value)
}

func (s stack) Pop() (stack, string) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) Has() bool {
	return len(s) > 0
}

func (s stack) Peek() string {
	l := len(s)
	return s[l-1]
}

const (
	openingBraces = "({["
	closingBraces = "]})"
)

var matches = map[string]string{
	")": "(",
	"}": "{",
	"]": "[",
}

// ValidBraces takes a string of braces, and returns true if the
// order of the braces is valid, and false if it is not.
func ValidBraces(s string) bool {
	tokens := strings.Split(s, "")
	stack := stack{}

	for _, token := range tokens {
		if strings.Contains(openingBraces, token) {
			stack = stack.Push(token)
		}
		if strings.Contains(closingBraces, token) {
			if !stack.Has() {
				return false
			}

			lastToken := stack.Peek()
			if matches[token] == lastToken {
				stack, _ = stack.Pop()
			} else {
				return false
			}
		}
	}

	return stack.Has() == false
}
