package braces

import "testing"

func TestValidBraces(t *testing.T) {
	cases := []struct {
		desc           string
		stringOfBraces string
		expected       bool
	}{
		{"An empty string containing no braces is valid.",
			"", true},
		{"A simple set of parentheses returns true.",
			"()", true},
		{"An open parenthese that is not closed return false.",
			"(", false},
		{"A simple set of brackets returns true.",
			"[]", true},
		{"A simple set of curly braces returns true.",
			"{}", true},
		{"Unmatched pairs of braces return false",
			"(}", false},
		{"Several pairs of braces return true",
			"()()()", true},
		{"Several pairs of different braces return true",
			"()[]{}", true},
		{"Several nested pairs return true",
			"((())())", true},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			got := ValidBraces(c.stringOfBraces)
			if got != c.expected {
				t.Fatalf("ValidBraces(%q) should return %t, returned %t",
					c.stringOfBraces, c.expected, got)
			}
		})
	}
}
