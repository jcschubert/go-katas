package spinwords

import (
	"testing"
)

type SpinWordsTestCase struct {
	input    string
	expected string
}

func TestCreatePhoneNumber(t *testing.T) {

	singleWordCases := []SpinWordsTestCase{
		{"Hi", "Hi"},
		{"Welcome", "emocleW"},
		{"to", "to"},
		{"CodeWars", "sraWedoC"},
	}

	sentenceCases := []SpinWordsTestCase{
		{"Hey fellow warriors", "Hey wollef sroirraw"},
		{"Burgers are my favorite fruit", "sregruB are my etirovaf tiurf"},
		{"Pizza is the best vegetable", "azziP is the best elbategev"},
	}

	runSpinWordsCases(t, singleWordCases)
	runSpinWordsCases(t, sentenceCases)
}

func runSpinWordsCases(t *testing.T, cases []SpinWordsTestCase) {
	for _, tc := range cases {
		t.Run("It reverts words in a sentence that are longer than 5 characters", func(t *testing.T) {
			result := SpinWords(tc.input)

			if result != tc.expected {
				t.Logf("Given: %+v, expected: %q, got: %q", tc.input, tc.expected, result)
				t.Fail()
			}
		})
	}
}
