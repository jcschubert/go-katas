package maximumlengthdifference

import "math"

func MxDifLg(a1 []string, a2 []string) int {
	strings := append(a1, a2...)

	if len(strings) == 0 {
		return -1
	}

	max, min := 0.0, 0.0

	for _, s := range strings {
		len := float64(len(s))
		max = math.Max(max, len)
		min = math.Min(min, len)
	}

	return int(max - min)
}
