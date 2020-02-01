package splitstrings

import (
	"reflect"
	"testing"
)

type TestCase struct {
	input    string
	expected []string
}

func TestSplitStrings(t *testing.T) {
	cases := []TestCase{
		{"ab", []string{"ab"}},
	}
	runCases(t, Solution, cases)
}

func runCases(t *testing.T, toTest func(string) []string, cases []TestCase) {
	for _, tc := range cases {
		t.Run("It turns words into weirdly cased words", func(t *testing.T) {
			result := toTest(tc.input)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Logf("Given: %+v, expected: %q, got: %q", tc.input, tc.expected, result)
				t.Fail()
			}
		})
	}
}
