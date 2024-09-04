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
			testSelect(t, tc.input, 0, len(tc.input), tc.k, "Select", func(input []int, a, b, k int) {
				Select(sort.IntSlice(input), k)
			})
		})

		t.Run("Ordered/"+tc.name, func(t *testing.T) {
			testSelect(t, tc.input, 0, len(tc.input), tc.k, "Ordered", func(input []int, a, b, k int) {
				Ordered(input, k)
			})
		})

		t.Run("Func/"+tc.name, func(t *testing.T) {
			testSelect(t, tc.input, 0, len(tc.input), tc.k, "Func", func(input []int, a, b, k int) {
				Func(input, k, cmp.Compare)
			})
		})
	}
}

func FuzzSelect(f *testing.F) {
	f.Add([]byte{1, 4}, uint16(1), uint16(0), uint16(2))
	f.Add([]byte{1, 4, 2}, uint16(2), uint16(0), uint16(3))
	f.Add([]byte{1, 4, 2, 1}, uint16(2), uint16(1), uint16(4))
	f.Add([]byte{1, 2, 3, 4, 5}, uint16(3), uint16(0), uint16(5))
	f.Add([]byte{5, 4, 3, 2, 1}, uint16(2), uint16(1), uint16(4))
	f.Add([]byte{1, 1, 1, 1, 1}, uint16(1), uint16(0), uint16(5))
	f.Add([]byte{1, 4, 7, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, uint16(7), uint16(3), uint16(12))
	f.Add([]byte{254, 4, 7, 2, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 253}, uint16(7), uint16(0), uint16(16))
	f.Add([]byte{0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 253, 0, 0, 0, 0, 0, 0}, uint16(0), uint16(20), uint16(12))

	f.Fuzz(func(t *testing.T, data []byte, k, a, b uint16) {
		if len(data) == 0 {
			return // Skip empty slices
		}

		// Convert byte slice to int slice
		input := make([]int, len(data))
		for i, v := range data {
			input[i] = int(v)
		}

		k = k % uint16(len(data))
		if k == 0 {
			k++
		}

		testSelect(t, input, 0, len(data), int(k), "Select", func(slice []int, a, b, k int) {
			Select(sort.IntSlice(slice), k)
		})

		testSelect(t, input, 0, len(data), int(k), "Ordered", func(slice []int, a, b, k int) {
			Ordered(slice, k)
		})

		testSelect(t, input, 0, len(data), int(k), "Func", func(slice []int, a, b, k int) {
			Func(slice, k, cmp.Compare)
		})

		// Ensure a, b, and k are within bounds
		a = a % uint16(len(data))
		b = b % uint16(len(data))
		if a > b {
			a, b = b, a // Ensure a < b
		} else if a == b {
			b++ // Ensure b is at least 1 greater than a
		}

		n := b - a
		k %= n
		if k == 0 {
			k++
		}

		// testSelect(t, input, int(a), int(b), int(k), "pdqselect", func(slice []int, a, b, k int) {
		// 	pdqselect(sort.IntSlice(slice), a, b, k-1, bits.Len(uint(n)))
		// })

		// testSelect(t, input, int(a), int(b), int(k), "pdqselectOrdered", func(slice []int, a, b, k int) {
		// 	pdqselectOrdered(slice, a, b, k-1, bits.Len(uint(n)))
		// })

		//testSelect(t, input, int(a), int(b), int(k), "pdqselectFunc", func(slice []int, a, b, k int) {
		//	pdqselectFunc(slice, a, b, k-1, bits.Len(uint(n)), cmp.Compare)
		//})

		testSelect(t, input, int(a), int(b), int(k), "heapSelect", func(slice []int, a, b, k int) {
			heapSelect(sort.IntSlice(slice), a, b, k)
		})

		testSelect(t, input, int(a), int(b), int(k), "heapSelectOrdered", func(slice []int, a, b, k int) {
			heapSelectOrdered(slice, a, b, k)
		})

		testSelect(t, input, int(a), int(b), int(k), "heapSelectFunc", func(slice []int, a, b, k int) {
			heapSelectFunc(slice, a, b, k, cmp.Compare)
		})
	})
}

func testSelect(t *testing.T, input []int, a, b, k int, name string, selectFunc func([]int, int, int, int)) {
	t.Helper()

	// Create a copy for sorting
	sorted := make([]int, len(input))
	copy(sorted, input)
	sort.Ints(sorted[a:b])

	// Create another copy for selecting
	output := make([]int, len(input))
	copy(output, input)

	// Run pdqselect
	selectFunc(output, a, b, k)

	// Get the first k elements, sort them, and compare with sorted slice
	firstK := make([]int, k)
	copy(firstK, output[a:a+k])
	sort.Ints(firstK)
	for i := range firstK {
		if firstK[i] != sorted[a+i] {
			t.Errorf("%s(a=%d, b=%d, k=%d, n=%d): sorted output element at index %d (%d) does not match sorted input (%d)\ninput:  %v\nsorted: %v\noutput: %v\nfirstK: %v",
				name, a, b, k, b-a, i, firstK[i], sorted[a+i], input, sorted, output, firstK)
		}
	}

	// Check if all elements before and including k are smaller or equal, and all elements after k are larger or equal
	for i := a; i < a+k; i++ {
		if output[i] > sorted[a+k-1] {
			t.Errorf("%s(a=%d, b=%d, k=%d, n=%d): element at index %d (%d) is larger than k-th element (%d)\ninput:  %v\nsorted: %v\noutput: %v",
				name, a, b, k, b-a, i, output[i], sorted[a+k-1], input, sorted, output)
		}
	}

	for i := a + k - 1; i < b; i++ {
		if output[i] < output[a+k-1] {
			t.Errorf("%s(a=%d, b=%d, k=%d, n=%d): element at index %d (%d) is smaller than k-th element (%d)\ninput:  %v\nsorted: %v\noutput: %v",
				name, a, b, k, b-a, i, output[i], output[a+k-1], input, sorted, output)
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

				b.Run("Select/"+benchName, func(b *testing.B) {
					b.ReportAllocs()
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						dataCopy := make([]int, len(data))
						copy(dataCopy, data)
						Select(sort.IntSlice(dataCopy), k)
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
