package pdqselect

import (
	"cmp"
	"fmt"
	"math/rand/v2"
	"slices"
	"sort"
	"testing"
)

func TestSelect(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		k     int
	}{
		{"Small sorted", []int{1, 2, 3, 4, 5}, 3},
		{"Small reversed", []int{5, 4, 3, 2, 1}, 3},
		{"Medium random", []int{3, 7, 2, 1, 4, 6, 5, 8, 9}, 5},
		{"Large random", []int{15, 3, 9, 8, 5, 2, 7, 1, 6, 13, 11, 12, 10, 4, 14}, 8},
		{"All equal", []int{1, 1, 1, 1, 1}, 3},
		{"Mostly equal", []int{2, 2, 2, 2, 1, 2, 2, 3, 2, 2}, 6},
		{"Single element", []int{42}, 1},
		{"Two elements", []int{2, 1}, 1},
	}

	for _, tc := range testCases {
		t.Run("Select/"+tc.name, func(t *testing.T) {
			testSelect(t, tc.input, tc.k, func(input []int, k int) {
				Select(sort.IntSlice(input), k)
			})
		})

		t.Run("Ordered/"+tc.name, func(t *testing.T) {
			testSelect(t, tc.input, tc.k, func(input []int, k int) {
				Ordered(input, k)
			})
		})

		t.Run("Func/"+tc.name, func(t *testing.T) {
			testSelect(t, tc.input, tc.k, func(input []int, k int) {
				Func(input, k, cmp.Compare)
			})
		})
	}
}

func testSelect(t *testing.T, input []int, k int, selectFunc func([]int, int)) {
	t.Helper()
	inputCopy := make([]int, len(input))
	copy(inputCopy, input)

	selectFunc(inputCopy, k)

	// Check that all elements to the left of k are <= input[k-1]
	for i := 0; i < k-1; i++ {
		if inputCopy[i] > inputCopy[k-1] {
			t.Errorf("Element at index %d (%d) is greater than the k-th element (%d)", i, inputCopy[i], inputCopy[k-1])
		}
	}

	// Check that all elements to the right of k are >= input[k-1]
	for i := k; i < len(inputCopy); i++ {
		if inputCopy[i] < inputCopy[k-1] {
			t.Errorf("Element at index %d (%d) is less than the k-th element (%d)", i, inputCopy[i], inputCopy[k-1])
		}
	}
}

func FuzzSelect(f *testing.F) {
	f.Add([]byte{1, 4}, uint(2))
	f.Add([]byte{1, 4, 2}, uint(2))
	f.Add([]byte{1, 4, 2, 1}, uint(2))
	f.Add([]byte{1, 2, 3, 4, 5}, uint(3))
	f.Add([]byte{5, 4, 3, 2, 1}, uint(2))
	f.Add([]byte{1, 1, 1, 1, 1}, uint(1))
	f.Add([]byte{1, 4, 7, 2, 0}, uint(4))

	f.Fuzz(func(t *testing.T, data []byte, k uint) {
		if len(data) == 0 {
			return // Skip empty slices
		}
		if k == 0 || int(k) > len(data) {
			return // Skip invalid k values
		}

		// Convert byte slice to int slice
		input := make([]int, len(data))
		for i, b := range data {
			input[i] = int(b)
		}

		fuzzSelect(t, input, int(k), "Select", func(slice []int, k int) {
			Select(sort.IntSlice(slice), k)
		})

		fuzzSelect(t, input, int(k), "Ordered", func(slice []int, k int) {
			Ordered(slice, k)
		})

		fuzzSelect(t, input, int(k), "Func", func(slice []int, k int) {
			Func(slice, k, cmp.Compare)
		})
	})
}

func fuzzSelect(t *testing.T, input []int, k int, name string, selectFunc func([]int, int)) {
	t.Helper()

	// Create a copy for sorting
	sorted := make([]int, len(input))
	copy(sorted, input)
	sort.Ints(sorted)

	// Create another copy for selecting
	output := make([]int, len(input))
	copy(output, input)

	// Run Select
	selectFunc(output, k)

	// Check if the first k elements are the k smallest (unsorted)
	firstK := make([]int, k)
	copy(firstK, output[:k])
	sort.Ints(firstK)

	for i := 0; i < k; i++ {
		if firstK[i] != sorted[i] {
			t.Errorf("%s(k=%d, n=%d): expected %d in first k elements, but got %d\ninput:  %v\nsorted: %v\noutput: %v",
				name, k, len(input), sorted[i], firstK[i], input, sorted, output)
		}
	}

	// Check if all elements in the first k are smaller than or equal to all elements after k
	max := findMax(output[:k])
	for i := k; i < len(output); i++ {
		if output[i] < max {
			t.Errorf("%s(k=%d, n=%d): element at index %d (%d) is smaller than max of first k elements (%d)\ninput:  %v\nsorted: %v\noutput: %v",
				name, k, len(input), i, input[i], max, input, sorted, output)
		}
	}
}

func BenchmarkSelect(b *testing.B) {
	sizes := []int{100000, 10000, 1000}
	kRatios := []float64{0.01, 0.05, 0.1}
	distributions := []string{"random", "sorted", "reversed", "equal", "mostly_equal"}

	for _, size := range sizes {
		for _, kRatio := range kRatios {
			k := int(float64(size) * kRatio)
			for _, dist := range distributions {
				benchName := fmt.Sprintf("n=%d/k=%d/%s", size, k, dist)

				b.Run("Select/"+benchName, func(b *testing.B) {
					data := generateSlice(size, dist)
					b.ReportAllocs()
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						dataCopy := make([]int, len(data))
						copy(dataCopy, data)
						Select(sort.IntSlice(dataCopy), k)
					}
				})

				b.Run("Ordered/"+benchName, func(b *testing.B) {
					data := generateSlice(size, dist)
					b.ReportAllocs()
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						dataCopy := make([]int, len(data))
						copy(dataCopy, data)
						Ordered(dataCopy, k)
					}
				})

				b.Run("Func/"+benchName, func(b *testing.B) {
					data := generateSlice(size, dist)
					b.ReportAllocs()
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						dataCopy := make([]int, len(data))
						copy(dataCopy, data)
						Func(dataCopy, k, cmp.Compare)
					}
				})

				b.Run("Sort/"+benchName, func(b *testing.B) {
					data := generateSlice(size, dist)
					b.ReportAllocs()
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						dataCopy := make([]int, len(data))
						copy(dataCopy, data)
						slices.Sort(dataCopy)
						_ = dataCopy[k-1] // Select k-th element
					}
				})
			}
		}
	}
}

func findMax(slice []int) int {
	if len(slice) == 0 {
		return 0
	}
	max := slice[0]
	for _, v := range slice[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// generateSlice creates a slice of ints with the specified size and distribution
func generateSlice(size int, distribution string) []int {
	slice := make([]int, size)
	switch distribution {
	case "random":
		for i := range slice {
			slice[i] = rand.Int()
		}
	case "sorted":
		for i := range slice {
			slice[i] = i
		}
	case "reversed":
		for i := range slice {
			slice[i] = size - i
		}
	case "equal":
		for i := range slice {
			slice[i] = 42
		}
	case "mostly_equal":
		for i := range slice {
			if rand.Float32() < 0.9 {
				slice[i] = 42
			} else {
				slice[i] = rand.Int()
			}
		}
	}
	return slice
}

func benchmarkPDQSelect(b *testing.B, size int, k int, distribution string) {
	data := generateSlice(size, distribution)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dataCopy := make([]int, len(data))
		copy(dataCopy, data)
		Select(sort.IntSlice(dataCopy), k)
	}
}

func benchmarkSortSelect(b *testing.B, size int, k int, distribution string) {
	data := generateSlice(size, distribution)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dataCopy := make([]int, len(data))
		copy(dataCopy, data)
		slices.Sort(dataCopy)
	}
}
