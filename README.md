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
                                                │    bench     │
                                                │    sec/op    │
Select/Select/n=100000/k=1000/random-16           1.109m ± ∞ ¹
Select/Ordered/n=100000/k=1000/random-16          438.3µ ± ∞ ¹
Select/Func/n=100000/k=1000/random-16             740.7µ ± ∞ ¹
Select/Sort/n=100000/k=1000/random-16             4.602m ± ∞ ¹
Select/Select/n=100000/k=1000/sorted-16           243.4µ ± ∞ ¹
Select/Ordered/n=100000/k=1000/sorted-16          100.0µ ± ∞ ¹
Select/Func/n=100000/k=1000/sorted-16             233.0µ ± ∞ ¹
Select/Sort/n=100000/k=1000/sorted-16             100.5µ ± ∞ ¹
Select/Select/n=100000/k=1000/reversed-16         322.3µ ± ∞ ¹
Select/Ordered/n=100000/k=1000/reversed-16        120.9µ ± ∞ ¹
Select/Func/n=100000/k=1000/reversed-16           248.5µ ± ∞ ¹
Select/Sort/n=100000/k=1000/reversed-16           118.4µ ± ∞ ¹
Select/Select/n=100000/k=1000/equal-16            244.4µ ± ∞ ¹
Select/Ordered/n=100000/k=1000/equal-16           99.21µ ± ∞ ¹
Select/Func/n=100000/k=1000/equal-16              228.2µ ± ∞ ¹
Select/Sort/n=100000/k=1000/equal-16              99.99µ ± ∞ ¹
Select/Select/n=100000/k=1000/mostly_equal-16     525.3µ ± ∞ ¹
Select/Ordered/n=100000/k=1000/mostly_equal-16    156.1µ ± ∞ ¹
Select/Func/n=100000/k=1000/mostly_equal-16       516.7µ ± ∞ ¹
Select/Sort/n=100000/k=1000/mostly_equal-16       323.6µ ± ∞ ¹
Select/Select/n=100000/k=5000/random-16           767.2µ ± ∞ ¹
Select/Ordered/n=100000/k=5000/random-16          375.0µ ± ∞ ¹
Select/Func/n=100000/k=5000/random-16             1.004m ± ∞ ¹
Select/Sort/n=100000/k=5000/random-16             4.654m ± ∞ ¹
Select/Select/n=100000/k=5000/sorted-16           243.8µ ± ∞ ¹
Select/Ordered/n=100000/k=5000/sorted-16          100.9µ ± ∞ ¹
Select/Func/n=100000/k=5000/sorted-16             232.2µ ± ∞ ¹
Select/Sort/n=100000/k=5000/sorted-16             99.97µ ± ∞ ¹
Select/Select/n=100000/k=5000/reversed-16         322.4µ ± ∞ ¹
Select/Ordered/n=100000/k=5000/reversed-16        118.2µ ± ∞ ¹
Select/Func/n=100000/k=5000/reversed-16           250.0µ ± ∞ ¹
Select/Sort/n=100000/k=5000/reversed-16           121.3µ ± ∞ ¹
Select/Select/n=100000/k=5000/equal-16            246.5µ ± ∞ ¹
Select/Ordered/n=100000/k=5000/equal-16           100.2µ ± ∞ ¹
Select/Func/n=100000/k=5000/equal-16              229.8µ ± ∞ ¹
Select/Sort/n=100000/k=5000/equal-16              99.95µ ± ∞ ¹
Select/Select/n=100000/k=5000/mostly_equal-16     502.4µ ± ∞ ¹
Select/Ordered/n=100000/k=5000/mostly_equal-16    152.6µ ± ∞ ¹
Select/Func/n=100000/k=5000/mostly_equal-16       537.3µ ± ∞ ¹
Select/Sort/n=100000/k=5000/mostly_equal-16       323.9µ ± ∞ ¹
Select/Select/n=100000/k=10000/random-16          850.1µ ± ∞ ¹
Select/Ordered/n=100000/k=10000/random-16         357.5µ ± ∞ ¹
Select/Func/n=100000/k=10000/random-16            857.7µ ± ∞ ¹
Select/Sort/n=100000/k=10000/random-16            4.698m ± ∞ ¹
Select/Select/n=100000/k=10000/sorted-16          244.6µ ± ∞ ¹
Select/Ordered/n=100000/k=10000/sorted-16         100.5µ ± ∞ ¹
Select/Func/n=100000/k=10000/sorted-16            230.5µ ± ∞ ¹
Select/Sort/n=100000/k=10000/sorted-16            100.5µ ± ∞ ¹
Select/Select/n=100000/k=10000/reversed-16        321.8µ ± ∞ ¹
Select/Ordered/n=100000/k=10000/reversed-16       121.3µ ± ∞ ¹
Select/Func/n=100000/k=10000/reversed-16          249.8µ ± ∞ ¹
Select/Sort/n=100000/k=10000/reversed-16          118.9µ ± ∞ ¹
Select/Select/n=100000/k=10000/equal-16           247.1µ ± ∞ ¹
Select/Ordered/n=100000/k=10000/equal-16          100.8µ ± ∞ ¹
Select/Func/n=100000/k=10000/equal-16             232.1µ ± ∞ ¹
Select/Sort/n=100000/k=10000/equal-16             100.2µ ± ∞ ¹
Select/Select/n=100000/k=10000/mostly_equal-16    533.0µ ± ∞ ¹
Select/Ordered/n=100000/k=10000/mostly_equal-16   159.0µ ± ∞ ¹
Select/Func/n=100000/k=10000/mostly_equal-16      558.8µ ± ∞ ¹
Select/Sort/n=100000/k=10000/mostly_equal-16      349.2µ ± ∞ ¹
Select/Select/n=10000/k=100/random-16             40.67µ ± ∞ ¹
Select/Ordered/n=10000/k=100/random-16            21.82µ ± ∞ ¹
Select/Func/n=10000/k=100/random-16               58.41µ ± ∞ ¹
Select/Sort/n=10000/k=100/random-16               173.7µ ± ∞ ¹
Select/Select/n=10000/k=100/sorted-16             24.40µ ± ∞ ¹
Select/Ordered/n=10000/k=100/sorted-16            9.727µ ± ∞ ¹
Select/Func/n=10000/k=100/sorted-16               23.68µ ± ∞ ¹
Select/Sort/n=10000/k=100/sorted-16               9.919µ ± ∞ ¹
Select/Select/n=10000/k=100/reversed-16           32.03µ ± ∞ ¹
Select/Ordered/n=10000/k=100/reversed-16          11.93µ ± ∞ ¹
Select/Func/n=10000/k=100/reversed-16             25.82µ ± ∞ ¹
Select/Sort/n=10000/k=100/reversed-16             11.94µ ± ∞ ¹
Select/Select/n=10000/k=100/equal-16              24.77µ ± ∞ ¹
Select/Ordered/n=10000/k=100/equal-16             9.904µ ± ∞ ¹
Select/Func/n=10000/k=100/equal-16                23.45µ ± ∞ ¹
Select/Sort/n=10000/k=100/equal-16                9.787µ ± ∞ ¹
Select/Select/n=10000/k=100/mostly_equal-16       52.41µ ± ∞ ¹
Select/Ordered/n=10000/k=100/mostly_equal-16      14.79µ ± ∞ ¹
Select/Func/n=10000/k=100/mostly_equal-16         51.93µ ± ∞ ¹
Select/Sort/n=10000/k=100/mostly_equal-16         22.94µ ± ∞ ¹
Select/Select/n=10000/k=500/random-16             60.96µ ± ∞ ¹
Select/Ordered/n=10000/k=500/random-16            19.98µ ± ∞ ¹
Select/Func/n=10000/k=500/random-16               56.48µ ± ∞ ¹
Select/Sort/n=10000/k=500/random-16               172.7µ ± ∞ ¹
Select/Select/n=10000/k=500/sorted-16             24.74µ ± ∞ ¹
Select/Ordered/n=10000/k=500/sorted-16            10.03µ ± ∞ ¹
Select/Func/n=10000/k=500/sorted-16               23.85µ ± ∞ ¹
Select/Sort/n=10000/k=500/sorted-16               9.902µ ± ∞ ¹
Select/Select/n=10000/k=500/reversed-16           32.25µ ± ∞ ¹
Select/Ordered/n=10000/k=500/reversed-16          12.00µ ± ∞ ¹
Select/Func/n=10000/k=500/reversed-16             26.48µ ± ∞ ¹
Select/Sort/n=10000/k=500/reversed-16             11.94µ ± ∞ ¹
Select/Select/n=10000/k=500/equal-16              24.97µ ± ∞ ¹
Select/Ordered/n=10000/k=500/equal-16             9.679µ ± ∞ ¹
Select/Func/n=10000/k=500/equal-16                24.14µ ± ∞ ¹
Select/Sort/n=10000/k=500/equal-16                9.542µ ± ∞ ¹
Select/Select/n=10000/k=500/mostly_equal-16       53.27µ ± ∞ ¹
Select/Ordered/n=10000/k=500/mostly_equal-16      15.04µ ± ∞ ¹
Select/Func/n=10000/k=500/mostly_equal-16         52.19µ ± ∞ ¹
Select/Sort/n=10000/k=500/mostly_equal-16         23.19µ ± ∞ ¹
Select/Select/n=10000/k=1000/random-16            56.20µ ± ∞ ¹
Select/Ordered/n=10000/k=1000/random-16           24.28µ ± ∞ ¹
Select/Func/n=10000/k=1000/random-16              60.06µ ± ∞ ¹
Select/Sort/n=10000/k=1000/random-16              163.1µ ± ∞ ¹
Select/Select/n=10000/k=1000/sorted-16            24.98µ ± ∞ ¹
Select/Ordered/n=10000/k=1000/sorted-16           9.726µ ± ∞ ¹
Select/Func/n=10000/k=1000/sorted-16              24.16µ ± ∞ ¹
Select/Sort/n=10000/k=1000/sorted-16              9.896µ ± ∞ ¹
Select/Select/n=10000/k=1000/reversed-16          32.30µ ± ∞ ¹
Select/Ordered/n=10000/k=1000/reversed-16         11.87µ ± ∞ ¹
Select/Func/n=10000/k=1000/reversed-16            26.31µ ± ∞ ¹
Select/Sort/n=10000/k=1000/reversed-16            11.92µ ± ∞ ¹
Select/Select/n=10000/k=1000/equal-16             24.88µ ± ∞ ¹
Select/Ordered/n=10000/k=1000/equal-16            9.717µ ± ∞ ¹
Select/Func/n=10000/k=1000/equal-16               24.16µ ± ∞ ¹
Select/Sort/n=10000/k=1000/equal-16               9.663µ ± ∞ ¹
Select/Select/n=10000/k=1000/mostly_equal-16      52.25µ ± ∞ ¹
Select/Ordered/n=10000/k=1000/mostly_equal-16     15.08µ ± ∞ ¹
Select/Func/n=10000/k=1000/mostly_equal-16        55.14µ ± ∞ ¹
Select/Sort/n=10000/k=1000/mostly_equal-16        22.73µ ± ∞ ¹
Select/Select/n=1000/k=10/random-16               5.975µ ± ∞ ¹
Select/Ordered/n=1000/k=10/random-16              1.983µ ± ∞ ¹
Select/Func/n=1000/k=10/random-16                 5.552µ ± ∞ ¹
Select/Sort/n=1000/k=10/random-16                 10.22µ ± ∞ ¹
Select/Select/n=1000/k=10/sorted-16               2.640µ ± ∞ ¹
Select/Ordered/n=1000/k=10/sorted-16              1.118µ ± ∞ ¹
Select/Func/n=1000/k=10/sorted-16                 2.530µ ± ∞ ¹
Select/Sort/n=1000/k=10/sorted-16                 1.139µ ± ∞ ¹
Select/Select/n=1000/k=10/reversed-16             3.403µ ± ∞ ¹
Select/Ordered/n=1000/k=10/reversed-16            1.371µ ± ∞ ¹
Select/Func/n=1000/k=10/reversed-16               2.787µ ± ∞ ¹
Select/Sort/n=1000/k=10/reversed-16               1.365µ ± ∞ ¹
Select/Select/n=1000/k=10/equal-16                2.670µ ± ∞ ¹
Select/Ordered/n=1000/k=10/equal-16               1.123µ ± ∞ ¹
Select/Func/n=1000/k=10/equal-16                  2.542µ ± ∞ ¹
Select/Sort/n=1000/k=10/equal-16                  1.120µ ± ∞ ¹
Select/Select/n=1000/k=10/mostly_equal-16         9.048µ ± ∞ ¹
Select/Ordered/n=1000/k=10/mostly_equal-16        1.699µ ± ∞ ¹
Select/Func/n=1000/k=10/mostly_equal-16           5.554µ ± ∞ ¹
Select/Sort/n=1000/k=10/mostly_equal-16           2.479µ ± ∞ ¹
Select/Select/n=1000/k=50/random-16               5.498µ ± ∞ ¹
Select/Ordered/n=1000/k=50/random-16              1.843µ ± ∞ ¹
Select/Func/n=1000/k=50/random-16                 5.092µ ± ∞ ¹
Select/Sort/n=1000/k=50/random-16                 10.42µ ± ∞ ¹
Select/Select/n=1000/k=50/sorted-16               2.634µ ± ∞ ¹
Select/Ordered/n=1000/k=50/sorted-16              1.139µ ± ∞ ¹
Select/Func/n=1000/k=50/sorted-16                 2.512µ ± ∞ ¹
Select/Sort/n=1000/k=50/sorted-16                 1.108µ ± ∞ ¹
Select/Select/n=1000/k=50/reversed-16             3.377µ ± ∞ ¹
Select/Ordered/n=1000/k=50/reversed-16            1.366µ ± ∞ ¹
Select/Func/n=1000/k=50/reversed-16               2.781µ ± ∞ ¹
Select/Sort/n=1000/k=50/reversed-16               1.370µ ± ∞ ¹
Select/Select/n=1000/k=50/equal-16                2.641µ ± ∞ ¹
Select/Ordered/n=1000/k=50/equal-16               1.123µ ± ∞ ¹
Select/Func/n=1000/k=50/equal-16                  2.524µ ± ∞ ¹
Select/Sort/n=1000/k=50/equal-16                  1.103µ ± ∞ ¹
Select/Select/n=1000/k=50/mostly_equal-16         5.315µ ± ∞ ¹
Select/Ordered/n=1000/k=50/mostly_equal-16        2.088µ ± ∞ ¹
Select/Func/n=1000/k=50/mostly_equal-16           5.826µ ± ∞ ¹
Select/Sort/n=1000/k=50/mostly_equal-16           2.357µ ± ∞ ¹
Select/Select/n=1000/k=100/random-16              4.894µ ± ∞ ¹
Select/Ordered/n=1000/k=100/random-16             2.303µ ± ∞ ¹
Select/Func/n=1000/k=100/random-16                5.595µ ± ∞ ¹
Select/Sort/n=1000/k=100/random-16                10.35µ ± ∞ ¹
Select/Select/n=1000/k=100/sorted-16              2.638µ ± ∞ ¹
Select/Ordered/n=1000/k=100/sorted-16             1.128µ ± ∞ ¹
Select/Func/n=1000/k=100/sorted-16                2.526µ ± ∞ ¹
Select/Sort/n=1000/k=100/sorted-16                1.117µ ± ∞ ¹
Select/Select/n=1000/k=100/reversed-16            3.381µ ± ∞ ¹
Select/Ordered/n=1000/k=100/reversed-16           1.365µ ± ∞ ¹
Select/Func/n=1000/k=100/reversed-16              2.762µ ± ∞ ¹
Select/Sort/n=1000/k=100/reversed-16              1.363µ ± ∞ ¹
Select/Select/n=1000/k=100/equal-16               2.652µ ± ∞ ¹
Select/Ordered/n=1000/k=100/equal-16              1.127µ ± ∞ ¹
Select/Func/n=1000/k=100/equal-16                 2.530µ ± ∞ ¹
Select/Sort/n=1000/k=100/equal-16                 1.107µ ± ∞ ¹
Select/Select/n=1000/k=100/mostly_equal-16        6.335µ ± ∞ ¹
Select/Ordered/n=1000/k=100/mostly_equal-16       1.829µ ± ∞ ¹
Select/Func/n=1000/k=100/mostly_equal-16          6.150µ ± ∞ ¹
Select/Sort/n=1000/k=100/mostly_equal-16          2.264µ ± ∞ ¹
geomean                                           25.73µ
¹ need >= 6 samples for confidence interval at level 0.95

                                                │     bench     │
                                                │     B/op      │
Select/Select/n=100000/k=1000/random-16           784.0Ki ± ∞ ¹
Select/Ordered/n=100000/k=1000/random-16          784.0Ki ± ∞ ¹
Select/Func/n=100000/k=1000/random-16             784.0Ki ± ∞ ¹
Select/Sort/n=100000/k=1000/random-16             784.0Ki ± ∞ ¹
Select/Select/n=100000/k=1000/sorted-16           784.0Ki ± ∞ ¹
Select/Ordered/n=100000/k=1000/sorted-16          784.0Ki ± ∞ ¹
Select/Func/n=100000/k=1000/sorted-16             784.0Ki ± ∞ ¹
Select/Sort/n=100000/k=1000/sorted-16             784.0Ki ± ∞ ¹
Select/Select/n=100000/k=1000/reversed-16         784.0Ki ± ∞ ¹
Select/Ordered/n=100000/k=1000/reversed-16        784.0Ki ± ∞ ¹
Select/Func/n=100000/k=1000/reversed-16           784.0Ki ± ∞ ¹
Select/Sort/n=100000/k=1000/reversed-16           784.0Ki ± ∞ ¹
Select/Select/n=100000/k=1000/equal-16            784.0Ki ± ∞ ¹
Select/Ordered/n=100000/k=1000/equal-16           784.0Ki ± ∞ ¹
Select/Func/n=100000/k=1000/equal-16              784.0Ki ± ∞ ¹
Select/Sort/n=100000/k=1000/equal-16              784.0Ki ± ∞ ¹
Select/Select/n=100000/k=1000/mostly_equal-16     784.0Ki ± ∞ ¹
Select/Ordered/n=100000/k=1000/mostly_equal-16    784.0Ki ± ∞ ¹
Select/Func/n=100000/k=1000/mostly_equal-16       784.0Ki ± ∞ ¹
Select/Sort/n=100000/k=1000/mostly_equal-16       784.0Ki ± ∞ ¹
Select/Select/n=100000/k=5000/random-16           784.0Ki ± ∞ ¹
Select/Ordered/n=100000/k=5000/random-16          784.0Ki ± ∞ ¹
Select/Func/n=100000/k=5000/random-16             784.0Ki ± ∞ ¹
Select/Sort/n=100000/k=5000/random-16             784.0Ki ± ∞ ¹
Select/Select/n=100000/k=5000/sorted-16           784.0Ki ± ∞ ¹
Select/Ordered/n=100000/k=5000/sorted-16          784.0Ki ± ∞ ¹
Select/Func/n=100000/k=5000/sorted-16             784.0Ki ± ∞ ¹
Select/Sort/n=100000/k=5000/sorted-16             784.0Ki ± ∞ ¹
Select/Select/n=100000/k=5000/reversed-16         784.0Ki ± ∞ ¹
Select/Ordered/n=100000/k=5000/reversed-16        784.0Ki ± ∞ ¹
Select/Func/n=100000/k=5000/reversed-16           784.0Ki ± ∞ ¹
Select/Sort/n=100000/k=5000/reversed-16           784.0Ki ± ∞ ¹
Select/Select/n=100000/k=5000/equal-16            784.0Ki ± ∞ ¹
Select/Ordered/n=100000/k=5000/equal-16           784.0Ki ± ∞ ¹
Select/Func/n=100000/k=5000/equal-16              784.0Ki ± ∞ ¹
Select/Sort/n=100000/k=5000/equal-16              784.0Ki ± ∞ ¹
Select/Select/n=100000/k=5000/mostly_equal-16     784.0Ki ± ∞ ¹
Select/Ordered/n=100000/k=5000/mostly_equal-16    784.0Ki ± ∞ ¹
Select/Func/n=100000/k=5000/mostly_equal-16       784.0Ki ± ∞ ¹
Select/Sort/n=100000/k=5000/mostly_equal-16       784.0Ki ± ∞ ¹
Select/Select/n=100000/k=10000/random-16          784.0Ki ± ∞ ¹
Select/Ordered/n=100000/k=10000/random-16         784.0Ki ± ∞ ¹
Select/Func/n=100000/k=10000/random-16            784.0Ki ± ∞ ¹
Select/Sort/n=100000/k=10000/random-16            784.0Ki ± ∞ ¹
Select/Select/n=100000/k=10000/sorted-16          784.0Ki ± ∞ ¹
Select/Ordered/n=100000/k=10000/sorted-16         784.0Ki ± ∞ ¹
Select/Func/n=100000/k=10000/sorted-16            784.0Ki ± ∞ ¹
Select/Sort/n=100000/k=10000/sorted-16            784.0Ki ± ∞ ¹
Select/Select/n=100000/k=10000/reversed-16        784.0Ki ± ∞ ¹
Select/Ordered/n=100000/k=10000/reversed-16       784.0Ki ± ∞ ¹
Select/Func/n=100000/k=10000/reversed-16          784.0Ki ± ∞ ¹
Select/Sort/n=100000/k=10000/reversed-16          784.0Ki ± ∞ ¹
Select/Select/n=100000/k=10000/equal-16           784.0Ki ± ∞ ¹
Select/Ordered/n=100000/k=10000/equal-16          784.0Ki ± ∞ ¹
Select/Func/n=100000/k=10000/equal-16             784.0Ki ± ∞ ¹
Select/Sort/n=100000/k=10000/equal-16             784.0Ki ± ∞ ¹
Select/Select/n=100000/k=10000/mostly_equal-16    784.0Ki ± ∞ ¹
Select/Ordered/n=100000/k=10000/mostly_equal-16   784.0Ki ± ∞ ¹
Select/Func/n=100000/k=10000/mostly_equal-16      784.0Ki ± ∞ ¹
Select/Sort/n=100000/k=10000/mostly_equal-16      784.0Ki ± ∞ ¹
Select/Select/n=10000/k=100/random-16             80.02Ki ± ∞ ¹
Select/Ordered/n=10000/k=100/random-16            80.00Ki ± ∞ ¹
Select/Func/n=10000/k=100/random-16               80.00Ki ± ∞ ¹
Select/Sort/n=10000/k=100/random-16               80.00Ki ± ∞ ¹
Select/Select/n=10000/k=100/sorted-16             80.02Ki ± ∞ ¹
Select/Ordered/n=10000/k=100/sorted-16            80.00Ki ± ∞ ¹
Select/Func/n=10000/k=100/sorted-16               80.00Ki ± ∞ ¹
Select/Sort/n=10000/k=100/sorted-16               80.00Ki ± ∞ ¹
Select/Select/n=10000/k=100/reversed-16           80.02Ki ± ∞ ¹
Select/Ordered/n=10000/k=100/reversed-16          80.00Ki ± ∞ ¹
Select/Func/n=10000/k=100/reversed-16             80.00Ki ± ∞ ¹
Select/Sort/n=10000/k=100/reversed-16             80.00Ki ± ∞ ¹
Select/Select/n=10000/k=100/equal-16              80.02Ki ± ∞ ¹
Select/Ordered/n=10000/k=100/equal-16             80.00Ki ± ∞ ¹
Select/Func/n=10000/k=100/equal-16                80.00Ki ± ∞ ¹
Select/Sort/n=10000/k=100/equal-16                80.00Ki ± ∞ ¹
Select/Select/n=10000/k=100/mostly_equal-16       80.02Ki ± ∞ ¹
Select/Ordered/n=10000/k=100/mostly_equal-16      80.00Ki ± ∞ ¹
Select/Func/n=10000/k=100/mostly_equal-16         80.00Ki ± ∞ ¹
Select/Sort/n=10000/k=100/mostly_equal-16         80.00Ki ± ∞ ¹
Select/Select/n=10000/k=500/random-16             80.02Ki ± ∞ ¹
Select/Ordered/n=10000/k=500/random-16            80.00Ki ± ∞ ¹
Select/Func/n=10000/k=500/random-16               80.00Ki ± ∞ ¹
Select/Sort/n=10000/k=500/random-16               80.00Ki ± ∞ ¹
Select/Select/n=10000/k=500/sorted-16             80.02Ki ± ∞ ¹
Select/Ordered/n=10000/k=500/sorted-16            80.00Ki ± ∞ ¹
Select/Func/n=10000/k=500/sorted-16               80.00Ki ± ∞ ¹
Select/Sort/n=10000/k=500/sorted-16               80.00Ki ± ∞ ¹
Select/Select/n=10000/k=500/reversed-16           80.02Ki ± ∞ ¹
Select/Ordered/n=10000/k=500/reversed-16          80.00Ki ± ∞ ¹
Select/Func/n=10000/k=500/reversed-16             80.00Ki ± ∞ ¹
Select/Sort/n=10000/k=500/reversed-16             80.00Ki ± ∞ ¹
Select/Select/n=10000/k=500/equal-16              80.02Ki ± ∞ ¹
Select/Ordered/n=10000/k=500/equal-16             80.00Ki ± ∞ ¹
Select/Func/n=10000/k=500/equal-16                80.00Ki ± ∞ ¹
Select/Sort/n=10000/k=500/equal-16                80.00Ki ± ∞ ¹
Select/Select/n=10000/k=500/mostly_equal-16       80.02Ki ± ∞ ¹
Select/Ordered/n=10000/k=500/mostly_equal-16      80.00Ki ± ∞ ¹
Select/Func/n=10000/k=500/mostly_equal-16         80.00Ki ± ∞ ¹
Select/Sort/n=10000/k=500/mostly_equal-16         80.00Ki ± ∞ ¹
Select/Select/n=10000/k=1000/random-16            80.02Ki ± ∞ ¹
Select/Ordered/n=10000/k=1000/random-16           80.00Ki ± ∞ ¹
Select/Func/n=10000/k=1000/random-16              80.00Ki ± ∞ ¹
Select/Sort/n=10000/k=1000/random-16              80.00Ki ± ∞ ¹
Select/Select/n=10000/k=1000/sorted-16            80.02Ki ± ∞ ¹
Select/Ordered/n=10000/k=1000/sorted-16           80.00Ki ± ∞ ¹
Select/Func/n=10000/k=1000/sorted-16              80.00Ki ± ∞ ¹
Select/Sort/n=10000/k=1000/sorted-16              80.00Ki ± ∞ ¹
Select/Select/n=10000/k=1000/reversed-16          80.02Ki ± ∞ ¹
Select/Ordered/n=10000/k=1000/reversed-16         80.00Ki ± ∞ ¹
Select/Func/n=10000/k=1000/reversed-16            80.00Ki ± ∞ ¹
Select/Sort/n=10000/k=1000/reversed-16            80.00Ki ± ∞ ¹
Select/Select/n=10000/k=1000/equal-16             80.02Ki ± ∞ ¹
Select/Ordered/n=10000/k=1000/equal-16            80.00Ki ± ∞ ¹
Select/Func/n=10000/k=1000/equal-16               80.00Ki ± ∞ ¹
Select/Sort/n=10000/k=1000/equal-16               80.00Ki ± ∞ ¹
Select/Select/n=10000/k=1000/mostly_equal-16      80.02Ki ± ∞ ¹
Select/Ordered/n=10000/k=1000/mostly_equal-16     80.00Ki ± ∞ ¹
Select/Func/n=10000/k=1000/mostly_equal-16        80.00Ki ± ∞ ¹
Select/Sort/n=10000/k=1000/mostly_equal-16        80.00Ki ± ∞ ¹
Select/Select/n=1000/k=10/random-16               8.023Ki ± ∞ ¹
Select/Ordered/n=1000/k=10/random-16              8.000Ki ± ∞ ¹
Select/Func/n=1000/k=10/random-16                 8.000Ki ± ∞ ¹
Select/Sort/n=1000/k=10/random-16                 8.000Ki ± ∞ ¹
Select/Select/n=1000/k=10/sorted-16               8.023Ki ± ∞ ¹
Select/Ordered/n=1000/k=10/sorted-16              8.000Ki ± ∞ ¹
Select/Func/n=1000/k=10/sorted-16                 8.000Ki ± ∞ ¹
Select/Sort/n=1000/k=10/sorted-16                 8.000Ki ± ∞ ¹
Select/Select/n=1000/k=10/reversed-16             8.023Ki ± ∞ ¹
Select/Ordered/n=1000/k=10/reversed-16            8.000Ki ± ∞ ¹
Select/Func/n=1000/k=10/reversed-16               8.000Ki ± ∞ ¹
Select/Sort/n=1000/k=10/reversed-16               8.000Ki ± ∞ ¹
Select/Select/n=1000/k=10/equal-16                8.023Ki ± ∞ ¹
Select/Ordered/n=1000/k=10/equal-16               8.000Ki ± ∞ ¹
Select/Func/n=1000/k=10/equal-16                  8.000Ki ± ∞ ¹
Select/Sort/n=1000/k=10/equal-16                  8.000Ki ± ∞ ¹
Select/Select/n=1000/k=10/mostly_equal-16         8.023Ki ± ∞ ¹
Select/Ordered/n=1000/k=10/mostly_equal-16        8.000Ki ± ∞ ¹
Select/Func/n=1000/k=10/mostly_equal-16           8.000Ki ± ∞ ¹
Select/Sort/n=1000/k=10/mostly_equal-16           8.000Ki ± ∞ ¹
Select/Select/n=1000/k=50/random-16               8.023Ki ± ∞ ¹
Select/Ordered/n=1000/k=50/random-16              8.000Ki ± ∞ ¹
Select/Func/n=1000/k=50/random-16                 8.000Ki ± ∞ ¹
Select/Sort/n=1000/k=50/random-16                 8.000Ki ± ∞ ¹
Select/Select/n=1000/k=50/sorted-16               8.023Ki ± ∞ ¹
Select/Ordered/n=1000/k=50/sorted-16              8.000Ki ± ∞ ¹
Select/Func/n=1000/k=50/sorted-16                 8.000Ki ± ∞ ¹
Select/Sort/n=1000/k=50/sorted-16                 8.000Ki ± ∞ ¹
Select/Select/n=1000/k=50/reversed-16             8.023Ki ± ∞ ¹
Select/Ordered/n=1000/k=50/reversed-16            8.000Ki ± ∞ ¹
Select/Func/n=1000/k=50/reversed-16               8.000Ki ± ∞ ¹
Select/Sort/n=1000/k=50/reversed-16               8.000Ki ± ∞ ¹
Select/Select/n=1000/k=50/equal-16                8.023Ki ± ∞ ¹
Select/Ordered/n=1000/k=50/equal-16               8.000Ki ± ∞ ¹
Select/Func/n=1000/k=50/equal-16                  8.000Ki ± ∞ ¹
Select/Sort/n=1000/k=50/equal-16                  8.000Ki ± ∞ ¹
Select/Select/n=1000/k=50/mostly_equal-16         8.023Ki ± ∞ ¹
Select/Ordered/n=1000/k=50/mostly_equal-16        8.000Ki ± ∞ ¹
Select/Func/n=1000/k=50/mostly_equal-16           8.000Ki ± ∞ ¹
Select/Sort/n=1000/k=50/mostly_equal-16           8.000Ki ± ∞ ¹
Select/Select/n=1000/k=100/random-16              8.023Ki ± ∞ ¹
Select/Ordered/n=1000/k=100/random-16             8.000Ki ± ∞ ¹
Select/Func/n=1000/k=100/random-16                8.000Ki ± ∞ ¹
Select/Sort/n=1000/k=100/random-16                8.000Ki ± ∞ ¹
Select/Select/n=1000/k=100/sorted-16              8.023Ki ± ∞ ¹
Select/Ordered/n=1000/k=100/sorted-16             8.000Ki ± ∞ ¹
Select/Func/n=1000/k=100/sorted-16                8.000Ki ± ∞ ¹
Select/Sort/n=1000/k=100/sorted-16                8.000Ki ± ∞ ¹
Select/Select/n=1000/k=100/reversed-16            8.023Ki ± ∞ ¹
Select/Ordered/n=1000/k=100/reversed-16           8.000Ki ± ∞ ¹
Select/Func/n=1000/k=100/reversed-16              8.000Ki ± ∞ ¹
Select/Sort/n=1000/k=100/reversed-16              8.000Ki ± ∞ ¹
Select/Select/n=1000/k=100/equal-16               8.023Ki ± ∞ ¹
Select/Ordered/n=1000/k=100/equal-16              8.000Ki ± ∞ ¹
Select/Func/n=1000/k=100/equal-16                 8.000Ki ± ∞ ¹
Select/Sort/n=1000/k=100/equal-16                 8.000Ki ± ∞ ¹
Select/Select/n=1000/k=100/mostly_equal-16        8.023Ki ± ∞ ¹
Select/Ordered/n=1000/k=100/mostly_equal-16       8.000Ki ± ∞ ¹
Select/Func/n=1000/k=100/mostly_equal-16          8.000Ki ± ∞ ¹
Select/Sort/n=1000/k=100/mostly_equal-16          8.000Ki ± ∞ ¹
geomean                                           79.48Ki
¹ need >= 6 samples for confidence interval at level 0.95

                                                │    bench    │
                                                │  allocs/op  │
Select/Select/n=100000/k=1000/random-16           2.000 ± ∞ ¹
Select/Ordered/n=100000/k=1000/random-16          1.000 ± ∞ ¹
Select/Func/n=100000/k=1000/random-16             1.000 ± ∞ ¹
Select/Sort/n=100000/k=1000/random-16             1.000 ± ∞ ¹
Select/Select/n=100000/k=1000/sorted-16           2.000 ± ∞ ¹
Select/Ordered/n=100000/k=1000/sorted-16          1.000 ± ∞ ¹
Select/Func/n=100000/k=1000/sorted-16             1.000 ± ∞ ¹
Select/Sort/n=100000/k=1000/sorted-16             1.000 ± ∞ ¹
Select/Select/n=100000/k=1000/reversed-16         2.000 ± ∞ ¹
Select/Ordered/n=100000/k=1000/reversed-16        1.000 ± ∞ ¹
Select/Func/n=100000/k=1000/reversed-16           1.000 ± ∞ ¹
Select/Sort/n=100000/k=1000/reversed-16           1.000 ± ∞ ¹
Select/Select/n=100000/k=1000/equal-16            2.000 ± ∞ ¹
Select/Ordered/n=100000/k=1000/equal-16           1.000 ± ∞ ¹
Select/Func/n=100000/k=1000/equal-16              1.000 ± ∞ ¹
Select/Sort/n=100000/k=1000/equal-16              1.000 ± ∞ ¹
Select/Select/n=100000/k=1000/mostly_equal-16     2.000 ± ∞ ¹
Select/Ordered/n=100000/k=1000/mostly_equal-16    1.000 ± ∞ ¹
Select/Func/n=100000/k=1000/mostly_equal-16       1.000 ± ∞ ¹
Select/Sort/n=100000/k=1000/mostly_equal-16       1.000 ± ∞ ¹
Select/Select/n=100000/k=5000/random-16           2.000 ± ∞ ¹
Select/Ordered/n=100000/k=5000/random-16          1.000 ± ∞ ¹
Select/Func/n=100000/k=5000/random-16             1.000 ± ∞ ¹
Select/Sort/n=100000/k=5000/random-16             1.000 ± ∞ ¹
Select/Select/n=100000/k=5000/sorted-16           2.000 ± ∞ ¹
Select/Ordered/n=100000/k=5000/sorted-16          1.000 ± ∞ ¹
Select/Func/n=100000/k=5000/sorted-16             1.000 ± ∞ ¹
Select/Sort/n=100000/k=5000/sorted-16             1.000 ± ∞ ¹
Select/Select/n=100000/k=5000/reversed-16         2.000 ± ∞ ¹
Select/Ordered/n=100000/k=5000/reversed-16        1.000 ± ∞ ¹
Select/Func/n=100000/k=5000/reversed-16           1.000 ± ∞ ¹
Select/Sort/n=100000/k=5000/reversed-16           1.000 ± ∞ ¹
Select/Select/n=100000/k=5000/equal-16            2.000 ± ∞ ¹
Select/Ordered/n=100000/k=5000/equal-16           1.000 ± ∞ ¹
Select/Func/n=100000/k=5000/equal-16              1.000 ± ∞ ¹
Select/Sort/n=100000/k=5000/equal-16              1.000 ± ∞ ¹
Select/Select/n=100000/k=5000/mostly_equal-16     2.000 ± ∞ ¹
Select/Ordered/n=100000/k=5000/mostly_equal-16    1.000 ± ∞ ¹
Select/Func/n=100000/k=5000/mostly_equal-16       1.000 ± ∞ ¹
Select/Sort/n=100000/k=5000/mostly_equal-16       1.000 ± ∞ ¹
Select/Select/n=100000/k=10000/random-16          2.000 ± ∞ ¹
Select/Ordered/n=100000/k=10000/random-16         1.000 ± ∞ ¹
Select/Func/n=100000/k=10000/random-16            1.000 ± ∞ ¹
Select/Sort/n=100000/k=10000/random-16            1.000 ± ∞ ¹
Select/Select/n=100000/k=10000/sorted-16          2.000 ± ∞ ¹
Select/Ordered/n=100000/k=10000/sorted-16         1.000 ± ∞ ¹
Select/Func/n=100000/k=10000/sorted-16            1.000 ± ∞ ¹
Select/Sort/n=100000/k=10000/sorted-16            1.000 ± ∞ ¹
Select/Select/n=100000/k=10000/reversed-16        2.000 ± ∞ ¹
Select/Ordered/n=100000/k=10000/reversed-16       1.000 ± ∞ ¹
Select/Func/n=100000/k=10000/reversed-16          1.000 ± ∞ ¹
Select/Sort/n=100000/k=10000/reversed-16          1.000 ± ∞ ¹
Select/Select/n=100000/k=10000/equal-16           2.000 ± ∞ ¹
Select/Ordered/n=100000/k=10000/equal-16          1.000 ± ∞ ¹
Select/Func/n=100000/k=10000/equal-16             1.000 ± ∞ ¹
Select/Sort/n=100000/k=10000/equal-16             1.000 ± ∞ ¹
Select/Select/n=100000/k=10000/mostly_equal-16    2.000 ± ∞ ¹
Select/Ordered/n=100000/k=10000/mostly_equal-16   1.000 ± ∞ ¹
Select/Func/n=100000/k=10000/mostly_equal-16      1.000 ± ∞ ¹
Select/Sort/n=100000/k=10000/mostly_equal-16      1.000 ± ∞ ¹
Select/Select/n=10000/k=100/random-16             2.000 ± ∞ ¹
Select/Ordered/n=10000/k=100/random-16            1.000 ± ∞ ¹
Select/Func/n=10000/k=100/random-16               1.000 ± ∞ ¹
Select/Sort/n=10000/k=100/random-16               1.000 ± ∞ ¹
Select/Select/n=10000/k=100/sorted-16             2.000 ± ∞ ¹
Select/Ordered/n=10000/k=100/sorted-16            1.000 ± ∞ ¹
Select/Func/n=10000/k=100/sorted-16               1.000 ± ∞ ¹
Select/Sort/n=10000/k=100/sorted-16               1.000 ± ∞ ¹
Select/Select/n=10000/k=100/reversed-16           2.000 ± ∞ ¹
Select/Ordered/n=10000/k=100/reversed-16          1.000 ± ∞ ¹
Select/Func/n=10000/k=100/reversed-16             1.000 ± ∞ ¹
Select/Sort/n=10000/k=100/reversed-16             1.000 ± ∞ ¹
Select/Select/n=10000/k=100/equal-16              2.000 ± ∞ ¹
Select/Ordered/n=10000/k=100/equal-16             1.000 ± ∞ ¹
Select/Func/n=10000/k=100/equal-16                1.000 ± ∞ ¹
Select/Sort/n=10000/k=100/equal-16                1.000 ± ∞ ¹
Select/Select/n=10000/k=100/mostly_equal-16       2.000 ± ∞ ¹
Select/Ordered/n=10000/k=100/mostly_equal-16      1.000 ± ∞ ¹
Select/Func/n=10000/k=100/mostly_equal-16         1.000 ± ∞ ¹
Select/Sort/n=10000/k=100/mostly_equal-16         1.000 ± ∞ ¹
Select/Select/n=10000/k=500/random-16             2.000 ± ∞ ¹
Select/Ordered/n=10000/k=500/random-16            1.000 ± ∞ ¹
Select/Func/n=10000/k=500/random-16               1.000 ± ∞ ¹
Select/Sort/n=10000/k=500/random-16               1.000 ± ∞ ¹
Select/Select/n=10000/k=500/sorted-16             2.000 ± ∞ ¹
Select/Ordered/n=10000/k=500/sorted-16            1.000 ± ∞ ¹
Select/Func/n=10000/k=500/sorted-16               1.000 ± ∞ ¹
Select/Sort/n=10000/k=500/sorted-16               1.000 ± ∞ ¹
Select/Select/n=10000/k=500/reversed-16           2.000 ± ∞ ¹
Select/Ordered/n=10000/k=500/reversed-16          1.000 ± ∞ ¹
Select/Func/n=10000/k=500/reversed-16             1.000 ± ∞ ¹
Select/Sort/n=10000/k=500/reversed-16             1.000 ± ∞ ¹
Select/Select/n=10000/k=500/equal-16              2.000 ± ∞ ¹
Select/Ordered/n=10000/k=500/equal-16             1.000 ± ∞ ¹
Select/Func/n=10000/k=500/equal-16                1.000 ± ∞ ¹
Select/Sort/n=10000/k=500/equal-16                1.000 ± ∞ ¹
Select/Select/n=10000/k=500/mostly_equal-16       2.000 ± ∞ ¹
Select/Ordered/n=10000/k=500/mostly_equal-16      1.000 ± ∞ ¹
Select/Func/n=10000/k=500/mostly_equal-16         1.000 ± ∞ ¹
Select/Sort/n=10000/k=500/mostly_equal-16         1.000 ± ∞ ¹
Select/Select/n=10000/k=1000/random-16            2.000 ± ∞ ¹
Select/Ordered/n=10000/k=1000/random-16           1.000 ± ∞ ¹
Select/Func/n=10000/k=1000/random-16              1.000 ± ∞ ¹
Select/Sort/n=10000/k=1000/random-16              1.000 ± ∞ ¹
Select/Select/n=10000/k=1000/sorted-16            2.000 ± ∞ ¹
Select/Ordered/n=10000/k=1000/sorted-16           1.000 ± ∞ ¹
Select/Func/n=10000/k=1000/sorted-16              1.000 ± ∞ ¹
Select/Sort/n=10000/k=1000/sorted-16              1.000 ± ∞ ¹
Select/Select/n=10000/k=1000/reversed-16          2.000 ± ∞ ¹
Select/Ordered/n=10000/k=1000/reversed-16         1.000 ± ∞ ¹
Select/Func/n=10000/k=1000/reversed-16            1.000 ± ∞ ¹
Select/Sort/n=10000/k=1000/reversed-16            1.000 ± ∞ ¹
Select/Select/n=10000/k=1000/equal-16             2.000 ± ∞ ¹
Select/Ordered/n=10000/k=1000/equal-16            1.000 ± ∞ ¹
Select/Func/n=10000/k=1000/equal-16               1.000 ± ∞ ¹
Select/Sort/n=10000/k=1000/equal-16               1.000 ± ∞ ¹
Select/Select/n=10000/k=1000/mostly_equal-16      2.000 ± ∞ ¹
Select/Ordered/n=10000/k=1000/mostly_equal-16     1.000 ± ∞ ¹
Select/Func/n=10000/k=1000/mostly_equal-16        1.000 ± ∞ ¹
Select/Sort/n=10000/k=1000/mostly_equal-16        1.000 ± ∞ ¹
Select/Select/n=1000/k=10/random-16               2.000 ± ∞ ¹
Select/Ordered/n=1000/k=10/random-16              1.000 ± ∞ ¹
Select/Func/n=1000/k=10/random-16                 1.000 ± ∞ ¹
Select/Sort/n=1000/k=10/random-16                 1.000 ± ∞ ¹
Select/Select/n=1000/k=10/sorted-16               2.000 ± ∞ ¹
Select/Ordered/n=1000/k=10/sorted-16              1.000 ± ∞ ¹
Select/Func/n=1000/k=10/sorted-16                 1.000 ± ∞ ¹
Select/Sort/n=1000/k=10/sorted-16                 1.000 ± ∞ ¹
Select/Select/n=1000/k=10/reversed-16             2.000 ± ∞ ¹
Select/Ordered/n=1000/k=10/reversed-16            1.000 ± ∞ ¹
Select/Func/n=1000/k=10/reversed-16               1.000 ± ∞ ¹
Select/Sort/n=1000/k=10/reversed-16               1.000 ± ∞ ¹
Select/Select/n=1000/k=10/equal-16                2.000 ± ∞ ¹
Select/Ordered/n=1000/k=10/equal-16               1.000 ± ∞ ¹
Select/Func/n=1000/k=10/equal-16                  1.000 ± ∞ ¹
Select/Sort/n=1000/k=10/equal-16                  1.000 ± ∞ ¹
Select/Select/n=1000/k=10/mostly_equal-16         2.000 ± ∞ ¹
Select/Ordered/n=1000/k=10/mostly_equal-16        1.000 ± ∞ ¹
Select/Func/n=1000/k=10/mostly_equal-16           1.000 ± ∞ ¹
Select/Sort/n=1000/k=10/mostly_equal-16           1.000 ± ∞ ¹
Select/Select/n=1000/k=50/random-16               2.000 ± ∞ ¹
Select/Ordered/n=1000/k=50/random-16              1.000 ± ∞ ¹
Select/Func/n=1000/k=50/random-16                 1.000 ± ∞ ¹
Select/Sort/n=1000/k=50/random-16                 1.000 ± ∞ ¹
Select/Select/n=1000/k=50/sorted-16               2.000 ± ∞ ¹
Select/Ordered/n=1000/k=50/sorted-16              1.000 ± ∞ ¹
Select/Func/n=1000/k=50/sorted-16                 1.000 ± ∞ ¹
Select/Sort/n=1000/k=50/sorted-16                 1.000 ± ∞ ¹
Select/Select/n=1000/k=50/reversed-16             2.000 ± ∞ ¹
Select/Ordered/n=1000/k=50/reversed-16            1.000 ± ∞ ¹
Select/Func/n=1000/k=50/reversed-16               1.000 ± ∞ ¹
Select/Sort/n=1000/k=50/reversed-16               1.000 ± ∞ ¹
Select/Select/n=1000/k=50/equal-16                2.000 ± ∞ ¹
Select/Ordered/n=1000/k=50/equal-16               1.000 ± ∞ ¹
Select/Func/n=1000/k=50/equal-16                  1.000 ± ∞ ¹
Select/Sort/n=1000/k=50/equal-16                  1.000 ± ∞ ¹
Select/Select/n=1000/k=50/mostly_equal-16         2.000 ± ∞ ¹
Select/Ordered/n=1000/k=50/mostly_equal-16        1.000 ± ∞ ¹
Select/Func/n=1000/k=50/mostly_equal-16           1.000 ± ∞ ¹
Select/Sort/n=1000/k=50/mostly_equal-16           1.000 ± ∞ ¹
Select/Select/n=1000/k=100/random-16              2.000 ± ∞ ¹
Select/Ordered/n=1000/k=100/random-16             1.000 ± ∞ ¹
Select/Func/n=1000/k=100/random-16                1.000 ± ∞ ¹
Select/Sort/n=1000/k=100/random-16                1.000 ± ∞ ¹
Select/Select/n=1000/k=100/sorted-16              2.000 ± ∞ ¹
Select/Ordered/n=1000/k=100/sorted-16             1.000 ± ∞ ¹
Select/Func/n=1000/k=100/sorted-16                1.000 ± ∞ ¹
Select/Sort/n=1000/k=100/sorted-16                1.000 ± ∞ ¹
Select/Select/n=1000/k=100/reversed-16            2.000 ± ∞ ¹
Select/Ordered/n=1000/k=100/reversed-16           1.000 ± ∞ ¹
Select/Func/n=1000/k=100/reversed-16              1.000 ± ∞ ¹
Select/Sort/n=1000/k=100/reversed-16              1.000 ± ∞ ¹
Select/Select/n=1000/k=100/equal-16               2.000 ± ∞ ¹
Select/Ordered/n=1000/k=100/equal-16              1.000 ± ∞ ¹
Select/Func/n=1000/k=100/equal-16                 1.000 ± ∞ ¹
Select/Sort/n=1000/k=100/equal-16                 1.000 ± ∞ ¹
Select/Select/n=1000/k=100/mostly_equal-16        2.000 ± ∞ ¹
Select/Ordered/n=1000/k=100/mostly_equal-16       1.000 ± ∞ ¹
Select/Func/n=1000/k=100/mostly_equal-16          1.000 ± ∞ ¹
Select/Sort/n=1000/k=100/mostly_equal-16          1.000 ± ∞ ¹
geomean                                           1.189
¹ need >= 6 samples for confidence interval at level 0.95
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built upon Go's internal `pdqsort` algorithm.
