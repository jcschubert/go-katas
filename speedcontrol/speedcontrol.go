package speedcontrol

import "math"

// Gps calculates the floor of the *maximum* average speed
// per hour for of the distance covered at that particular
// amount of seconds.
func Gps(interval int, distances []float64) int {
	max, avgSpeed := 0.0, 0.0
	for i := 1; i < len(distances); i++ {
		d := distances[i] - distances[i-1]
		avgSpeed = math.Floor((3600 * d) / float64(interval))
		max = math.Max(max, avgSpeed)
	}
	return int(max)
}
