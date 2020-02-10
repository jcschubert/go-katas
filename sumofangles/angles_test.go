package sumofangles

import "testing"

func TestAngles(t *testing.T) {
	cases := []struct {
		n           int
		sumOfAngles int
	}{
		{3, 180},
		{4, 360},
	}

	const msg = "Angles(%d) should return %d, but returned %d instead."

	for _, c := range cases {
		got := Angle(c.n)
		if c.sumOfAngles != got {
			t.Fatalf(msg, c.n, c.sumOfAngles, got)
		}
	}
}
