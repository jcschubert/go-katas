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

	for _, t := range tokens {
		if strings.Contains(openingBraces, t) {
			stack = stack.Push(t)
		}
		if strings.Contains(closingBraces, t) {
			if !stack.Has() {
				return false
			}

			lastT := stack.Peek()
			if matches[t] == lastT {
				stack, _ = stack.Pop()
			}
		}
	}

	return stack.Has() == false
}
