package speedcontrol

import "testing"

func TestGps(t *testing.T) {
	cases := []struct {
		interval  int
		distances []float64
		expected  int
	}{
		{15, []float64{}, 0},
		{15, []float64{0.0}, 0},
		{15, []float64{0.0, 0.19, 0.5, 0.75, 1.0, 1.25, 1.5, 1.75, 2.0, 2.25}, 74},
		{20, []float64{0.0, 0.23, 0.46, 0.69, 0.92, 1.15, 1.38, 1.61}, 41},
	}

	const msg string = "Gps(%d, %v) returned %d, but %d was expected."

	for _, c := range cases {
		got := Gps(c.interval, c.distances)
		if got != c.expected {
			t.Fatalf(msg, c.interval, c.distances, got, c.expected)
		}
	}
}
