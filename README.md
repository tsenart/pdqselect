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
BenchmarkSelect/Select/n=100000/k=1000/random-16                  3258	   1108941 ns/op	  802844 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=100000/k=1000/random-16  	              2344	    438276 ns/op	  802822 B/op	       1 allocs/op
BenchmarkSelect/Func/n=100000/k=1000/random-16     	              1526	    740673 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=100000/k=1000/random-16           	               249	   4601888 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Select/n=100000/k=1000/sorted-16         	              4862	    243432 ns/op	  802844 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=100000/k=1000/sorted-16  	             10000	    100042 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Func/n=100000/k=1000/sorted-16     	              5061	    232981 ns/op	  802818 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=100000/k=1000/sorted-16           	             12018	    100507 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Select/n=100000/k=1000/reversed-16       	              3636	    322333 ns/op	  802845 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=100000/k=1000/reversed-16         	    9745	    120925 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Func/n=100000/k=1000/reversed-16            	    4698	    248515 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=100000/k=1000/reversed-16                  	    9674	    118380 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Select/n=100000/k=1000/equal-16                   	    4837	    244395 ns/op	  802844 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=100000/k=1000/equal-16            	   12447	     99205 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Func/n=100000/k=1000/equal-16               	    5127	    228209 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=100000/k=1000/equal-16                     	   12330	     99987 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Select/n=100000/k=1000/mostly_equal-16            	    2350	    525269 ns/op	  802844 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=100000/k=1000/mostly_equal-16     	    6949	    156052 ns/op	  802818 B/op	       1 allocs/op
BenchmarkSelect/Func/n=100000/k=1000/mostly_equal-16        	    2215	    516669 ns/op	  802818 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=100000/k=1000/mostly_equal-16              	    3674	    323622 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Select/n=100000/k=5000/random-16                  	    1431	    767205 ns/op	  802844 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=100000/k=5000/random-16           	    2710	    375032 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Func/n=100000/k=5000/random-16              	    1484	   1003712 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=100000/k=5000/random-16                    	     249	   4654386 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Select/n=100000/k=5000/sorted-16                  	    4842	    243750 ns/op	  802844 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=100000/k=5000/sorted-16           	   12247	    100854 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Func/n=100000/k=5000/sorted-16              	    4858	    232220 ns/op	  802818 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=100000/k=5000/sorted-16                    	   12028	     99973 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Select/n=100000/k=5000/reversed-16                	    3612	    322395 ns/op	  802844 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=100000/k=5000/reversed-16         	    9670	    118168 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Func/n=100000/k=5000/reversed-16            	    4551	    250005 ns/op	  802818 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=100000/k=5000/reversed-16                  	    9660	    121263 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Select/n=100000/k=5000/equal-16                   	    4765	    246508 ns/op	  802844 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=100000/k=5000/equal-16            	   10000	    100224 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Func/n=100000/k=5000/equal-16               	    4947	    229783 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=100000/k=5000/equal-16                     	   12201	     99952 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Select/n=100000/k=5000/mostly_equal-16            	    2272	    502406 ns/op	  802844 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=100000/k=5000/mostly_equal-16     	    6620	    152593 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Func/n=100000/k=5000/mostly_equal-16        	    2170	    537337 ns/op	  802818 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=100000/k=5000/mostly_equal-16              	    3524	    323867 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Select/n=100000/k=10000/random-16                 	    1416	    850070 ns/op	  802848 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=100000/k=10000/random-16          	    4720	    357472 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Func/n=100000/k=10000/random-16             	    2252	    857684 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=100000/k=10000/random-16                   	     253	   4698393 ns/op	  802818 B/op	       1 allocs/op
BenchmarkSelect/Select/n=100000/k=10000/sorted-16                 	    4857	    244634 ns/op	  802844 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=100000/k=10000/sorted-16          	   12151	    100510 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Func/n=100000/k=10000/sorted-16             	    5071	    230501 ns/op	  802818 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=100000/k=10000/sorted-16                   	   12350	    100514 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Select/n=100000/k=10000/reversed-16               	    3674	    321801 ns/op	  802843 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=100000/k=10000/reversed-16        	    9727	    121273 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Func/n=100000/k=10000/reversed-16           	    4692	    249779 ns/op	  802818 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=100000/k=10000/reversed-16                 	    9750	    118881 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Select/n=100000/k=10000/equal-16                  	    4744	    247068 ns/op	  802845 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=100000/k=10000/equal-16           	   10000	    100762 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Func/n=100000/k=10000/equal-16              	    4924	    232139 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=100000/k=10000/equal-16                    	   10000	    100225 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Select/n=100000/k=10000/mostly_equal-16           	    2272	    533003 ns/op	  802845 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=100000/k=10000/mostly_equal-16    	    6667	    159041 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Func/n=100000/k=10000/mostly_equal-16       	    2151	    558818 ns/op	  802818 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=100000/k=10000/mostly_equal-16             	    3249	    349239 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSelect/Select/n=10000/k=100/random-16                    	   28977	     40668 ns/op	   81944 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=10000/k=100/random-16             	   76352	     21823 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Func/n=10000/k=100/random-16                	   20977	     58405 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=10000/k=100/random-16                      	    6198	    173667 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Select/n=10000/k=100/sorted-16                    	   49035	     24397 ns/op	   81944 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=10000/k=100/sorted-16             	  120390	      9727 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Func/n=10000/k=100/sorted-16                	   48834	     23677 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=10000/k=100/sorted-16                      	  121506	      9919 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Select/n=10000/k=100/reversed-16                  	   37465	     32030 ns/op	   81944 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=10000/k=100/reversed-16           	  101815	     11925 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Func/n=10000/k=100/reversed-16              	   46480	     25820 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=10000/k=100/reversed-16                    	  100923	     11943 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Select/n=10000/k=100/equal-16                     	   48374	     24769 ns/op	   81944 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=10000/k=100/equal-16              	  119108	      9904 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Func/n=10000/k=100/equal-16                 	   50722	     23453 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=10000/k=100/equal-16                       	  120790	      9787 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Select/n=10000/k=100/mostly_equal-16              	   22870	     52412 ns/op	   81944 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=10000/k=100/mostly_equal-16       	   77167	     14787 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Func/n=10000/k=100/mostly_equal-16          	   23708	     51928 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=10000/k=100/mostly_equal-16                	   52896	     22942 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Select/n=10000/k=500/random-16                    	   23290	     60960 ns/op	   81944 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=10000/k=500/random-16             	   66494	     19980 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Func/n=10000/k=500/random-16                	   21919	     56477 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=10000/k=500/random-16                      	    6745	    172693 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Select/n=10000/k=500/sorted-16                    	   49100	     24744 ns/op	   81944 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=10000/k=500/sorted-16             	  121021	     10033 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Func/n=10000/k=500/sorted-16                	   50082	     23846 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=10000/k=500/sorted-16                      	  121392	      9902 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Select/n=10000/k=500/reversed-16                  	   37406	     32248 ns/op	   81944 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=10000/k=500/reversed-16           	   99088	     12000 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Func/n=10000/k=500/reversed-16              	   44866	     26477 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=10000/k=500/reversed-16                    	  101508	     11942 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Select/n=10000/k=500/equal-16                     	   48124	     24967 ns/op	   81944 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=10000/k=500/equal-16              	  120061	      9679 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Func/n=10000/k=500/equal-16                 	   49365	     24143 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=10000/k=500/equal-16                       	  123847	      9542 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Select/n=10000/k=500/mostly_equal-16              	   22813	     53273 ns/op	   81944 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=10000/k=500/mostly_equal-16       	   77716	     15036 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Func/n=10000/k=500/mostly_equal-16          	   22905	     52193 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=10000/k=500/mostly_equal-16                	   51708	     23190 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Select/n=10000/k=1000/random-16                   	   30990	     56201 ns/op	   81944 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=10000/k=1000/random-16            	   57715	     24283 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Func/n=10000/k=1000/random-16               	   25384	     60057 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=10000/k=1000/random-16                     	    6507	    163102 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Select/n=10000/k=1000/sorted-16                   	   48370	     24983 ns/op	   81944 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=10000/k=1000/sorted-16            	  121900	      9726 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Func/n=10000/k=1000/sorted-16               	   49191	     24161 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=10000/k=1000/sorted-16                     	  125038	      9896 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Select/n=10000/k=1000/reversed-16                 	   37262	     32302 ns/op	   81944 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=10000/k=1000/reversed-16          	  100383	     11868 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Func/n=10000/k=1000/reversed-16             	   45468	     26306 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=10000/k=1000/reversed-16                   	  101955	     11924 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Select/n=10000/k=1000/equal-16                    	   48316	     24879 ns/op	   81944 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=10000/k=1000/equal-16             	  123121	      9717 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Func/n=10000/k=1000/equal-16                	   49434	     24162 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=10000/k=1000/equal-16                      	  122238	      9663 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Select/n=10000/k=1000/mostly_equal-16             	   23989	     52254 ns/op	   81944 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=10000/k=1000/mostly_equal-16      	   81284	     15079 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Func/n=10000/k=1000/mostly_equal-16         	   22436	     55135 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=10000/k=1000/mostly_equal-16               	   52615	     22733 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSelect/Select/n=1000/k=10/random-16                      	  231312	      5975 ns/op	    8216 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=1000/k=10/random-16               	  543343	      1983 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Func/n=1000/k=10/random-16                  	  185162	      5552 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=1000/k=10/random-16                        	  114211	     10220 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Select/n=1000/k=10/sorted-16                      	  456824	      2640 ns/op	    8216 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=1000/k=10/sorted-16               	 1000000	      1118 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Func/n=1000/k=10/sorted-16                  	  470763	      2530 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=1000/k=10/sorted-16                        	 1000000	      1139 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Select/n=1000/k=10/reversed-16                    	  348313	      3403 ns/op	    8216 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=1000/k=10/reversed-16             	  867884	      1371 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Func/n=1000/k=10/reversed-16                	  426879	      2787 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=1000/k=10/reversed-16                      	  881938	      1365 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Select/n=1000/k=10/equal-16                       	  450236	      2670 ns/op	    8216 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=1000/k=10/equal-16                	 1000000	      1123 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Func/n=1000/k=10/equal-16                   	  464678	      2542 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=1000/k=10/equal-16                         	 1000000	      1120 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Select/n=1000/k=10/mostly_equal-16                	  223219	      9048 ns/op	    8216 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=1000/k=10/mostly_equal-16         	  589564	      1699 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Func/n=1000/k=10/mostly_equal-16            	  207612	      5554 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=1000/k=10/mostly_equal-16                  	  536198	      2479 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Select/n=1000/k=50/random-16                      	  248930	      5498 ns/op	    8216 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=1000/k=50/random-16               	  717555	      1843 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Func/n=1000/k=50/random-16                  	  287151	      5092 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=1000/k=50/random-16                        	  118708	     10420 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Select/n=1000/k=50/sorted-16                      	  447147	      2634 ns/op	    8216 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=1000/k=50/sorted-16               	 1000000	      1139 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Func/n=1000/k=50/sorted-16                  	  473912	      2512 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=1000/k=50/sorted-16                        	 1000000	      1108 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Select/n=1000/k=50/reversed-16                    	  352773	      3377 ns/op	    8216 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=1000/k=50/reversed-16             	  828056	      1366 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Func/n=1000/k=50/reversed-16                	  425404	      2781 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=1000/k=50/reversed-16                      	  899734	      1370 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Select/n=1000/k=50/equal-16                       	  452950	      2641 ns/op	    8216 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=1000/k=50/equal-16                	 1000000	      1123 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Func/n=1000/k=50/equal-16                   	  465646	      2524 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=1000/k=50/equal-16                         	 1000000	      1103 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Select/n=1000/k=50/mostly_equal-16                	  210338	      5315 ns/op	    8216 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=1000/k=50/mostly_equal-16         	  695546	      2088 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Func/n=1000/k=50/mostly_equal-16            	  216892	      5826 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=1000/k=50/mostly_equal-16                  	  536475	      2357 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Select/n=1000/k=100/random-16                     	  213662	      4894 ns/op	    8216 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=1000/k=100/random-16              	  543393	      2303 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Func/n=1000/k=100/random-16                 	  213747	      5595 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=1000/k=100/random-16                       	  116188	     10348 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Select/n=1000/k=100/sorted-16                     	  447568	      2638 ns/op	    8216 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=1000/k=100/sorted-16              	 1000000	      1128 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Func/n=1000/k=100/sorted-16                 	  464602	      2526 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=1000/k=100/sorted-16                       	 1000000	      1117 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Select/n=1000/k=100/reversed-16                   	  354590	      3381 ns/op	    8216 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=1000/k=100/reversed-16            	  847422	      1365 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Func/n=1000/k=100/reversed-16               	  430978	      2762 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=1000/k=100/reversed-16                     	  865695	      1363 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Select/n=1000/k=100/equal-16                      	  449875	      2652 ns/op	    8216 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=1000/k=100/equal-16               	 1000000	      1127 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Func/n=1000/k=100/equal-16                  	  467913	      2530 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=1000/k=100/equal-16                        	 1000000	      1107 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Select/n=1000/k=100/mostly_equal-16               	  243644	      6335 ns/op	    8216 B/op	       2 allocs/op
BenchmarkSelect/Ordered/n=1000/k=100/mostly_equal-16        	  665876	      1829 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Func/n=1000/k=100/mostly_equal-16           	  203697	      6150 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSelect/Sort/n=1000/k=100/mostly_equal-16                 	  526977	      2264 ns/op	    8192 B/op	       1 allocs/op
PASS
ok  	github.com/tsenart/pdqselect	262.149s
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built upon Go's internal `pdqsort` algorithm.
