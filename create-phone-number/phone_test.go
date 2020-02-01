package phone

import (
	"testing"
)

type PhoneTestCase struct {
	input    [10]uint
	expected string
}

func TestCreatePhoneNumber(t *testing.T) {

	testCases := []PhoneTestCase{
		{[10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, "(123) 456-7890"},
		{[10]uint{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, "(000) 000-0000"},
		{[10]uint{1, 7, 1, 2, 2, 3, 4, 8, 7, 1}, "(171) 223-4871"},
		{[10]uint{0, 4, 0, 1, 5, 2, 9, 2, 1, 0}, "(040) 152-9210"},
	}

	for _, tc := range testCases {
		t.Run("It formats a phone number", func(t *testing.T) {
			result := CreatePhoneNumber(tc.input)

			if result != tc.expected {
				t.Logf("Given: %+v, expected: %q, got: %q", tc.input, tc.expected, result)
				t.Fail()
			}
		})
	}
}
