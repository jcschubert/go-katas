package maximumlengthdifference

import "testing"

func TestMaxLengthDifference(t *testing.T) {
	cases := []struct {
		a1     []string
		a2     []string
		result int
	}{
		{[]string{}, []string{}, -1},
		{
			[]string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"},
			[]string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"},
			13,
		},
	}

	const msg string = "MxDifLg(%v, %v) should return %d, but returned %d instead."

	for _, c := range cases {
		got := MxDifLg(c.a1, c.a2)
		if got != c.result {
			t.Fatalf(msg, c.a1, c.a2, c.result, got)
		}
	}
}
