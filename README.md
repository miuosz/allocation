# Allocation

Allocation is a simple library used for performing memory allocations.

## Quickstart

```go
package main

import (
	"github.com/miuosz/allocation"
)

func main() {
	allocation.New(1, 10*allocation.MegaByte, false, nil)
}
```

## New or NewBackground?
Allocations made using the `New` function are not recycled until they are no longer used. 
To illustrate this, let's consider an example where we perform 100 allocations of 1MB each:
```go
func main() {
	allocation.New(100, allocation.MegaByte, true, nil)
}
```
In this case, the garbage collector (GC) will run approximately 5 times (in Go 1.19) without reclaiming any memory. 
Consequently, the total memory usage will reach 100MB. 
As the GC does not reclaim any live heap memory, the heap size continuously grows, leading to an increased next heap target.

In contrast, allocations made by the `NewBackground` function are recycled `in real time`.
Even before the function returns, the allocations made by NewBackground become eligible for cleaning by the GC. 
Let's use a similar code example as before:
```go
func main() {
	allocation.NewBackground(100, allocation.MegaByte, true, nil)
}
```
In this scenario, the GC will run approximately 29 times (in Go 1.19), with each run claiming some of the allocated memory. As a result, the next heap target does not significantly increase or may slightly increase over time.

## Allocations Size Precision
### New
The library does not directly account for the size of `slice headers`.  
As a result, the reported memory usage may be slightly higher than the requested size, particularly noticeable when allocating small sizes.

For a single allocation of 1 byte, the memory usage will be as follows:

1. Slice header `24 bytes`
2. Payload itself: `1 byte`

Total: `25 bytes`

The impact of this overhead grows linearly with the number of allocations.
For example:

1. Allocation of single 10MB object will take 10MB and 24 bytes
```go
allocation.New(1, 10*allocation.MegaByte)
```

2. Allocation of ten 10MB objects will take 10MB and 240 bytes
```go
allocation.New(10, 1*allocation.MegaByte)
```

The approximate formula is: `(24 * amount) + (size * amount)`.

```go
allocation.New(1, Byte, false, nil)
```
```
Benchmark_New-8   	25325204	        46.78 ns/op	      25 B/op	       2 allocs/op
```

### NewBackground
Unlike the `New` function, `NewBackground` does not need to store slice headers on the heap, resulting in the exact specified allocation size.
```go
allocation.NewBackground(1, Byte, false, nil)
```
```
Benchmark_NewBackground-8   	81592506	        14.52 ns/op	       1 B/op	       1 allocs/op
```

## Number of Allocations
### New
Beside specified amount of allocations, there's always 1 additional allocation.

### NewBackground
Always allocates only specified amount of allocations.
