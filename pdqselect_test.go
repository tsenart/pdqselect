package pdqselect

import (
	"cmp"
	"encoding/binary"
	"fmt"
	"math"
	"math/rand/v2"
	"slices"
	"sort"
	"testing"
	"time"
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
	f.Add(encodeInts(1, 4), uint16(1), uint16(0), uint16(2))
	f.Add(encodeInts(1, 4, 2), uint16(2), uint16(0), uint16(3))
	f.Add(encodeInts(1, 4, 2, 1), uint16(2), uint16(1), uint16(4))
	f.Add(encodeInts(1, 2, 3, 4, 5), uint16(3), uint16(0), uint16(5))
	f.Add(encodeInts(5, 4, 3, 2, 1), uint16(2), uint16(1), uint16(4))
	f.Add(encodeInts(1, 1, 1, 1, 1), uint16(1), uint16(0), uint16(5))
	f.Add(encodeInts(1, 4, 7, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1), uint16(7), uint16(3), uint16(12))
	f.Add(encodeInts(254, 4, 7, 2, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 253), uint16(7), uint16(0), uint16(16))
	f.Add(encodeInts(0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 253, 0, 0, 0, 0, 0, 0), uint16(0), uint16(20), uint16(12))

	now := time.Now().UnixNano()
	rng := rand.New(rand.NewPCG(uint64(now), uint64(now>>32)))

	for _, dist := range []string{
		"random",
		"sorted",
		"reversed",
		"mostly_sorted",
		"organ_pipe",
		"sawtooth",
		"push_front",
		"push_middle",
		"zipf",
	} {
		for _, size := range []int{10, 100, 1000} {
			data := generateSlice(rng, size, dist)
			encodedData := encodeInts(data...)
			f.Add(encodedData, uint16(size/2), uint16(0), uint16(size))
			f.Add(encodedData, uint16(1), uint16(0), uint16(size))
			f.Add(encodedData, uint16(size), uint16(0), uint16(size))
		}
	}

	f.Fuzz(func(t *testing.T, data []byte, k, a, b uint16) {
		if len(data)%4 != 0 {
			return // Skip if data length is not a multiple of 4
		}

		// Convert byte slice to int slice
		input := decodeInts(data)

		if len(input) == 0 {
			return // Skip empty slices
		}

		k = k % uint16(len(input))
		if k == 0 {
			k++
		}

		testSelect(t, input, 0, len(input), int(k), "Select", func(slice []int, a, b, k int) {
			Select(sort.IntSlice(slice), k)
		})

		testSelect(t, input, 0, len(input), int(k), "Ordered", func(slice []int, a, b, k int) {
			Ordered(slice, k)
		})

		testSelect(t, input, 0, len(input), int(k), "Func", func(slice []int, a, b, k int) {
			Func(slice, k, cmp.Compare)
		})

		testSelect(t, input, 0, len(input), int(k), "pdqselect", func(slice []int, a, b, k int) {
			pdqselect(sort.IntSlice(slice), 0, len(slice), k-1, 0)
		})

		testSelect(t, input, 0, len(input), int(k), "pdqselectOrdered", func(slice []int, a, b, k int) {
			pdqselectOrdered(slice, 0, len(slice), k-1, 0)
		})

		testSelect(t, input, 0, len(input), int(k), "pdqselectFunc", func(slice []int, a, b, k int) {
			pdqselectFunc(slice, 0, len(slice), k-1, 0, cmp.Compare)
		})

		// Ensure a, b, and k are within bounds
		a = a % uint16(len(input))
		b = b % uint16(len(input))
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

		testSelect(t, input, int(a), int(b), int(k), "heapSelect", func(slice []int, a, b, k int) {
			heapSelect(sort.IntSlice(slice), a, b, k-1)
		})

		testSelect(t, input, int(a), int(b), int(k), "heapSelectOrdered", func(slice []int, a, b, k int) {
			heapSelectOrdered(slice, a, b, k-1)
		})

		testSelect(t, input, int(a), int(b), int(k), "heapSelectFunc", func(slice []int, a, b, k int) {
			heapSelectFunc(slice, a, b, k-1, cmp.Compare)
		})
	})
}

func encodeInts(ints ...int) []byte {
	buf := make([]byte, len(ints)*4)
	for i, v := range ints {
		binary.BigEndian.PutUint32(buf[i*4:], uint32(v))
	}
	return buf
}

func decodeInts(data []byte) []int {
	ints := make([]int, len(data)/4)
	for i := range ints {
		ints[i] = int(binary.BigEndian.Uint32(data[i*4:]))
	}
	return ints
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

	// Assert that the kth element is the expected one
	if output[a+k-1] != sorted[a+k-1] {
		t.Errorf("%s(a=%d, b=%d, k=%d, n=%d): k-th element (%d) does not match sorted input (%d)\ninput:  %v\nsorted: %v\noutput: %v",
			name, a, b, k, b-a, output[a+k-1], sorted[a+k-1], input, sorted, output)
	}

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
	rng := rand.New(rand.NewPCG(42, 42)) // Static seed for consistent benchmarks
	for _, n := range []int{1e6, 1e4, 100} {
		for _, k := range []int{1, 100, 1000} {
			if k > n {
				continue
			}

			for _, dist := range []string{"random", "sorted", "reversed", "mostly_sorted"} {
				benchName := fmt.Sprintf("n=%d/k=%d/%s", n, k, dist)
				data := generateSlice(rng, n, dist)

				b.Run("fn=Sort/"+benchName, func(b *testing.B) {
					b.ReportAllocs()
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						dataCopy := make([]int, len(data))
						copy(dataCopy, data)
						slices.Sort(dataCopy)
						_ = dataCopy[k-1] // Select k-th element
					}
				})

				b.Run("fn=Ordered/"+benchName, func(b *testing.B) {
					b.ReportAllocs()
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						dataCopy := make([]int, len(data))
						copy(dataCopy, data)
						Ordered(dataCopy, k)
					}
				})

				b.Run("fn=Func/"+benchName, func(b *testing.B) {
					b.ReportAllocs()
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						dataCopy := make([]int, len(data))
						copy(dataCopy, data)
						Func(dataCopy, k, cmp.Compare)
					}
				})

				b.Run("fn=Select/"+benchName, func(b *testing.B) {
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

// generateSlice creates a slice of ints with the specified size and distribution
func generateSlice(rng *rand.Rand, size int, distribution string) []int {
	slice := make([]int, size)
	switch distribution {
	case "random":
		for i := range slice {
			slice[i] = rng.Int()
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
			slice[i] = i
		}
		// Shuffle about 10% of the elements
		for i := 0; i < size/10; i++ {
			j := rng.IntN(size)
			k := rng.IntN(size)
			slice[j], slice[k] = slice[k], slice[j]
		}
	case "organ_pipe":
		mid := size / 2
		for i := 0; i < mid; i++ {
			slice[i] = i
			slice[size-1-i] = i
		}
		if size%2 != 0 {
			slice[mid] = mid
		}
	case "sawtooth":
		period := int(math.Sqrt(float64(size)))
		for i := range slice {
			slice[i] = i % period
		}
	case "push_front":
		for i := range slice {
			if i < size/2 {
				slice[i] = 0
			} else {
				slice[i] = i - size/2 + 1
			}
		}
	case "push_middle":
		for i := range slice {
			if i < size/4 || i >= 3*size/4 {
				slice[i] = i
			} else {
				slice[i] = size / 2
			}
		}
	case "zipf":
		zipf := rand.NewZipf(rng, 1.5, 1.0, uint64(size-1))
		for i := range slice {
			slice[i] = int(zipf.Uint64())
		}
	default:
		panic("unknown distribution")
	}
	return slice
}
