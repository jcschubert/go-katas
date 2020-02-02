package ipvalidation

import (
	"strconv"
	"strings"
)

// IsValidIP check wether ip is a valid IPv4 address
// in dot-decimal format.
//
// Examples :
//            "1.2.3.4" -> true
//            "192.168.0.7" -> true
//            "1.2.3" -> false
//            "999.999.999.999" -> false
func IsValidIP(ip string) bool {
	octets := strings.Split(ip, ".")
	return has4Octets(octets) &&
		allOctetsWithinAllowedRange(octets)
}

func has4Octets(octets []string) bool {
	return len(octets) == 4
}

func allOctetsWithinAllowedRange(octets []string) bool {
	for _, octet := range octets {
		if len(octet) > 1 && octet[0] == '0' {
			return false
		}

		intVal, err := strconv.Atoi(octet)

		if err != nil || intVal < 0 || intVal > 255 {
			return false
		}
	}
	return true
}
