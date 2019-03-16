package arrays

// Sum sums numbers in an array
func Sum(numbers []int) (acc int) {
	for _, number := range numbers {
		acc += number
	}
	return
}

// SumTail sums numbers in an array except Head
func SumTail(numbers []int) (acc int) {
	for i, number := range numbers {
		if i == 0 {
			continue
		}
		acc += number
	}
	return
}

// SumAll returns the sum of each list passed to it
func SumAll(numbers ...[]int) (results []int) {
	for _, intList := range numbers {
		results = append(results, Sum(intList))
	}
	return
}

// SumAllTails adds up all except Head in a list
func SumAllTails(numbers ...[]int) (results []int) {
	for _, intList := range numbers {
		results = append(results, SumTail(intList))
	}
	return
}
