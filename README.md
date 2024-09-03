# pdqselect: Pattern-Defeating QuickSelect for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/tsenart/pdqselect.svg)](https://pkg.go.dev/github.com/tsenart/pdqselect)
[![Go Report Card](https://goreportcard.com/badge/github.com/tsenart/pdqselect)](https://goreportcard.com/report/github.com/tsenart/pdqselect)

`pdqselect` is a high-performance, adaptive selection algorithm for Go that finds the k-th smallest element in a data structure. It's based on Go's internal `pdqsort` implementation, combining the best aspects of quicksort, insertion sort, and heapsort with pattern-defeating techniques to achieve optimal performance across a wide range of data distributions.

## Features

- **O(n) Average Time Complexity**: Outperforms sorting-based selection methods for large datasets.
- **Adaptive**: Efficiently handles various data patterns, including already sorted data, reverse-sorted data, and data with many duplicates.
- **In-Place**: Operates directly on the input slice without requiring additional memory allocation.
- **Generic**: Supports multiple data types and custom comparison functions.
- **Robust**: Gracefully degrades to heapsort for pathological cases, ensuring O(n log n) worst-case performance.

## Installation

```bash
go get github.com/tsenart/pdqselect
```

## Usage

The package provides three main functions:

1. `Select`: Works with any type that implements the `sort.Interface`.
2. `SelectOrdered`: Specialized for slices of ordered types (implements `cmp.Ordered`).
3. `SelectFunc`: Generic version that allows a custom comparison function.

### Examples

Using `Select` with `sort.Interface`:

```go
import (
    "fmt"
    "sort"

    "github.com/tsenart/pdqselect"
)

func main() {
    data := sort.IntSlice{5, 4, 0, 10, 1, 2, 1}
    pdqselect.Select(data, 3)
    fmt.Println(data[:3]) // Output: [1 0 1]
}
```

Using `SelectOrdered` with a slice of ordered types:

```go
import (
    "fmt"

    "github.com/tsenart/pdqselect"
)

func main() {
    data := []int{5, 4, 0, 10, 1, 2, 1}
    pdqselect.SelectOrdered(data, 3)
    fmt.Println(data[:3]) // Output: [1 0 1]
}
```

Using `SelectFunc` with a custom comparison function:

```go
import (
    "fmt"

    "github.com/tsenart/pdqselect"
)

func main() {
    data := []float64{5.5, 4.4, 0.0, 10.1, 1.1, 2.2, 1.0}
    pdqselect.SelectFunc(data, 3, cmp.Compare)
    fmt.Println(data[:3]) // Output: [1 0 1.1]
}
```

## Benchmarks

`pdqselect` significantly outperforms standard library's `sort.Slice` followed by indexing for selecting k-th elements, especially for large datasets:

```
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built upon Go's internal `pdqsort` algorithm.
