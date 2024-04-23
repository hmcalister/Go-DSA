package defaultcomparators

// Comparators take two items, a and b, and returns a negative value if a<b, a positive value if a>b, and zero if a==b
type ComparatorFunction[T any] func(a, b T) int

func DefaultIntegerComparator(a, b int) int {
	return a - b
}

// Floats are difficult to compare, and hence we need some slightly ugly logic

func DefaultFloat32Comparator(a, b float32) int {
	if a < b {
		return -1
	}

	if a > b {
		return +1
	}

	return 0
}

func DefaultFloat64Comparator(a, b float64) int {
	if a < b {
		return -1
	}

	if a > b {
		return +1
	}

	return 0
}

// Strings are compared lexicographically

func DefaultStringComparator(a, b string) int {
	if a < b {
		return -1
	}

	if a > b {
		return +1
	}

	return 0
}
