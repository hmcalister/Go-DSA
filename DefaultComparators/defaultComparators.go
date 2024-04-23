package defaultcomparators

// Comparators take two items, a and b, and returns a negative value if a<b, a positive value if a>b, and zero if a==b
type ComparatorFunction[T any] func(a, b T) int

func DefaultIntegerComparator(a, b int) int {
	return a - b
}

