# pdqselect: Pattern-Defeating QuickSelect for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/tsenart/pdqselect.svg)](https://pkg.go.dev/github.com/tsenart/pdqselect)
[![Go Report Card](https://goreportcard.com/badge/github.com/tsenart/pdqselect)](https://goreportcard.com/report/github.com/tsenart/pdqselect)

`pdqselect` is a high-performance, adaptive selection algorithm for Go that finds the k-th smallest elements in an ordered data structure. It's based on Go's internal `pdqsort` implementation, combining the best aspects of quicksort, insertion sort, and heapsort with pattern-defeating techniques to achieve optimal performance across a wide range of data distributions.

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
                                         │     Sort      │                Select                │              Ordered               │                 Func                 │
                                         │    sec/op     │    sec/op     vs base                │   sec/op     vs base               │    sec/op     vs base                │
Select/n=1000000/k=1/random-16             57.451m ±  1%   11.919m ± 1%   -79.25% (p=0.002 n=6)   5.992m ± 1%  -89.57% (p=0.002 n=6)    9.955m ± 1%   -82.67% (p=0.002 n=6)
Select/n=1000000/k=1/sorted-16              726.3µ ±  2%   2181.0µ ± 1%  +200.28% (p=0.002 n=6)   721.3µ ± 3%        ~ (p=0.589 n=6)   2095.2µ ± 1%  +188.47% (p=0.002 n=6)
Select/n=1000000/k=1/reversed-16            933.0µ ±  4%   3001.1µ ± 1%  +221.65% (p=0.002 n=6)   939.1µ ± 1%        ~ (p=0.394 n=6)   2269.7µ ± 2%  +143.26% (p=0.002 n=6)
Select/n=1000000/k=1/mostly_sorted-16      18.558m ±  0%    3.991m ± 1%   -78.49% (p=0.002 n=6)   1.339m ± 2%  -92.78% (p=0.002 n=6)    3.463m ± 1%   -81.34% (p=0.002 n=6)
Select/n=1000000/k=100/random-16           57.763m ±  1%    8.934m ± 1%   -84.53% (p=0.002 n=6)   4.826m ± 2%  -91.64% (p=0.002 n=6)    7.857m ± 1%   -86.40% (p=0.002 n=6)
Select/n=1000000/k=100/sorted-16            732.9µ ±  3%   2161.3µ ± 2%  +194.90% (p=0.002 n=6)   729.0µ ± 2%        ~ (p=0.818 n=6)   2111.5µ ± 1%  +188.10% (p=0.002 n=6)
Select/n=1000000/k=100/reversed-16          972.1µ ±  3%   3041.4µ ± 1%  +212.88% (p=0.002 n=6)   961.5µ ± 3%        ~ (p=1.000 n=6)   2289.1µ ± 2%  +135.49% (p=0.002 n=6)
Select/n=1000000/k=100/mostly_sorted-16    19.365m ±  1%    5.067m ± 1%   -73.84% (p=0.002 n=6)   1.818m ± 2%  -90.61% (p=0.002 n=6)    4.386m ± 2%   -77.35% (p=0.002 n=6)
Select/n=1000000/k=1000/random-16          58.366m ±  1%    9.013m ± 1%   -84.56% (p=0.002 n=6)   4.739m ± 2%  -91.88% (p=0.002 n=6)    7.885m ± 1%   -86.49% (p=0.002 n=6)
Select/n=1000000/k=1000/sorted-16           737.3µ ±  3%   2164.1µ ± 1%  +193.52% (p=0.002 n=6)   716.8µ ± 5%        ~ (p=0.485 n=6)   2141.6µ ± 2%  +190.46% (p=0.002 n=6)
Select/n=1000000/k=1000/reversed-16         953.8µ ±  1%   3068.9µ ± 1%  +221.76% (p=0.002 n=6)   972.7µ ± 1%   +1.98% (p=0.041 n=6)   2306.0µ ± 1%  +141.77% (p=0.002 n=6)
Select/n=1000000/k=1000/mostly_sorted-16   18.418m ±  2%    4.050m ± 0%   -78.01% (p=0.002 n=6)   1.337m ± 2%  -92.74% (p=0.002 n=6)    3.503m ± 1%   -80.98% (p=0.002 n=6)
Select/n=10000/k=1/random-16               150.87µ ± 10%    47.18µ ± 0%   -68.73% (p=0.002 n=6)   19.33µ ± 1%  -87.19% (p=0.002 n=6)    54.13µ ± 2%   -64.12% (p=0.002 n=6)
Select/n=10000/k=1/sorted-16                9.668µ ±  1%   24.861µ ± 0%  +157.15% (p=0.002 n=6)   9.716µ ± 1%        ~ (p=0.240 n=6)   24.174µ ± 0%  +150.04% (p=0.002 n=6)
Select/n=10000/k=1/reversed-16              11.93µ ±  0%    32.42µ ± 0%  +171.76% (p=0.002 n=6)   11.94µ ± 0%        ~ (p=0.699 n=6)    26.23µ ± 0%  +119.83% (p=0.002 n=6)
Select/n=10000/k=1/mostly_sorted-16        110.59µ ±  2%    41.21µ ± 0%   -62.74% (p=0.002 n=6)   12.51µ ± 0%  -88.68% (p=0.002 n=6)    37.31µ ± 0%   -66.26% (p=0.002 n=6)
Select/n=10000/k=100/random-16             152.40µ ±  5%    49.96µ ± 1%   -67.22% (p=0.002 n=6)   20.31µ ± 0%  -86.67% (p=0.002 n=6)    60.04µ ± 7%   -60.60% (p=0.002 n=6)
Select/n=10000/k=100/sorted-16              9.979µ ±  3%   24.533µ ± 3%  +145.85% (p=0.002 n=6)   9.699µ ± 0%   -2.80% (p=0.004 n=6)   23.301µ ± 1%  +133.51% (p=0.002 n=6)
Select/n=10000/k=100/reversed-16            11.77µ ±  1%    31.59µ ± 1%  +168.28% (p=0.002 n=6)   11.72µ ± 0%   -0.42% (p=0.009 n=6)    25.17µ ± 0%  +113.75% (p=0.002 n=6)
Select/n=10000/k=100/mostly_sorted-16       76.43µ ±  1%    40.54µ ± 1%   -46.96% (p=0.002 n=6)   12.44µ ± 1%  -83.73% (p=0.002 n=6)    35.03µ ± 2%   -54.17% (p=0.002 n=6)
Select/n=10000/k=1000/random-16            154.94µ ± 10%    53.02µ ± 2%   -65.78% (p=0.002 n=6)   21.09µ ± 1%  -86.39% (p=0.002 n=6)    59.30µ ± 2%   -61.73% (p=0.002 n=6)
Select/n=10000/k=1000/sorted-16             9.743µ ±  1%   24.469µ ± 1%  +151.14% (p=0.002 n=6)   9.770µ ± 1%        ~ (p=0.331 n=6)   23.377µ ± 1%  +139.94% (p=0.002 n=6)
Select/n=10000/k=1000/reversed-16           11.80µ ±  0%    31.74µ ± 1%  +168.99% (p=0.002 n=6)   11.87µ ± 1%        ~ (p=0.087 n=6)    25.59µ ± 1%  +116.84% (p=0.002 n=6)
Select/n=10000/k=1000/mostly_sorted-16      86.10µ ±  1%    41.30µ ± 0%   -52.03% (p=0.002 n=6)   12.43µ ± 1%  -85.56% (p=0.002 n=6)    35.38µ ± 1%   -58.91% (p=0.002 n=6)
Select/n=100/k=1/random-16                  673.8n ±  1%    568.7n ± 0%   -15.60% (p=0.002 n=6)   233.8n ± 1%  -65.30% (p=0.002 n=6)    555.2n ± 1%   -17.62% (p=0.002 n=6)
Select/n=100/k=1/sorted-16                  131.4n ±  0%    322.0n ± 1%  +145.02% (p=0.002 n=6)   133.5n ± 1%   +1.60% (p=0.002 n=6)    288.9n ± 1%  +119.86% (p=0.002 n=6)
Select/n=100/k=1/reversed-16                155.6n ±  1%    400.7n ± 0%  +157.40% (p=0.002 n=6)   158.9n ± 0%   +2.06% (p=0.002 n=6)    308.0n ± 0%   +97.88% (p=0.002 n=6)
Select/n=100/k=1/mostly_sorted-16           585.6n ±  2%    466.2n ± 0%   -20.40% (p=0.002 n=6)   177.1n ± 0%  -69.77% (p=0.002 n=6)    444.4n ± 1%   -24.12% (p=0.002 n=6)
Select/n=100/k=100/random-16                663.7n ±  1%    704.4n ± 0%    +6.13% (p=0.002 n=6)   281.5n ± 1%  -57.59% (p=0.002 n=6)    684.8n ± 1%    +3.17% (p=0.002 n=6)
Select/n=100/k=100/sorted-16                132.7n ±  1%    322.6n ± 1%  +143.10% (p=0.002 n=6)   133.8n ± 1%   +0.83% (p=0.006 n=6)    290.3n ± 0%  +118.80% (p=0.002 n=6)
Select/n=100/k=100/reversed-16              155.9n ±  0%    401.1n ± 0%  +157.25% (p=0.002 n=6)   159.1n ± 0%   +2.08% (p=0.002 n=6)    309.2n ± 0%   +98.36% (p=0.002 n=6)
Select/n=100/k=100/mostly_sorted-16         495.7n ±  1%    583.3n ± 0%   +17.69% (p=0.002 n=6)   221.4n ± 1%  -55.32% (p=0.002 n=6)    595.6n ± 0%   +20.17% (p=0.002 n=6)
geomean                                     69.71µ          70.96µ         +1.80%                 26.66µ       -61.76%                  63.89µ         -8.35%

                                         │     Sort     │               Select               │               Ordered                │                 Func                 │
                                         │     B/op     │     B/op      vs base              │     B/op      vs base                │     B/op      vs base                │
Select/n=1000000/k=1/random-16             7.633Mi ± 0%   7.633Mi ± 0%       ~ (p=0.063 n=6)   7.633Mi ± 0%       ~ (p=0.753 n=6)     7.633Mi ± 0%       ~ (p=0.556 n=6)
Select/n=1000000/k=1/sorted-16             7.633Mi ± 0%   7.633Mi ± 0%  +0.00% (p=0.002 n=6)   7.633Mi ± 0%       ~ (p=0.515 n=6)     7.633Mi ± 0%       ~ (p=0.316 n=6)
Select/n=1000000/k=1/reversed-16           7.633Mi ± 0%   7.633Mi ± 0%  +0.00% (p=0.002 n=6)   7.633Mi ± 0%       ~ (p=1.000 n=6)     7.633Mi ± 0%       ~ (p=1.000 n=6)
Select/n=1000000/k=1/mostly_sorted-16      7.633Mi ± 0%   7.633Mi ± 0%  +0.00% (p=0.002 n=6)   7.633Mi ± 0%       ~ (p=0.290 n=6)     7.633Mi ± 0%       ~ (p=0.316 n=6)
Select/n=1000000/k=100/random-16           7.633Mi ± 0%   7.633Mi ± 0%  +0.00% (p=0.002 n=6)   7.633Mi ± 0%  +0.00% (p=0.026 n=6)     7.633Mi ± 0%       ~ (p=0.113 n=6)
Select/n=1000000/k=100/sorted-16           7.633Mi ± 0%   7.633Mi ± 0%  +0.00% (p=0.002 n=6)   7.633Mi ± 0%       ~ (p=1.000 n=6)     7.633Mi ± 0%       ~ (p=0.636 n=6)
Select/n=1000000/k=100/reversed-16         7.633Mi ± 0%   7.633Mi ± 0%  +0.00% (p=0.002 n=6)   7.633Mi ± 0%       ~ (p=0.545 n=6)     7.633Mi ± 0%       ~ (p=0.545 n=6)
Select/n=1000000/k=100/mostly_sorted-16    7.633Mi ± 0%   7.633Mi ± 0%  +0.00% (p=0.002 n=6)   7.633Mi ± 0%       ~ (p=0.123 n=6)     7.633Mi ± 0%       ~ (p=0.123 n=6)
Select/n=1000000/k=1000/random-16          7.633Mi ± 0%   7.633Mi ± 0%  +0.00% (p=0.002 n=6)   7.633Mi ± 0%       ~ (p=0.394 n=6)     7.633Mi ± 0%       ~ (p=0.351 n=6)
Select/n=1000000/k=1000/sorted-16          7.633Mi ± 0%   7.633Mi ± 0%  +0.00% (p=0.002 n=6)   7.633Mi ± 0%       ~ (p=1.000 n=6)     7.633Mi ± 0%       ~ (p=0.061 n=6)
Select/n=1000000/k=1000/reversed-16        7.633Mi ± 0%   7.633Mi ± 0%  +0.00% (p=0.002 n=6)   7.633Mi ± 0%       ~ (p=0.182 n=6)     7.633Mi ± 0%       ~ (p=0.182 n=6)
Select/n=1000000/k=1000/mostly_sorted-16   7.633Mi ± 0%   7.633Mi ± 0%  +0.00% (p=0.002 n=6)   7.633Mi ± 0%       ~ (p=0.537 n=6)     7.633Mi ± 0%       ~ (p=0.509 n=6)
Select/n=10000/k=1/random-16               80.00Ki ± 0%   80.02Ki ± 0%  +0.03% (p=0.002 n=6)   80.00Ki ± 0%       ~ (p=1.000 n=6) ¹   80.00Ki ± 0%       ~ (p=1.000 n=6) ¹
Select/n=10000/k=1/sorted-16               80.00Ki ± 0%   80.02Ki ± 0%  +0.03% (p=0.002 n=6)   80.00Ki ± 0%       ~ (p=1.000 n=6) ¹   80.00Ki ± 0%       ~ (p=1.000 n=6) ¹
Select/n=10000/k=1/reversed-16             80.00Ki ± 0%   80.02Ki ± 0%  +0.03% (p=0.002 n=6)   80.00Ki ± 0%       ~ (p=1.000 n=6) ¹   80.00Ki ± 0%       ~ (p=1.000 n=6) ¹
Select/n=10000/k=1/mostly_sorted-16        80.00Ki ± 0%   80.02Ki ± 0%  +0.03% (p=0.002 n=6)   80.00Ki ± 0%       ~ (p=1.000 n=6) ¹   80.00Ki ± 0%       ~ (p=1.000 n=6) ¹
Select/n=10000/k=100/random-16             80.00Ki ± 0%   80.02Ki ± 0%  +0.03% (p=0.002 n=6)   80.00Ki ± 0%       ~ (p=1.000 n=6) ¹   80.00Ki ± 0%       ~ (p=1.000 n=6) ¹
Select/n=10000/k=100/sorted-16             80.00Ki ± 0%   80.02Ki ± 0%  +0.03% (p=0.002 n=6)   80.00Ki ± 0%       ~ (p=1.000 n=6) ¹   80.00Ki ± 0%       ~ (p=1.000 n=6) ¹
Select/n=10000/k=100/reversed-16           80.00Ki ± 0%   80.02Ki ± 0%  +0.03% (p=0.002 n=6)   80.00Ki ± 0%       ~ (p=1.000 n=6) ¹   80.00Ki ± 0%       ~ (p=1.000 n=6) ¹
Select/n=10000/k=100/mostly_sorted-16      80.00Ki ± 0%   80.02Ki ± 0%  +0.03% (p=0.002 n=6)   80.00Ki ± 0%       ~ (p=1.000 n=6) ¹   80.00Ki ± 0%       ~ (p=1.000 n=6) ¹
Select/n=10000/k=1000/random-16            80.00Ki ± 0%   80.02Ki ± 0%  +0.03% (p=0.002 n=6)   80.00Ki ± 0%       ~ (p=1.000 n=6) ¹   80.00Ki ± 0%       ~ (p=1.000 n=6) ¹
Select/n=10000/k=1000/sorted-16            80.00Ki ± 0%   80.02Ki ± 0%  +0.03% (p=0.002 n=6)   80.00Ki ± 0%       ~ (p=1.000 n=6) ¹   80.00Ki ± 0%       ~ (p=1.000 n=6) ¹
Select/n=10000/k=1000/reversed-16          80.00Ki ± 0%   80.02Ki ± 0%  +0.03% (p=0.002 n=6)   80.00Ki ± 0%       ~ (p=1.000 n=6) ¹   80.00Ki ± 0%       ~ (p=1.000 n=6) ¹
Select/n=10000/k=1000/mostly_sorted-16     80.00Ki ± 0%   80.02Ki ± 0%  +0.03% (p=0.002 n=6)   80.00Ki ± 0%       ~ (p=1.000 n=6) ¹   80.00Ki ± 0%       ~ (p=1.000 n=6) ¹
Select/n=100/k=1/random-16                   896.0 ± 0%     920.0 ± 0%  +2.68% (p=0.002 n=6)     896.0 ± 0%       ~ (p=1.000 n=6) ¹     896.0 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=100/k=1/sorted-16                   896.0 ± 0%     920.0 ± 0%  +2.68% (p=0.002 n=6)     896.0 ± 0%       ~ (p=1.000 n=6) ¹     896.0 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=100/k=1/reversed-16                 896.0 ± 0%     920.0 ± 0%  +2.68% (p=0.002 n=6)     896.0 ± 0%       ~ (p=1.000 n=6) ¹     896.0 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=100/k=1/mostly_sorted-16            896.0 ± 0%     920.0 ± 0%  +2.68% (p=0.002 n=6)     896.0 ± 0%       ~ (p=1.000 n=6) ¹     896.0 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=100/k=100/random-16                 896.0 ± 0%     920.0 ± 0%  +2.68% (p=0.002 n=6)     896.0 ± 0%       ~ (p=1.000 n=6) ¹     896.0 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=100/k=100/sorted-16                 896.0 ± 0%     920.0 ± 0%  +2.68% (p=0.002 n=6)     896.0 ± 0%       ~ (p=1.000 n=6) ¹     896.0 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=100/k=100/reversed-16               896.0 ± 0%     920.0 ± 0%  +2.68% (p=0.002 n=6)     896.0 ± 0%       ~ (p=1.000 n=6) ¹     896.0 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=100/k=100/mostly_sorted-16          896.0 ± 0%     920.0 ± 0%  +2.68% (p=0.002 n=6)     896.0 ± 0%       ~ (p=1.000 n=6) ¹     896.0 ± 0%       ~ (p=1.000 n=6) ¹
geomean                                    144.2Ki        145.2Ki       +0.67%                 144.2Ki       +0.00%                   144.2Ki       +0.00%
¹ all samples are equal

                                         │    Sort    │               Select               │              Ordered               │                Func                │
                                         │ allocs/op  │ allocs/op   vs base                │ allocs/op   vs base                │ allocs/op   vs base                │
Select/n=1000000/k=1/random-16             1.000 ± 0%   2.000 ± 0%  +100.00% (p=0.002 n=6)   1.000 ± 0%       ~ (p=1.000 n=6) ¹   1.000 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=1000000/k=1/sorted-16             1.000 ± 0%   2.000 ± 0%  +100.00% (p=0.002 n=6)   1.000 ± 0%       ~ (p=1.000 n=6) ¹   1.000 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=1000000/k=1/reversed-16           1.000 ± 0%   2.000 ± 0%  +100.00% (p=0.002 n=6)   1.000 ± 0%       ~ (p=1.000 n=6) ¹   1.000 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=1000000/k=1/mostly_sorted-16      1.000 ± 0%   2.000 ± 0%  +100.00% (p=0.002 n=6)   1.000 ± 0%       ~ (p=1.000 n=6) ¹   1.000 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=1000000/k=100/random-16           1.000 ± 0%   2.000 ± 0%  +100.00% (p=0.002 n=6)   1.000 ± 0%       ~ (p=1.000 n=6) ¹   1.000 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=1000000/k=100/sorted-16           1.000 ± 0%   2.000 ± 0%  +100.00% (p=0.002 n=6)   1.000 ± 0%       ~ (p=1.000 n=6) ¹   1.000 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=1000000/k=100/reversed-16         1.000 ± 0%   2.000 ± 0%  +100.00% (p=0.002 n=6)   1.000 ± 0%       ~ (p=1.000 n=6) ¹   1.000 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=1000000/k=100/mostly_sorted-16    1.000 ± 0%   2.000 ± 0%  +100.00% (p=0.002 n=6)   1.000 ± 0%       ~ (p=1.000 n=6) ¹   1.000 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=1000000/k=1000/random-16          1.000 ± 0%   2.000 ± 0%  +100.00% (p=0.002 n=6)   1.000 ± 0%       ~ (p=1.000 n=6) ¹   1.000 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=1000000/k=1000/sorted-16          1.000 ± 0%   2.000 ± 0%  +100.00% (p=0.002 n=6)   1.000 ± 0%       ~ (p=1.000 n=6) ¹   1.000 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=1000000/k=1000/reversed-16        1.000 ± 0%   2.000 ± 0%  +100.00% (p=0.002 n=6)   1.000 ± 0%       ~ (p=1.000 n=6) ¹   1.000 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=1000000/k=1000/mostly_sorted-16   1.000 ± 0%   2.000 ± 0%  +100.00% (p=0.002 n=6)   1.000 ± 0%       ~ (p=1.000 n=6) ¹   1.000 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=10000/k=1/random-16               1.000 ± 0%   2.000 ± 0%  +100.00% (p=0.002 n=6)   1.000 ± 0%       ~ (p=1.000 n=6) ¹   1.000 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=10000/k=1/sorted-16               1.000 ± 0%   2.000 ± 0%  +100.00% (p=0.002 n=6)   1.000 ± 0%       ~ (p=1.000 n=6) ¹   1.000 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=10000/k=1/reversed-16             1.000 ± 0%   2.000 ± 0%  +100.00% (p=0.002 n=6)   1.000 ± 0%       ~ (p=1.000 n=6) ¹   1.000 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=10000/k=1/mostly_sorted-16        1.000 ± 0%   2.000 ± 0%  +100.00% (p=0.002 n=6)   1.000 ± 0%       ~ (p=1.000 n=6) ¹   1.000 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=10000/k=100/random-16             1.000 ± 0%   2.000 ± 0%  +100.00% (p=0.002 n=6)   1.000 ± 0%       ~ (p=1.000 n=6) ¹   1.000 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=10000/k=100/sorted-16             1.000 ± 0%   2.000 ± 0%  +100.00% (p=0.002 n=6)   1.000 ± 0%       ~ (p=1.000 n=6) ¹   1.000 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=10000/k=100/reversed-16           1.000 ± 0%   2.000 ± 0%  +100.00% (p=0.002 n=6)   1.000 ± 0%       ~ (p=1.000 n=6) ¹   1.000 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=10000/k=100/mostly_sorted-16      1.000 ± 0%   2.000 ± 0%  +100.00% (p=0.002 n=6)   1.000 ± 0%       ~ (p=1.000 n=6) ¹   1.000 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=10000/k=1000/random-16            1.000 ± 0%   2.000 ± 0%  +100.00% (p=0.002 n=6)   1.000 ± 0%       ~ (p=1.000 n=6) ¹   1.000 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=10000/k=1000/sorted-16            1.000 ± 0%   2.000 ± 0%  +100.00% (p=0.002 n=6)   1.000 ± 0%       ~ (p=1.000 n=6) ¹   1.000 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=10000/k=1000/reversed-16          1.000 ± 0%   2.000 ± 0%  +100.00% (p=0.002 n=6)   1.000 ± 0%       ~ (p=1.000 n=6) ¹   1.000 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=10000/k=1000/mostly_sorted-16     1.000 ± 0%   2.000 ± 0%  +100.00% (p=0.002 n=6)   1.000 ± 0%       ~ (p=1.000 n=6) ¹   1.000 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=100/k=1/random-16                 1.000 ± 0%   2.000 ± 0%  +100.00% (p=0.002 n=6)   1.000 ± 0%       ~ (p=1.000 n=6) ¹   1.000 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=100/k=1/sorted-16                 1.000 ± 0%   2.000 ± 0%  +100.00% (p=0.002 n=6)   1.000 ± 0%       ~ (p=1.000 n=6) ¹   1.000 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=100/k=1/reversed-16               1.000 ± 0%   2.000 ± 0%  +100.00% (p=0.002 n=6)   1.000 ± 0%       ~ (p=1.000 n=6) ¹   1.000 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=100/k=1/mostly_sorted-16          1.000 ± 0%   2.000 ± 0%  +100.00% (p=0.002 n=6)   1.000 ± 0%       ~ (p=1.000 n=6) ¹   1.000 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=100/k=100/random-16               1.000 ± 0%   2.000 ± 0%  +100.00% (p=0.002 n=6)   1.000 ± 0%       ~ (p=1.000 n=6) ¹   1.000 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=100/k=100/sorted-16               1.000 ± 0%   2.000 ± 0%  +100.00% (p=0.002 n=6)   1.000 ± 0%       ~ (p=1.000 n=6) ¹   1.000 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=100/k=100/reversed-16             1.000 ± 0%   2.000 ± 0%  +100.00% (p=0.002 n=6)   1.000 ± 0%       ~ (p=1.000 n=6) ¹   1.000 ± 0%       ~ (p=1.000 n=6) ¹
Select/n=100/k=100/mostly_sorted-16        1.000 ± 0%   2.000 ± 0%  +100.00% (p=0.002 n=6)   1.000 ± 0%       ~ (p=1.000 n=6) ¹   1.000 ± 0%       ~ (p=1.000 n=6) ¹
geomean                                    1.000        2.000       +100.00%                 1.000       +0.00%                   1.000       +0.00%
¹ all samples are equal
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built upon Go's internal `pdqsort` algorithm.
