package ipvalidation

import "testing"

type ipValidationCases struct {
	input    string
	expected bool
}

func TestIpValidation(t *testing.T) {
	runTests(t, IsValidIP, []ipValidationCases{
		{"1.2.3.4", true},
		{"192.168.0.7", true},
		{"1.2.3", false},
		{"123.45.67.89", true},
		{"1.2.3.4.5", false},
		{"123.456.78.90", false},
		{"123.045.078.089", false},
	})
}

func runTests(t *testing.T, f func(string) bool, cases []ipValidationCases) {
	t.Run("Testing Ip Validation", func(t *testing.T) {
		for _, c := range cases {
			result := f(c.input)
			if result != c.expected {
				t.Fatalf("Input: %q, expected %t, but got %t", c.input, c.expected, result)
			}
		}
	})
}
