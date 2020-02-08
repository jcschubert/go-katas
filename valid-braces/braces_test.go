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
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			got := ValidBraces(c.stringOfBraces)
			if got != c.expected {
				t.Fatalf("ValidBraces(%s) should return %t, returned %t",
					c.stringOfBraces, c.expected, got)
			}
		})
	}
}
