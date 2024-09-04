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
			testSelect(t, tc.input, tc.k, "Select", func(input []int, k int) {
				Select(sort.IntSlice(input), k)
			})
		})

		t.Run("Ordered/"+tc.name, func(t *testing.T) {
			testSelect(t, tc.input, tc.k, "Ordered", func(input []int, k int) {
				Ordered(input, k)
			})
		})

		t.Run("Func/"+tc.name, func(t *testing.T) {
			testSelect(t, tc.input, tc.k, "Func", func(input []int, k int) {
				Func(input, k, cmp.Compare)
			})
		})

		t.Run("heapSelect/"+tc.name, func(t *testing.T) {
			testSelect(t, tc.input, tc.k, "heapSelect", func(input []int, k int) {
				pdqselect(sort.IntSlice(input), 0, len(input), k-1, 0) // limit = 0 means we'll use heapSelect
			})
		})

		t.Run("heapSelectOrdered/"+tc.name, func(t *testing.T) {
			testSelect(t, tc.input, tc.k, "heapSelectOrdered", func(input []int, k int) {
				pdqselectOrdered(input, 0, len(input), k-1, 0) // limit = 0 means we'll use heapSelect
			})
		})

		t.Run("heapSelectFunc/"+tc.name, func(t *testing.T) {
			testSelect(t, tc.input, tc.k, "heapSelectFunc", func(input []int, k int) {
				pdqselectFunc(input, 0, len(input), k-1, 0, cmp.Compare) // limit = 0 means we'll use heapSelect
			})
		})
	}
}

func FuzzSelect(f *testing.F) {
	f.Add([]byte{1, 4}, uint(2))
	f.Add([]byte{1, 4, 2}, uint(2))
	f.Add([]byte{1, 4, 2, 1}, uint(2))
	f.Add([]byte{1, 2, 3, 4, 5}, uint(3))
	f.Add([]byte{5, 4, 3, 2, 1}, uint(2))
	f.Add([]byte{1, 1, 1, 1, 1}, uint(1))
	f.Add([]byte{1, 4, 7, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, uint(7))
	f.Add([]byte{254, 4, 7, 2, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 253}, uint(7))

	f.Fuzz(func(t *testing.T, data []byte, k uint) {
		if len(data) == 0 {
			return // Skip empty slices
		}

		// Ensure k is within bounds
		k = k % uint(len(data))

		// Convert byte slice to int slice
		input := make([]int, len(data))
		for i, b := range data {
			input[i] = int(b)
		}

		testSelect(t, input, int(k), "Select", func(slice []int, k int) {
			Select(sort.IntSlice(slice), k)
		})

		testSelect(t, input, int(k), "Ordered", func(slice []int, k int) {
			Ordered(slice, k)
		})

		testSelect(t, input, int(k), "Func", func(slice []int, k int) {
			Func(slice, k, cmp.Compare)
		})

		testSelect(t, input, int(k), "heapSelect", func(slice []int, k int) {
			pdqselect(sort.IntSlice(slice), 0, len(slice), k-1, 0) // limit = 0 means we'll use heapSelect
		})

		testSelect(t, input, int(k), "heapSelectOrdered", func(slice []int, k int) {
			pdqselectOrdered(slice, 0, len(slice), k-1, 0) // limit = 0 means we'll use heapSelect
		})

		testSelect(t, input, int(k), "heapSelectFunc", func(slice []int, k int) {
			pdqselectFunc(slice, 0, len(slice), k-1, 0, cmp.Compare) // limit = 0 means we'll use heapSelect
		})
	})
}

func testSelect(t *testing.T, input []int, k int, name string, selectFunc func([]int, int)) {
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
	for _, n := range []int{1e6, 1e4, 100} {
		for _, k := range []int{1, 100, 1000} {
			if k > n {
				continue
			}

			for _, dist := range []string{"random", "sorted", "reversed", "mostly_sorted"} {
				benchName := fmt.Sprintf("n=%d/k=%d/%s", n, k, dist)
				data := generateSlice(n, dist)

				b.Run("Sort/"+benchName, func(b *testing.B) {
					b.ReportAllocs()
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						dataCopy := make([]int, len(data))
						copy(dataCopy, data)
						slices.Sort(dataCopy)
						_ = dataCopy[k-1] // Select k-th element
					}
				})

				b.Run("Select/"+benchName, func(b *testing.B) {
					b.ReportAllocs()
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						dataCopy := make([]int, len(data))
						copy(dataCopy, data)
						Select(sort.IntSlice(dataCopy), k)
					}
				})

				b.Run("Ordered/"+benchName, func(b *testing.B) {
					b.ReportAllocs()
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						dataCopy := make([]int, len(data))
						copy(dataCopy, data)
						Ordered(dataCopy, k)
					}
				})

				b.Run("Func/"+benchName, func(b *testing.B) {
					b.ReportAllocs()
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						dataCopy := make([]int, len(data))
						copy(dataCopy, data)
						Func(dataCopy, k, cmp.Compare)
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
	case "mostly_sorted":
		for i := range slice {
			if rand.Float32() < 0.9 {
				slice[i] = i
			} else {
				slice[i] = rand.Int()
			}
		}
	}
	return slice
}
