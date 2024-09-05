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
                                         │     Sort     │               Ordered                │                 Func                  │                Select                 │
                                         │    sec/op    │    sec/op     vs base                │    sec/op     vs base                 │    sec/op     vs base                 │
Select/n=1000000/k=1/random-16             53.871m ± 1%    5.685m ± 3%  -89.45% (p=0.000 n=15)    9.024m ± 0%   -83.25% (p=0.000 n=15)   10.602m ± 0%   -80.32% (p=0.000 n=15)
Select/n=1000000/k=1/sorted-16              678.7µ ± 4%    682.1µ ± 0%        ~ (p=0.061 n=15)   1960.4µ ± 2%  +188.86% (p=0.000 n=15)   2203.8µ ± 0%  +224.72% (p=0.000 n=15)
Select/n=1000000/k=1/reversed-16            939.8µ ± 1%    941.6µ ± 0%   +0.19% (p=0.045 n=15)   2159.8µ ± 0%  +129.81% (p=0.000 n=15)   3093.6µ ± 0%  +229.18% (p=0.000 n=15)
Select/n=1000000/k=1/mostly_sorted-16      21.317m ± 0%    1.829m ± 2%  -91.42% (p=0.000 n=15)    4.474m ± 0%   -79.01% (p=0.000 n=15)    4.518m ± 0%   -78.81% (p=0.000 n=15)
Select/n=1000000/k=100/random-16           58.093m ± 0%    6.625m ± 0%  -88.60% (p=0.000 n=15)   10.727m ± 0%   -81.53% (p=0.000 n=15)   12.614m ± 0%   -78.29% (p=0.000 n=15)
Select/n=1000000/k=100/sorted-16            705.3µ ± 2%    706.2µ ± 0%        ~ (p=0.486 n=15)   2041.4µ ± 0%  +189.44% (p=0.000 n=15)   2172.9µ ± 0%  +208.09% (p=0.000 n=15)
Select/n=1000000/k=100/reversed-16          932.0µ ± 3%    929.8µ ± 0%   -0.24% (p=0.001 n=15)   2230.6µ ± 0%  +139.33% (p=0.000 n=15)   3094.2µ ± 1%  +231.98% (p=0.000 n=15)
Select/n=1000000/k=100/mostly_sorted-16    21.512m ± 0%    1.896m ± 0%  -91.18% (p=0.000 n=15)    4.491m ± 1%   -79.12% (p=0.000 n=15)    4.456m ± 1%   -79.29% (p=0.000 n=15)
Select/n=1000000/k=1000/random-16          55.375m ± 0%    3.405m ± 5%  -93.85% (p=0.000 n=15)    6.110m ± 2%   -88.97% (p=0.000 n=15)    7.125m ± 1%   -87.13% (p=0.000 n=15)
Select/n=1000000/k=1000/sorted-16           721.4µ ± 1%    724.2µ ± 1%        ~ (p=0.567 n=15)   1970.4µ ± 1%  +173.12% (p=0.000 n=15)   2219.0µ ± 1%  +207.58% (p=0.000 n=15)
Select/n=1000000/k=1000/reversed-16         958.1µ ± 1%    938.0µ ± 1%   -2.11% (p=0.000 n=15)   2248.2µ ± 1%  +134.65% (p=0.000 n=15)   3094.9µ ± 1%  +223.01% (p=0.000 n=15)
Select/n=1000000/k=1000/mostly_sorted-16   20.990m ± 0%    1.861m ± 0%  -91.13% (p=0.000 n=15)    4.389m ± 0%   -79.09% (p=0.000 n=15)    4.445m ± 0%   -78.82% (p=0.000 n=15)
Select/n=10000/k=1/random-16               136.23µ ± 2%    21.18µ ± 0%  -84.45% (p=0.000 n=15)    63.97µ ± 0%   -53.04% (p=0.000 n=15)    53.66µ ± 1%   -60.61% (p=0.000 n=15)
Select/n=10000/k=1/sorted-16                9.625µ ± 1%    9.643µ ± 0%        ~ (p=0.720 n=15)   22.956µ ± 0%  +138.50% (p=0.000 n=15)   24.184µ ± 0%  +151.26% (p=0.000 n=15)
Select/n=10000/k=1/reversed-16              11.63µ ± 0%    11.64µ ± 0%        ~ (p=0.894 n=15)    24.85µ ± 0%  +113.57% (p=0.000 n=15)    31.75µ ± 0%  +172.90% (p=0.000 n=15)
Select/n=10000/k=1/mostly_sorted-16         78.70µ ± 0%    13.36µ ± 0%  -83.03% (p=0.000 n=15)    47.12µ ± 1%   -40.13% (p=0.000 n=15)    44.12µ ± 1%   -43.95% (p=0.000 n=15)
Select/n=10000/k=100/random-16             149.02µ ± 9%    21.79µ ± 0%  -85.38% (p=0.000 n=15)    61.55µ ± 1%   -58.70% (p=0.000 n=15)    54.90µ ± 1%   -63.16% (p=0.000 n=15)
Select/n=10000/k=100/sorted-16              9.681µ ± 0%    9.709µ ± 0%        ~ (p=0.056 n=15)   23.083µ ± 0%  +138.44% (p=0.000 n=15)   24.324µ ± 0%  +151.26% (p=0.000 n=15)
Select/n=10000/k=100/reversed-16            11.70µ ± 0%    11.76µ ± 2%   +0.51% (p=0.002 n=15)    25.09µ ± 0%  +114.44% (p=0.000 n=15)    31.88µ ± 0%  +172.49% (p=0.000 n=15)
Select/n=10000/k=100/mostly_sorted-16       74.87µ ± 0%    13.26µ ± 2%  -82.29% (p=0.000 n=15)    44.28µ ± 0%   -40.86% (p=0.000 n=15)    42.40µ ± 0%   -43.37% (p=0.000 n=15)
Select/n=10000/k=1000/random-16            137.05µ ± 1%    21.82µ ± 0%  -84.08% (p=0.000 n=15)    61.33µ ± 1%   -55.25% (p=0.000 n=15)    54.26µ ± 1%   -60.41% (p=0.000 n=15)
Select/n=10000/k=1000/sorted-16             9.700µ ± 1%   10.003µ ± 0%   +3.12% (p=0.000 n=15)   23.112µ ± 0%  +138.27% (p=0.000 n=15)   24.330µ ± 0%  +150.82% (p=0.000 n=15)
Select/n=10000/k=1000/reversed-16           11.96µ ± 0%    11.99µ ± 0%   +0.25% (p=0.001 n=15)    25.18µ ± 0%  +110.58% (p=0.000 n=15)    32.00µ ± 0%  +167.61% (p=0.000 n=15)
Select/n=10000/k=1000/mostly_sorted-16      75.62µ ± 1%    13.72µ ± 1%  -81.86% (p=0.000 n=15)    46.81µ ± 0%   -38.10% (p=0.000 n=15)    45.14µ ± 1%   -40.31% (p=0.000 n=15)
Select/n=100/k=1/random-16                  713.4n ± 4%    271.0n ± 1%  -62.01% (p=0.000 n=15)    661.7n ± 1%    -7.25% (p=0.000 n=15)    640.4n ± 0%   -10.23% (p=0.000 n=15)
Select/n=100/k=1/sorted-16                  134.8n ± 1%    134.8n ± 1%        ~ (p=0.830 n=15)    282.7n ± 0%  +109.72% (p=0.000 n=15)    318.9n ± 0%  +136.57% (p=0.000 n=15)
Select/n=100/k=1/reversed-16                154.9n ± 0%    155.7n ± 0%   +0.52% (p=0.002 n=15)    301.6n ± 0%   +94.71% (p=0.000 n=15)    397.7n ± 0%  +156.75% (p=0.000 n=15)
Select/n=100/k=1/mostly_sorted-16           355.1n ± 0%    191.9n ± 1%  -45.96% (p=0.000 n=15)    487.3n ± 0%   +37.23% (p=0.000 n=15)    475.5n ± 0%   +33.91% (p=0.000 n=15)
Select/n=100/k=100/random-16                688.2n ± 1%    180.8n ± 1%  -73.73% (p=0.000 n=15)    410.5n ± 0%   -40.35% (p=0.000 n=15)    418.3n ± 0%   -39.22% (p=0.000 n=15)
Select/n=100/k=100/sorted-16                133.0n ± 0%    134.1n ± 1%   +0.83% (p=0.000 n=15)    281.8n ± 1%  +111.88% (p=0.000 n=15)    320.5n ± 0%  +140.98% (p=0.000 n=15)
Select/n=100/k=100/reversed-16              155.3n ± 1%    154.6n ± 1%        ~ (p=0.418 n=15)    302.2n ± 0%   +94.59% (p=0.000 n=15)    398.9n ± 1%  +156.86% (p=0.000 n=15)
Select/n=100/k=100/mostly_sorted-16         318.5n ± 1%    168.8n ± 1%  -47.00% (p=0.000 n=15)    421.6n ± 1%   +32.37% (p=0.000 n=15)    405.1n ± 1%   +27.19% (p=0.000 n=15)
geomean                                     66.47µ         26.97µ       -59.42%                   64.53µ         -2.92%                   70.53µ         +6.10%

                                         │     Sort     │                Ordered                │                 Func                  │               Select                │
                                         │     B/op     │     B/op      vs base                 │     B/op      vs base                 │     B/op      vs base               │
Select/n=1000000/k=1/random-16             7.633Mi ± 0%   7.633Mi ± 0%       ~ (p=0.276 n=15)     7.633Mi ± 0%       ~ (p=0.162 n=15)     7.633Mi ± 0%  +0.00% (p=0.000 n=15)
Select/n=1000000/k=1/sorted-16             7.633Mi ± 0%   7.633Mi ± 0%  -0.00% (p=0.010 n=15)     7.633Mi ± 0%       ~ (p=0.694 n=15)     7.633Mi ± 0%  +0.00% (p=0.000 n=15)
Select/n=1000000/k=1/reversed-16           7.633Mi ± 0%   7.633Mi ± 0%       ~ (p=0.994 n=15)     7.633Mi ± 0%  -0.00% (p=0.003 n=15)     7.633Mi ± 0%  +0.00% (p=0.000 n=15)
Select/n=1000000/k=1/mostly_sorted-16      7.633Mi ± 0%   7.633Mi ± 0%       ~ (p=0.916 n=15)     7.633Mi ± 0%       ~ (p=0.316 n=15)     7.633Mi ± 0%  +0.00% (p=0.000 n=15)
Select/n=1000000/k=100/random-16           7.633Mi ± 0%   7.633Mi ± 0%       ~ (p=0.469 n=15)     7.633Mi ± 0%       ~ (p=0.786 n=15)     7.633Mi ± 0%  +0.00% (p=0.000 n=15)
Select/n=1000000/k=100/sorted-16           7.633Mi ± 0%   7.633Mi ± 0%       ~ (p=0.483 n=15)     7.633Mi ± 0%       ~ (p=0.058 n=15)     7.633Mi ± 0%  +0.00% (p=0.000 n=15)
Select/n=1000000/k=100/reversed-16         7.633Mi ± 0%   7.633Mi ± 0%       ~ (p=0.259 n=15)     7.633Mi ± 0%       ~ (p=0.733 n=15)     7.633Mi ± 0%  +0.00% (p=0.000 n=15)
Select/n=1000000/k=100/mostly_sorted-16    7.633Mi ± 0%   7.633Mi ± 0%       ~ (p=0.449 n=15)     7.633Mi ± 0%       ~ (p=0.881 n=15)     7.633Mi ± 0%  +0.00% (p=0.000 n=15)
Select/n=1000000/k=1000/random-16          7.633Mi ± 0%   7.633Mi ± 0%       ~ (p=0.078 n=15)     7.633Mi ± 0%       ~ (p=0.108 n=15)     7.633Mi ± 0%  +0.00% (p=0.000 n=15)
Select/n=1000000/k=1000/sorted-16          7.633Mi ± 0%   7.633Mi ± 0%       ~ (p=0.272 n=15)     7.633Mi ± 0%       ~ (p=0.338 n=15)     7.633Mi ± 0%  +0.00% (p=0.000 n=15)
Select/n=1000000/k=1000/reversed-16        7.633Mi ± 0%   7.633Mi ± 0%   0.00% (p=0.009 n=15)     7.633Mi ± 0%       ~ (p=0.059 n=15)     7.633Mi ± 0%  +0.00% (p=0.000 n=15)
Select/n=1000000/k=1000/mostly_sorted-16   7.633Mi ± 0%   7.633Mi ± 0%       ~ (p=0.112 n=15)     7.633Mi ± 0%       ~ (p=0.112 n=15)     7.633Mi ± 0%  +0.00% (p=0.000 n=15)
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
