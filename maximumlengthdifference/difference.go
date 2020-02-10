package maximumlengthdifference

import "math"

func MxDifLg(a1 []string, a2 []string) int {

	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}

	max := 0.0

	for _, s1 := range a1 {
		for _, s2 := range a2 {
			diff := float64(len(s1) - len(s2))
			max = math.Max(max, math.Abs(diff))
		}
	}

	return int(max)
}
