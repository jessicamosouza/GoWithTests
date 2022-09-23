package arraysSlices

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum

}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// The tail of a collection is all items in the collection except the first one (the "head").
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		//take from 1 to the end" with numbers[1:]
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			//start with an empty slice sums and append to it the result of Sum as we work through
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
