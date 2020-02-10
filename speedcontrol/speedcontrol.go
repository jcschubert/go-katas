package speedcontrol

import "math"

// Gps calculates the floor of the *maximum* average speed
// per hour for of the distance covered at that particular
// amount of seconds.
func Gps(s int, distances []float64) int {
	avgSpeed, max, interval := 0.0, 0.0, float64(s)

	for i := 1; i < len(distances); i++ {
		delta := distances[i] - distances[i-1]
		avgSpeed = math.Floor((3600 * delta) / interval)
		max = math.Max(max, avgSpeed)
	}

	return int(max)
}
