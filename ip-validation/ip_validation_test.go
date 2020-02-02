package ipvalidation

import "testing"

type ipValidationCases struct {
	input    string
	expected bool
}

func TestIpValidation(t *testing.T) {
	runTests(t, is_valid_ip, []ipValidationCases{
		{"1.2.3.4", true},
	})
}

func runTests(t *testing.T, f func(string) bool, cases []ipValidationCases) {
	t.Run("Testing Ip Validation", func(t *testing.T) {
		for _, c := range cases {
			result := f(c.input)
			if result != c.expected {
				t.Fatalf("Input: %v, expected %t, but got %t", c.input, c.expected, result)
			}
		}
	})
}
