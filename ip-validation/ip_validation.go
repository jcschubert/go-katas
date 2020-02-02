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
		if hasLeadingZero(octet) || !isWithinRange(octet) {
			return false
		}
	}
	return true
}

func hasLeadingZero(octet string) bool {
	return len(octet) > 1 && octet[0] == '0'
}

func isWithinRange(octet string) bool {
	intVal, err := strconv.Atoi(octet)
	return err == nil && 0 <= intVal && intVal <= 255
}
