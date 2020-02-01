package squaresum

// SquareSum squares each number passed into it and
// then sums the results together.
func SquareSum(numbers []int) (result int) {
	for _, i := range numbers {
		result += i * i
	}
	return
}
