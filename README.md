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

## Allocations Size Precision
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

## Number of Allocations

Beside specified amount of allocations, there're always 1 additional allocation.
