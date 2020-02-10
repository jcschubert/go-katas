package speedcontrol

import "math"

// Gps calculates the floor of the *maximum* average speed
// per hour for of the distance covered at that particular
// amount of seconds.
func Gps(s int, x []float64) int {
	if len(x) < 2 {
		return 0
	}

	max := 0.0
	for i := 1; i < len(x); i++ {
		d := x[i] - x[i-1]
		avgSpeed := math.Floor((3600 * d) / float64(s))
		if avgSpeed > max {
			max = avgSpeed
		}
	}
	return int(max)
}
