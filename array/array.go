package array

// Sum returns sum of the slice of numbers passed
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll takes variable arg slice of integers and returns a slice of
// their sum
func SumAll(numbersToSum ...[]int) (sum []int) {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

// SumAllTails takes a variable arg slice and returns the sum of tail of each slice
func SumAllTails(numbersToSum ...[]int) (sum []int) {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
