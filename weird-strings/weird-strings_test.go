package weirdstrings

import (
	"testing"
)

type WeirdCaseTestCase struct {
	input    string
	expected string
}

func TestToWeirdCase(t *testing.T) {

	testCases := []WeirdCaseTestCase{
		{"a", "A"},
		{"A", "A"},
		{"aa", "Aa"},
		{"ABC", "AbC"},
		{"abc def", "AbC DeF"},
	}
	runWeirdCaseTestCases(t, testCases)
}

func runWeirdCaseTestCases(t *testing.T, cases []WeirdCaseTestCase) {
	for _, tc := range cases {
		t.Run("It turns words into weirdly cased words", func(t *testing.T) {
			result := toWeirdCase(tc.input)

			if result != tc.expected {
				t.Logf("Given: %+v, expected: %q, got: %q", tc.input, tc.expected, result)
				t.Fail()
			}
		})
	}
}
