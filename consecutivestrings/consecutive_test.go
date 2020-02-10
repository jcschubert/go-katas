package consecutivestrings

import "testing"

func TestLongestConsec(t *testing.T) {
	cases := []struct {
		strings  []string
		k        int
		expected string
	}{
		{
			[]string{"zone", "abigail", "theta", "form", "libe", "zas", "theta", "abigail"},
			2,
			"abigailtheta",
		},
		{
			[]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"},
			1,
			"oocccffuucccjjjkkkjyyyeehh",
		},
		{
			[]string{"itvayloxrp", "wkppqsztdkmvcuwvereiupccauycnjutlv", "vweqilsfytihvrzlaodfixoyxvyuyvgpck"},
			2,
			"wkppqsztdkmvcuwvereiupccauycnjutlvvweqilsfytihvrzlaodfixoyxvyuyvgpck",
		},
	}

	const msg = "LongestConsec(%v, %d) should return %q, but returned %q instead."

	for _, c := range cases {
		got := LongestConsec(c.strings, c.k)
		if got != c.expected {
			t.Fatalf(msg, c.strings, c.k, c.expected, got)
		}
	}
}
