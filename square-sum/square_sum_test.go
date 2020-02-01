package squaresum

import (
	"testing"
)

type SquareSumTestCase struct {
	input    []int
	expected int
}

func TestSquareSum(t *testing.T) {
	runSquareSumTestCases(t, []SquareSumTestCase{
		{[]int{1, 2}, 5},
	})
}

func runSquareSumTestCases(t *testing.T, cases []SquareSumTestCase) {
	for _, tc := range cases {
		t.Run("It turns words into weirdly cased words", func(t *testing.T) {
			result := SquareSum(tc.input)

			if result != tc.expected {
				t.Logf("Given: %+v, expected: %d, got: %d", tc.input, tc.expected, result)
				t.Fail()
			}
		})
	}
}
