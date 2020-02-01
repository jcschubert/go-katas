package phone

import (
	"fmt"
)

const formatString = "(%d%d%d) %d%d%d-%d%d%d%d"

// CreatePhoneNumber formats a phone number consisting of 10 digits.
// and returns it as a string.
//
// Example: [10]uint{1,2,3,4,5,6,7,8,9,10} is formatted as "(123) 456-7890"
func CreatePhoneNumber(numbers [10]uint) string {
	return fmt.Sprintf(formatString, numbers[0], numbers[1], numbers[2],
		numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8], numbers[9])
}
