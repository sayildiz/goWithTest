package revisitarraysslicesgenerics

// Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {

	add := func(acc, x int) int { return acc + x }

	return Reduce(0, add, numbers)
}

// SumAllTails calculates the sums of all but the first number given a collection of slices.
func SumAllTails(numbersToSum ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}

	return Reduce([]int{}, sumTail, numbersToSum)
}

// my solution
func Reduce[T, R any](initval R, f func(a R, b T) R, list []T) R {
	for _, l := range list {
		initval = f(initval, l)
	}
	return initval
}

func Find[A any](items []A, predicate func(A) bool) (value A, found bool) {
	for _, v := range items {
		if predicate(v) {
			return v, true
		}
	}
	return
}

func ReduceFromBook[A any](collection []A, f func(A, A) A, initialValue A) A {
	var result = initialValue

	for _, x := range collection {
		result = f(result, x)
	}
	return result
}
