// The pdqselect package implements an adaptive selection algorithm that finds
// the k-th smallest elements in an ordered data structure. It is based on Go's pdqsort
// implementation, which is a hybrid sorting algorithm that combines quicksort, insertion sort,
// heapsort and other pattern defeating techniques to achieve optimal performance on a wide range of data.
package pdqselect

import (
	"cmp"
	"math/bits"
	"sort"
)

// Select swaps elements in the data provided so that the first k elements
// (i.e. the elements occuping indices 0, 1, ..., k-1) are the smallest k elements
// in the data. It doesn't guarantee any particular order among the smallest k elements,
// only that they are the smallest k elements in the data.
//
// Select implements Hoare's Selection Algorithm and runs in O(n) time, so it
// is asymptotically faster than sorting or other heap-like implementations for
// finding the smallest k elements in a data structure.
//
// It's an adaptation of Go's internal pdqsort implementation, which makes it adaptive
// to bad data patterns like already sorted data, duplicate elements, and more.
func Select(data sort.Interface, k int) {
	n := data.Len()
	if k < 1 || k > n {
		return
	}
	pdqselect(data, 0, n, k-1, bits.Len(uint(n)))
}

// Ordered is a specialized version of Select that works with slices of
// ordered types (i.e. types that implement the cmp.Ordered interface).
func Ordered[T cmp.Ordered](data []T, k int) {
	n := len(data)
	if k < 1 || k > n {
		return
	}
	pdqselectOrdered(data, 0, n, k-1, bits.Len(uint(n)))
}

// Func is a generic version of Select that allows the caller to provide
// a custom comparison function to determine the order of elements.
func Func[E any](data []E, k int, cmp func(i, j E) int) {
	n := len(data)
	if k < 1 || k > n {
		return
	}
	pdqselectFunc(data, 0, n, k-1, bits.Len(uint(n)), cmp)
}

func pdqselect(data sort.Interface, a, b, k, limit int) {
	const maxInsertion = 12

	var (
		wasBalanced    = true
		wasPartitioned = true
	)

	for {
		length := b - a

		if length <= maxInsertion {
			insertionSort(data, a, b)
			return
		}

		// Fall back to heap select if too many bad choices were made.
		if limit == 0 {
			heapSelect(data, a, b, k+1)
			return
		}

		// Break patterns if the last partitioning was imbalanced
		if !wasBalanced {
			breakPatterns(data, a, b)
			limit--
		}

		pivot, hint := choosePivot(data, a, b)
		if hint == decreasingHint {
			reverseRange(data, a, b)
			// The chosen pivot was pivot-a elements after the start of the array.
			// After reversing it is pivot-a elements before the end of the array.
			// The idea came from Rust's implementation.
			pivot = (b - 1) - (pivot - a)
			hint = increasingHint
		}

		// Check if the slice is likely already sorted
		if wasBalanced && wasPartitioned && hint == increasingHint {
			if partialInsertionSort(data, a, b) {
				return
			}
		}

		// Probably the slice contains many duplicate elements, partition the slice into
		// elements equal to and elements greater than the pivot.
		if a > 0 && !data.Less(a-1, pivot) {
			mid := partitionEqual(data, a, b, pivot)
			a = mid
			continue
		}

		mid, alreadyPartitioned := partition(data, a, b, pivot)
		wasPartitioned = alreadyPartitioned

		leftLen, rightLen := mid-a, b-mid
		balanceThreshold := length / 8

		if k < mid {
			wasBalanced = leftLen >= balanceThreshold
			b = mid
		} else if k > mid {
			wasBalanced = rightLen >= balanceThreshold
			a = mid + 1
		} else {
			return
		}
	}
}

func pdqselectOrdered[T cmp.Ordered](data []T, a, b, k, limit int) {
	const maxInsertion = 12

	var (
		wasBalanced    = true
		wasPartitioned = true
	)

	for {
		length := b - a

		if length <= maxInsertion {
			insertionSortOrdered(data, a, b)
			return
		}

		// Fall back to heap select if too many bad choices were made.
		if limit == 0 {
			heapSelectOrdered(data, a, b, k+1)
			return
		}

		// Break patterns if the last partitioning was imbalanced
		if !wasBalanced {
			breakPatternsOrdered(data, a, b)
			limit--
		}

		pivot, hint := choosePivotOrdered(data, a, b)
		if hint == decreasingHint {
			reverseRangeOrdered(data, a, b)
			// The chosen pivot was pivot-a elements after the start of the array.
			// After reversing it is pivot-a elements before the end of the array.
			// The idea came from Rust's implementation.
			pivot = (b - 1) - (pivot - a)
			hint = increasingHint
		}

		// Check if the slice is likely already sorted
		if wasBalanced && wasPartitioned && hint == increasingHint {
			if partialInsertionSortOrdered(data, a, b) {
				return
			}
		}

		// Probably the slice contains many duplicate elements, partition the slice into
		// elements equal to and elements greater than the pivot.
		if a > 0 && !cmp.Less(data[a-1], data[pivot]) {
			mid := partitionEqualOrdered(data, a, b, pivot)
			a = mid
			continue
		}

		mid, alreadyPartitioned := partitionOrdered(data, a, b, pivot)
		wasPartitioned = alreadyPartitioned

		leftLen, rightLen := mid-a, b-mid
		balanceThreshold := length / 8

		if k < mid {
			wasBalanced = leftLen >= balanceThreshold
			b = mid
		} else if k > mid {
			wasBalanced = rightLen >= balanceThreshold
			a = mid + 1
		} else {
			return
		}
	}
}

func pdqselectFunc[E any](data []E, a, b, k, limit int, cmp func(a, b E) int) {
	const maxInsertion = 12

	var (
		wasBalanced    = true
		wasPartitioned = true
	)

	for {
		length := b - a

		if length <= maxInsertion {
			insertionSortCmpFunc(data, a, b, cmp)
			return
		}

		// Fall back to heap select if too many bad choices were made.
		if limit == 0 {
			heapSelectFunc(data, a, b, k+1, cmp)
			return
		}

		// Break patterns if the last partitioning was imbalanced
		if !wasBalanced {
			breakPatternsCmpFunc(data, a, b, cmp)
			limit--
		}

		pivot, hint := choosePivotCmpFunc(data, a, b, cmp)
		if hint == decreasingHint {
			reverseRangeCmpFunc(data, a, b, cmp)
			// The chosen pivot was pivot-a elements after the start of the array.
			// After reversing it is pivot-a elements before the end of the array.
			// The idea came from Rust's implementation.
			pivot = (b - 1) - (pivot - a)
			hint = increasingHint
		}

		// Check if the slice is likely already sorted
		if wasBalanced && wasPartitioned && hint == increasingHint {
			if partialInsertionSortCmpFunc(data, a, b, cmp) {
				return
			}
		}

		// Probably the slice contains many duplicate elements, partition the slice into
		// elements equal to and elements greater than the pivot.
		if a > 0 && !(cmp(data[a-1], data[pivot]) < 0) {
			mid := partitionEqualCmpFunc(data, a, b, pivot, cmp)
			a = mid
			continue
		}

		mid, alreadyPartitioned := partitionCmpFunc(data, a, b, pivot, cmp)
		wasPartitioned = alreadyPartitioned

		leftLen, rightLen := mid-a, b-mid
		balanceThreshold := length / 8

		if k < mid {
			wasBalanced = leftLen >= balanceThreshold
			b = mid
		} else if k > mid {
			wasBalanced = rightLen >= balanceThreshold
			a = mid + 1
		} else {
			return
		}
	}
}

func heapSelect(data sort.Interface, a, b, k int) {
	n := b - a

	// Build max-heap of first k elements
	for i := (k - 1) / 2; i >= 0; i-- {
		siftDown(data, i, k, a)
	}

	// Process remaining elements
	for i := k; i < n; i++ {
		if data.Less(a+i, a) {
			data.Swap(a, a+i)
			siftDown(data, 0, k, a)
		}
	}
}

func heapSelectOrdered[T cmp.Ordered](data []T, a, b, k int) {
	n := b - a

	// Build max-heap of first k elements
	for i := (k - 1) / 2; i >= 0; i-- {
		siftDownOrdered(data, i, k, a)
	}

	// Process remaining elements
	for i := k; i < n; i++ {
		if cmp.Less(data[a+i], data[a]) {
			data[a], data[a+i] = data[a+i], data[a]
			siftDownOrdered(data, 0, k, a)
		}
	}
}

func heapSelectFunc[E any](data []E, a, b, k int, cmp func(a, b E) int) {
	n := b - a

	// Build max-heap of first k elements
	for i := (k - 1) / 2; i >= 0; i-- {
		siftDownCmpFunc(data, i, k, a, cmp)
	}

	// Process remaining elements
	for i := k; i < n; i++ {
		if cmp(data[a+i], data[a]) < 0 {
			data[a], data[a+i] = data[a+i], data[a]
			siftDownCmpFunc(data, 0, k, a, cmp)
		}
	}
}
