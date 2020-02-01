package reversedstrings

import (
	"testing"
)

type WeirdCaseTestCase struct {
	input    string
	expected string
}

func TestReversedString(t *testing.T) {
	runReversedStringTestCases(t, Solution, []WeirdCaseTestCase{
		{"", ""},
		{"world", "dlrow"},
		{"Hello", "olleH"},
	})
}

func runReversedStringTestCases(t *testing.T, toTest func(string) string, cases []WeirdCaseTestCase) {
	for _, tc := range cases {
		t.Run("It reverses strings", func(t *testing.T) {
			result := toTest(tc.input)

			if result != tc.expected {
				t.Logf("Given: %q, expected: %q, got: %q", tc.input, tc.expected, result)
				t.Fail()
			}
		})
	}
}
