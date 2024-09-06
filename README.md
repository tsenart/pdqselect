# pdqselect: Pattern-Defeating QuickSelect for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/tsenart/pdqselect.svg)](https://pkg.go.dev/github.com/tsenart/pdqselect)
[![Go Report Card](https://goreportcard.com/badge/github.com/tsenart/pdqselect)](https://goreportcard.com/report/github.com/tsenart/pdqselect)

`pdqselect` is a high-performance, adaptive selection algorithm for Go that finds the k-th smallest elements in an ordered data structure. It's based on Go's internal `pdqsort` implementation, combining the best aspects of quicksort, insertion sort, and heapsort with pattern-defeating techniques to achieve optimal performance across a wide range of data distributions.

## Features

- **O(n) Average Time Complexity**: Outperforms sorting-based selection methods for large datasets.
- **Adaptive**: Efficiently handles various data patterns, including already sorted data, reverse-sorted data, and data with many duplicates.
- **In-Place**: Operates directly on the input slice without requiring additional memory allocation.
- **Generic**: Supports multiple data types and custom comparison functions.
- **Robust**: Gracefully degrades to heap select for pathological cases, ensuring O(n log k) worst-case performance.

## Installation

```bash
go get github.com/tsenart/pdqselect
```

## Usage

The package provides three main functions:

1. `Select`: Works with any type that implements the `sort.Interface`.
2. `Ordered`: Specialized for slices of ordered types (implements `cmp.Ordered`).
3. `Func`: Generic version that allows a custom comparison function.

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

Using `Ordered` with a slice of ordered types:

```go
import (
    "fmt"

    "github.com/tsenart/pdqselect"
)

func main() {
    data := []int{5, 4, 0, 10, 1, 2, 1}
    pdqselect.Ordered(data, 3)
    fmt.Println(data[:3]) // Output: [1 0 1]
}
```

Using `Func` with a custom comparison function:

```go
import (
    "fmt"

    "github.com/tsenart/pdqselect"
)

func main() {
    data := []float64{5.5, 4.4, 0.0, 10.1, 1.1, 2.2, 1.0}
    pdqselect.Func(data, 3, cmp.Compare)
    fmt.Println(data[:3]) // Output: [1 0 1.1]
}
```

## Benchmarks

`pdqselect` significantly outperforms standard library's `sort.Slice` followed by indexing for selecting k-th elements, especially for large datasets:

```
goos: darwin
goarch: arm64
pkg: github.com/tsenart/pdqselect
cpu: Apple M3 Max
                                         │     Sort      │               Ordered               │                 Func                  │                Select                 │
                                         │    sec/op     │   sec/op     vs base                │    sec/op     vs base                 │    sec/op     vs base                 │
Select/n=1000000/k=1/random-16             55352.8µ ± 1%   689.4µ ± 1%  -98.75% (p=0.000 n=15)   1946.7µ ± 1%   -96.48% (p=0.000 n=15)   2495.0µ ± 0%   -95.49% (p=0.000 n=15)
Select/n=1000000/k=1/sorted-16               696.0µ ± 2%   686.8µ ± 2%   -1.32% (p=0.000 n=15)   1958.7µ ± 0%  +181.43% (p=0.000 n=15)   1886.9µ ± 1%  +171.11% (p=0.000 n=15)
Select/n=1000000/k=1/reversed-16             954.5µ ± 2%   691.7µ ± 2%  -27.53% (p=0.000 n=15)   2018.0µ ± 0%  +111.43% (p=0.000 n=15)   2586.6µ ± 1%  +171.00% (p=0.000 n=15)
Select/n=1000000/k=1/mostly_sorted-16      20980.2µ ± 0%   684.6µ ± 1%  -96.74% (p=0.000 n=15)   2003.5µ ± 2%   -90.45% (p=0.000 n=15)   1899.6µ ± 1%   -90.95% (p=0.000 n=15)
Select/n=1000000/k=100/random-16            57.142m ± 0%   6.425m ± 1%  -88.76% (p=0.000 n=15)   10.514m ± 0%   -81.60% (p=0.000 n=15)   12.347m ± 1%   -78.39% (p=0.000 n=15)
Select/n=1000000/k=100/sorted-16             703.1µ ± 1%   707.9µ ± 1%   +0.69% (p=0.019 n=15)   2066.3µ ± 1%  +193.90% (p=0.000 n=15)   2085.7µ ± 1%  +196.66% (p=0.000 n=15)
Select/n=1000000/k=100/reversed-16           935.5µ ± 1%   940.1µ ± 1%        ~ (p=0.486 n=15)   2256.5µ ± 1%  +141.20% (p=0.000 n=15)   3123.7µ ± 0%  +233.90% (p=0.000 n=15)
Select/n=1000000/k=100/mostly_sorted-16     20.959m ± 0%   1.820m ± 0%  -91.32% (p=0.000 n=15)    4.372m ± 0%   -79.14% (p=0.000 n=15)    4.577m ± 0%   -78.16% (p=0.000 n=15)
Select/n=1000000/k=1000/random-16           56.427m ± 0%   3.394m ± 1%  -93.99% (p=0.000 n=15)    6.195m ± 1%   -89.02% (p=0.000 n=15)    6.950m ± 0%   -87.68% (p=0.000 n=15)
Select/n=1000000/k=1000/sorted-16            702.4µ ± 1%   710.1µ ± 1%        ~ (p=0.174 n=15)   2071.9µ ± 1%  +194.96% (p=0.000 n=15)   2084.6µ ± 1%  +196.77% (p=0.000 n=15)
Select/n=1000000/k=1000/reversed-16          935.3µ ± 1%   930.5µ ± 0%   -0.52% (p=0.037 n=15)   2247.1µ ± 1%  +140.25% (p=0.000 n=15)   3120.0µ ± 0%  +233.58% (p=0.000 n=15)
Select/n=1000000/k=1000/mostly_sorted-16    20.911m ± 0%   1.854m ± 1%  -91.13% (p=0.000 n=15)    4.444m ± 1%   -78.75% (p=0.000 n=15)    4.577m ± 0%   -78.11% (p=0.000 n=15)
Select/n=10000/k=1/random-16               143.998µ ± 4%   9.528µ ± 1%  -93.38% (p=0.000 n=15)   21.810µ ± 0%   -84.85% (p=0.000 n=15)   27.952µ ± 0%   -80.59% (p=0.000 n=15)
Select/n=10000/k=1/sorted-16                 9.698µ ± 0%   9.509µ ± 1%   -1.95% (p=0.001 n=15)   21.784µ ± 1%  +124.62% (p=0.000 n=15)   20.942µ ± 1%  +115.94% (p=0.000 n=15)
Select/n=10000/k=1/reversed-16              11.709µ ± 3%   9.559µ ± 2%  -18.36% (p=0.000 n=15)   22.503µ ± 1%   +92.19% (p=0.000 n=15)   28.405µ ± 1%  +142.59% (p=0.000 n=15)
Select/n=10000/k=1/mostly_sorted-16         78.436µ ± 0%   9.548µ ± 2%  -87.83% (p=0.000 n=15)   21.864µ ± 0%   -72.13% (p=0.000 n=15)   21.138µ ± 0%   -73.05% (p=0.000 n=15)
Select/n=10000/k=100/random-16              157.90µ ± 6%   22.22µ ± 1%  -85.93% (p=0.000 n=15)    63.68µ ± 1%   -59.67% (p=0.000 n=15)    58.68µ ± 0%   -62.84% (p=0.000 n=15)
Select/n=10000/k=100/sorted-16               9.772µ ± 0%   9.750µ ± 0%        ~ (p=0.616 n=15)   23.646µ ± 0%  +141.98% (p=0.000 n=15)   24.335µ ± 0%  +149.03% (p=0.000 n=15)
Select/n=10000/k=100/reversed-16             11.95µ ± 0%   11.98µ ± 0%        ~ (p=0.205 n=15)    26.02µ ± 0%  +117.68% (p=0.000 n=15)    32.00µ ± 0%  +167.73% (p=0.000 n=15)
Select/n=10000/k=100/mostly_sorted-16        76.74µ ± 0%   13.42µ ± 0%  -82.51% (p=0.000 n=15)    46.51µ ± 0%   -39.40% (p=0.000 n=15)    45.99µ ± 0%   -40.07% (p=0.000 n=15)
Select/n=10000/k=1000/random-16             165.25µ ± 3%   23.02µ ± 0%  -86.07% (p=0.000 n=15)    66.88µ ± 0%   -59.53% (p=0.000 n=15)    60.24µ ± 1%   -63.54% (p=0.000 n=15)
Select/n=10000/k=1000/sorted-16              9.814µ ± 1%   9.905µ ± 1%        ~ (p=0.113 n=15)   24.169µ ± 0%  +146.27% (p=0.000 n=15)   24.496µ ± 0%  +149.60% (p=0.000 n=15)
Select/n=10000/k=1000/reversed-16            12.11µ ± 0%   12.13µ ± 0%        ~ (p=0.184 n=15)    26.16µ ± 0%  +116.12% (p=0.000 n=15)    32.06µ ± 0%  +164.87% (p=0.000 n=15)
Select/n=10000/k=1000/mostly_sorted-16       76.53µ ± 0%   13.88µ ± 0%  -81.86% (p=0.000 n=15)    48.87µ ± 1%   -36.14% (p=0.000 n=15)    48.07µ ± 1%   -37.18% (p=0.000 n=15)
Select/n=100/k=1/random-16                   718.5n ± 1%   123.6n ± 0%  -82.80% (p=0.000 n=15)    252.7n ± 0%   -64.83% (p=0.000 n=15)    330.8n ± 0%   -53.96% (p=0.000 n=15)
Select/n=100/k=1/sorted-16                   132.7n ± 0%   124.4n ± 0%   -6.25% (p=0.000 n=15)    251.5n ± 1%   +89.53% (p=0.000 n=15)    272.2n ± 0%  +105.12% (p=0.000 n=15)
Select/n=100/k=1/reversed-16                 155.9n ± 0%   115.7n ± 0%  -25.79% (p=0.000 n=15)    260.1n ± 0%   +66.84% (p=0.000 n=15)    330.7n ± 0%  +112.12% (p=0.000 n=15)
Select/n=100/k=1/mostly_sorted-16            358.8n ± 0%   124.1n ± 1%  -65.41% (p=0.000 n=15)    258.8n ± 1%   -27.87% (p=0.000 n=15)    333.1n ± 0%    -7.16% (p=0.000 n=15)
Select/n=100/k=100/random-16                 693.7n ± 0%   127.0n ± 0%  -81.69% (p=0.000 n=15)    237.8n ± 0%   -65.72% (p=0.000 n=15)    325.9n ± 0%   -53.02% (p=0.000 n=15)
Select/n=100/k=100/sorted-16                 134.3n ± 0%   116.6n ± 0%  -13.18% (p=0.000 n=15)    269.0n ± 0%  +100.30% (p=0.000 n=15)    328.0n ± 1%  +144.23% (p=0.000 n=15)
Select/n=100/k=100/reversed-16               157.2n ± 1%   125.9n ± 0%  -19.91% (p=0.000 n=15)    227.1n ± 1%   +44.47% (p=0.000 n=15)    292.1n ± 1%   +85.81% (p=0.000 n=15)
Select/n=100/k=100/mostly_sorted-16          323.9n ± 1%   128.9n ± 1%  -60.20% (p=0.000 n=15)    243.1n ± 1%   -24.95% (p=0.000 n=15)    334.1n ± 0%    +3.15% (p=0.000 n=15)
geomean                                      67.37µ        21.49µ       -68.11%                   51.40µ        -23.71%                   58.24µ        -13.56%

                                         │     Sort     │                Ordered                │                 Func                  │               Select                │
                                         │     B/op     │     B/op      vs base                 │     B/op      vs base                 │     B/op      vs base               │
Select/n=1000000/k=1/random-16             7.633Mi ± 0%   7.633Mi ± 0%       ~ (p=0.731 n=15)     7.633Mi ± 0%       ~ (p=0.826 n=15)     7.633Mi ± 0%  +0.00% (p=0.000 n=15)
Select/n=1000000/k=1/sorted-16             7.633Mi ± 0%   7.633Mi ± 0%       ~ (p=0.339 n=15)     7.633Mi ± 0%       ~ (p=0.767 n=15)     7.633Mi ± 0%  +0.00% (p=0.000 n=15)
Select/n=1000000/k=1/reversed-16           7.633Mi ± 0%   7.633Mi ± 0%       ~ (p=0.174 n=15)     7.633Mi ± 0%       ~ (p=0.233 n=15)     7.633Mi ± 0%  +0.00% (p=0.000 n=15)
Select/n=1000000/k=1/mostly_sorted-16      7.633Mi ± 0%   7.633Mi ± 0%       ~ (p=0.148 n=15)     7.633Mi ± 0%       ~ (p=0.223 n=15)     7.633Mi ± 0%  +0.00% (p=0.000 n=15)
Select/n=1000000/k=100/random-16           7.633Mi ± 0%   7.633Mi ± 0%  +0.00% (p=0.001 n=15)     7.633Mi ± 0%  +0.00% (p=0.001 n=15)     7.633Mi ± 0%  +0.00% (p=0.000 n=15)
Select/n=1000000/k=100/sorted-16           7.633Mi ± 0%   7.633Mi ± 0%       ~ (p=1.000 n=15)     7.633Mi ± 0%       ~ (p=0.875 n=15)     7.633Mi ± 0%  +0.00% (p=0.000 n=15)
Select/n=1000000/k=100/reversed-16         7.633Mi ± 0%   7.633Mi ± 0%       ~ (p=0.758 n=15)     7.633Mi ± 0%       ~ (p=0.555 n=15)     7.633Mi ± 0%  +0.00% (p=0.000 n=15)
Select/n=1000000/k=100/mostly_sorted-16    7.633Mi ± 0%   7.633Mi ± 0%       ~ (p=0.865 n=15)     7.633Mi ± 0%       ~ (p=0.926 n=15)     7.633Mi ± 0%  +0.00% (p=0.000 n=15)
Select/n=1000000/k=1000/random-16          7.633Mi ± 0%   7.633Mi ± 0%  +0.00% (p=0.000 n=15)     7.633Mi ± 0%  +0.00% (p=0.017 n=15)     7.633Mi ± 0%  +0.00% (p=0.000 n=15)
Select/n=1000000/k=1000/sorted-16          7.633Mi ± 0%   7.633Mi ± 0%       ~ (p=1.000 n=15)     7.633Mi ± 0%       ~ (p=0.558 n=15)     7.633Mi ± 0%  +0.00% (p=0.000 n=15)
Select/n=1000000/k=1000/reversed-16        7.633Mi ± 0%   7.633Mi ± 0%       ~ (p=0.515 n=15)     7.633Mi ± 0%       ~ (p=0.092 n=15)     7.633Mi ± 0%  +0.00% (p=0.000 n=15)
Select/n=1000000/k=1000/mostly_sorted-16   7.633Mi ± 0%   7.633Mi ± 0%       ~ (p=0.064 n=15)     7.633Mi ± 0%       ~ (p=0.085 n=15)     7.633Mi ± 0%  +0.00% (p=0.000 n=15)
Select/n=10000/k=1/random-16               80.00Ki ± 0%   80.00Ki ± 0%       ~ (p=1.000 n=15) ¹   80.00Ki ± 0%       ~ (p=1.000 n=15) ¹   80.02Ki ± 0%  +0.03% (p=0.000 n=15)
Select/n=10000/k=1/sorted-16               80.00Ki ± 0%   80.00Ki ± 0%       ~ (p=1.000 n=15) ¹   80.00Ki ± 0%       ~ (p=1.000 n=15) ¹   80.02Ki ± 0%  +0.03% (p=0.000 n=15)
Select/n=10000/k=1/reversed-16             80.00Ki ± 0%   80.00Ki ± 0%       ~ (p=1.000 n=15) ¹   80.00Ki ± 0%       ~ (p=1.000 n=15) ¹   80.02Ki ± 0%  +0.03% (p=0.000 n=15)
Select/n=10000/k=1/mostly_sorted-16        80.00Ki ± 0%   80.00Ki ± 0%       ~ (p=1.000 n=15) ¹   80.00Ki ± 0%       ~ (p=1.000 n=15) ¹   80.02Ki ± 0%  +0.03% (p=0.000 n=15)
Select/n=10000/k=100/random-16             80.00Ki ± 0%   80.00Ki ± 0%       ~ (p=1.000 n=15) ¹   80.00Ki ± 0%       ~ (p=1.000 n=15) ¹   80.02Ki ± 0%  +0.03% (p=0.000 n=15)
Select/n=10000/k=100/sorted-16             80.00Ki ± 0%   80.00Ki ± 0%       ~ (p=1.000 n=15) ¹   80.00Ki ± 0%       ~ (p=1.000 n=15) ¹   80.02Ki ± 0%  +0.03% (p=0.000 n=15)
Select/n=10000/k=100/reversed-16           80.00Ki ± 0%   80.00Ki ± 0%       ~ (p=1.000 n=15) ¹   80.00Ki ± 0%       ~ (p=1.000 n=15) ¹   80.02Ki ± 0%  +0.03% (p=0.000 n=15)
Select/n=10000/k=100/mostly_sorted-16      80.00Ki ± 0%   80.00Ki ± 0%       ~ (p=1.000 n=15) ¹   80.00Ki ± 0%       ~ (p=1.000 n=15) ¹   80.02Ki ± 0%  +0.03% (p=0.000 n=15)
Select/n=10000/k=1000/random-16            80.00Ki ± 0%   80.00Ki ± 0%       ~ (p=1.000 n=15) ¹   80.00Ki ± 0%       ~ (p=1.000 n=15) ¹   80.02Ki ± 0%  +0.03% (p=0.000 n=15)
Select/n=10000/k=1000/sorted-16            80.00Ki ± 0%   80.00Ki ± 0%       ~ (p=1.000 n=15) ¹   80.00Ki ± 0%       ~ (p=1.000 n=15) ¹   80.02Ki ± 0%  +0.03% (p=0.000 n=15)
Select/n=10000/k=1000/reversed-16          80.00Ki ± 0%   80.00Ki ± 0%       ~ (p=1.000 n=15) ¹   80.00Ki ± 0%       ~ (p=1.000 n=15) ¹   80.02Ki ± 0%  +0.03% (p=0.000 n=15)
Select/n=10000/k=1000/mostly_sorted-16     80.00Ki ± 0%   80.00Ki ± 0%       ~ (p=1.000 n=15) ¹   80.00Ki ± 0%       ~ (p=1.000 n=15) ¹   80.02Ki ± 0%  +0.03% (p=0.000 n=15)
Select/n=100/k=1/random-16                   896.0 ± 0%     896.0 ± 0%       ~ (p=1.000 n=15) ¹     896.0 ± 0%       ~ (p=1.000 n=15) ¹     920.0 ± 0%  +2.68% (p=0.000 n=15)
Select/n=100/k=1/sorted-16                   896.0 ± 0%     896.0 ± 0%       ~ (p=1.000 n=15) ¹     896.0 ± 0%       ~ (p=1.000 n=15) ¹     920.0 ± 0%  +2.68% (p=0.000 n=15)
Select/n=100/k=1/reversed-16                 896.0 ± 0%     896.0 ± 0%       ~ (p=1.000 n=15) ¹     896.0 ± 0%       ~ (p=1.000 n=15) ¹     920.0 ± 0%  +2.68% (p=0.000 n=15)
Select/n=100/k=1/mostly_sorted-16            896.0 ± 0%     896.0 ± 0%       ~ (p=1.000 n=15) ¹     896.0 ± 0%       ~ (p=1.000 n=15) ¹     920.0 ± 0%  +2.68% (p=0.000 n=15)
Select/n=100/k=100/random-16                 896.0 ± 0%     896.0 ± 0%       ~ (p=1.000 n=15) ¹     896.0 ± 0%       ~ (p=1.000 n=15) ¹     920.0 ± 0%  +2.68% (p=0.000 n=15)
Select/n=100/k=100/sorted-16                 896.0 ± 0%     896.0 ± 0%       ~ (p=1.000 n=15) ¹     896.0 ± 0%       ~ (p=1.000 n=15) ¹     920.0 ± 0%  +2.68% (p=0.000 n=15)
Select/n=100/k=100/reversed-16               896.0 ± 0%     896.0 ± 0%       ~ (p=1.000 n=15) ¹     896.0 ± 0%       ~ (p=1.000 n=15) ¹     920.0 ± 0%  +2.68% (p=0.000 n=15)
Select/n=100/k=100/mostly_sorted-16          896.0 ± 0%     896.0 ± 0%       ~ (p=1.000 n=15) ¹     896.0 ± 0%       ~ (p=1.000 n=15) ¹     920.0 ± 0%  +2.68% (p=0.000 n=15)
geomean                                    144.2Ki        144.2Ki       +0.00%                    144.2Ki       +0.00%                    145.2Ki       +0.67%
¹ all samples are equal

                                         │    Sort    │               Ordered               │                Func                 │               Select                │
                                         │ allocs/op  │ allocs/op   vs base                 │ allocs/op   vs base                 │ allocs/op   vs base                 │
Select/n=1000000/k=1/random-16             1.000 ± 0%   1.000 ± 0%       ~ (p=1.000 n=15) ¹   1.000 ± 0%       ~ (p=1.000 n=15) ¹   2.000 ± 0%  +100.00% (p=0.000 n=15)
Select/n=1000000/k=1/sorted-16             1.000 ± 0%   1.000 ± 0%       ~ (p=1.000 n=15) ¹   1.000 ± 0%       ~ (p=1.000 n=15) ¹   2.000 ± 0%  +100.00% (p=0.000 n=15)
Select/n=1000000/k=1/reversed-16           1.000 ± 0%   1.000 ± 0%       ~ (p=1.000 n=15) ¹   1.000 ± 0%       ~ (p=1.000 n=15) ¹   2.000 ± 0%  +100.00% (p=0.000 n=15)
Select/n=1000000/k=1/mostly_sorted-16      1.000 ± 0%   1.000 ± 0%       ~ (p=1.000 n=15) ¹   1.000 ± 0%       ~ (p=1.000 n=15) ¹   2.000 ± 0%  +100.00% (p=0.000 n=15)
Select/n=1000000/k=100/random-16           1.000 ± 0%   1.000 ± 0%       ~ (p=1.000 n=15) ¹   1.000 ± 0%       ~ (p=1.000 n=15) ¹   2.000 ± 0%  +100.00% (p=0.000 n=15)
Select/n=1000000/k=100/sorted-16           1.000 ± 0%   1.000 ± 0%       ~ (p=1.000 n=15) ¹   1.000 ± 0%       ~ (p=1.000 n=15) ¹   2.000 ± 0%  +100.00% (p=0.000 n=15)
Select/n=1000000/k=100/reversed-16         1.000 ± 0%   1.000 ± 0%       ~ (p=1.000 n=15) ¹   1.000 ± 0%       ~ (p=1.000 n=15) ¹   2.000 ± 0%  +100.00% (p=0.000 n=15)
Select/n=1000000/k=100/mostly_sorted-16    1.000 ± 0%   1.000 ± 0%       ~ (p=1.000 n=15) ¹   1.000 ± 0%       ~ (p=1.000 n=15) ¹   2.000 ± 0%  +100.00% (p=0.000 n=15)
Select/n=1000000/k=1000/random-16          1.000 ± 0%   1.000 ± 0%       ~ (p=1.000 n=15) ¹   1.000 ± 0%       ~ (p=1.000 n=15) ¹   2.000 ± 0%  +100.00% (p=0.000 n=15)
Select/n=1000000/k=1000/sorted-16          1.000 ± 0%   1.000 ± 0%       ~ (p=1.000 n=15) ¹   1.000 ± 0%       ~ (p=1.000 n=15) ¹   2.000 ± 0%  +100.00% (p=0.000 n=15)
Select/n=1000000/k=1000/reversed-16        1.000 ± 0%   1.000 ± 0%       ~ (p=1.000 n=15) ¹   1.000 ± 0%       ~ (p=1.000 n=15) ¹   2.000 ± 0%  +100.00% (p=0.000 n=15)
Select/n=1000000/k=1000/mostly_sorted-16   1.000 ± 0%   1.000 ± 0%       ~ (p=1.000 n=15) ¹   1.000 ± 0%       ~ (p=1.000 n=15) ¹   2.000 ± 0%  +100.00% (p=0.000 n=15)
Select/n=10000/k=1/random-16               1.000 ± 0%   1.000 ± 0%       ~ (p=1.000 n=15) ¹   1.000 ± 0%       ~ (p=1.000 n=15) ¹   2.000 ± 0%  +100.00% (p=0.000 n=15)
Select/n=10000/k=1/sorted-16               1.000 ± 0%   1.000 ± 0%       ~ (p=1.000 n=15) ¹   1.000 ± 0%       ~ (p=1.000 n=15) ¹   2.000 ± 0%  +100.00% (p=0.000 n=15)
Select/n=10000/k=1/reversed-16             1.000 ± 0%   1.000 ± 0%       ~ (p=1.000 n=15) ¹   1.000 ± 0%       ~ (p=1.000 n=15) ¹   2.000 ± 0%  +100.00% (p=0.000 n=15)
Select/n=10000/k=1/mostly_sorted-16        1.000 ± 0%   1.000 ± 0%       ~ (p=1.000 n=15) ¹   1.000 ± 0%       ~ (p=1.000 n=15) ¹   2.000 ± 0%  +100.00% (p=0.000 n=15)
Select/n=10000/k=100/random-16             1.000 ± 0%   1.000 ± 0%       ~ (p=1.000 n=15) ¹   1.000 ± 0%       ~ (p=1.000 n=15) ¹   2.000 ± 0%  +100.00% (p=0.000 n=15)
Select/n=10000/k=100/sorted-16             1.000 ± 0%   1.000 ± 0%       ~ (p=1.000 n=15) ¹   1.000 ± 0%       ~ (p=1.000 n=15) ¹   2.000 ± 0%  +100.00% (p=0.000 n=15)
Select/n=10000/k=100/reversed-16           1.000 ± 0%   1.000 ± 0%       ~ (p=1.000 n=15) ¹   1.000 ± 0%       ~ (p=1.000 n=15) ¹   2.000 ± 0%  +100.00% (p=0.000 n=15)
Select/n=10000/k=100/mostly_sorted-16      1.000 ± 0%   1.000 ± 0%       ~ (p=1.000 n=15) ¹   1.000 ± 0%       ~ (p=1.000 n=15) ¹   2.000 ± 0%  +100.00% (p=0.000 n=15)
Select/n=10000/k=1000/random-16            1.000 ± 0%   1.000 ± 0%       ~ (p=1.000 n=15) ¹   1.000 ± 0%       ~ (p=1.000 n=15) ¹   2.000 ± 0%  +100.00% (p=0.000 n=15)
Select/n=10000/k=1000/sorted-16            1.000 ± 0%   1.000 ± 0%       ~ (p=1.000 n=15) ¹   1.000 ± 0%       ~ (p=1.000 n=15) ¹   2.000 ± 0%  +100.00% (p=0.000 n=15)
Select/n=10000/k=1000/reversed-16          1.000 ± 0%   1.000 ± 0%       ~ (p=1.000 n=15) ¹   1.000 ± 0%       ~ (p=1.000 n=15) ¹   2.000 ± 0%  +100.00% (p=0.000 n=15)
Select/n=10000/k=1000/mostly_sorted-16     1.000 ± 0%   1.000 ± 0%       ~ (p=1.000 n=15) ¹   1.000 ± 0%       ~ (p=1.000 n=15) ¹   2.000 ± 0%  +100.00% (p=0.000 n=15)
Select/n=100/k=1/random-16                 1.000 ± 0%   1.000 ± 0%       ~ (p=1.000 n=15) ¹   1.000 ± 0%       ~ (p=1.000 n=15) ¹   2.000 ± 0%  +100.00% (p=0.000 n=15)
Select/n=100/k=1/sorted-16                 1.000 ± 0%   1.000 ± 0%       ~ (p=1.000 n=15) ¹   1.000 ± 0%       ~ (p=1.000 n=15) ¹   2.000 ± 0%  +100.00% (p=0.000 n=15)
Select/n=100/k=1/reversed-16               1.000 ± 0%   1.000 ± 0%       ~ (p=1.000 n=15) ¹   1.000 ± 0%       ~ (p=1.000 n=15) ¹   2.000 ± 0%  +100.00% (p=0.000 n=15)
Select/n=100/k=1/mostly_sorted-16          1.000 ± 0%   1.000 ± 0%       ~ (p=1.000 n=15) ¹   1.000 ± 0%       ~ (p=1.000 n=15) ¹   2.000 ± 0%  +100.00% (p=0.000 n=15)
Select/n=100/k=100/random-16               1.000 ± 0%   1.000 ± 0%       ~ (p=1.000 n=15) ¹   1.000 ± 0%       ~ (p=1.000 n=15) ¹   2.000 ± 0%  +100.00% (p=0.000 n=15)
Select/n=100/k=100/sorted-16               1.000 ± 0%   1.000 ± 0%       ~ (p=1.000 n=15) ¹   1.000 ± 0%       ~ (p=1.000 n=15) ¹   2.000 ± 0%  +100.00% (p=0.000 n=15)
Select/n=100/k=100/reversed-16             1.000 ± 0%   1.000 ± 0%       ~ (p=1.000 n=15) ¹   1.000 ± 0%       ~ (p=1.000 n=15) ¹   2.000 ± 0%  +100.00% (p=0.000 n=15)
Select/n=100/k=100/mostly_sorted-16        1.000 ± 0%   1.000 ± 0%       ~ (p=1.000 n=15) ¹   1.000 ± 0%       ~ (p=1.000 n=15) ¹   2.000 ± 0%  +100.00% (p=0.000 n=15)
geomean                                    1.000        1.000       +0.00%                    1.000       +0.00%                    2.000       +100.00%
¹ all samples are equal
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built upon Go's internal `pdqsort` algorithm.
