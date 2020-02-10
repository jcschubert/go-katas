package speedcontrol

func Gps(s int, dists []float64) int {
	interval, maxSpeed := float64(s), 0.0

	for i := 1; i < len(dists); i++ {
		avgSpeed := (dists[i] - dists[i-1]) * 3600 / interval
		if avgSpeed > maxSpeed {
			maxSpeed = avgSpeed
		}
	}
	return int(maxSpeed)
}
