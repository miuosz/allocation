Allocation is simple library used to do allocations.

# Quickstart
```go
package main

import (
	"fmt"
	"runtime"
	"runtime/debug"

	"github.com/miuosz/allocation"
)

func main() {
	allocation.New(
		allocation.WithCount(1),
		allocation.WithSize(10*allocation.MegaByte),
	)
}
```

# Allocations size precision
Generally speaking this library does not take into account size of slice headers and size of `Allocation`.
This could impact results while allocating small size.  
To sum up `SliceHeader` takes 24 bytes, `Allocation` takes 24 bytes.
```go
type Allocation struct {
	payload [][]byte
}
```
With such structure single allocation of `1 byte` will take up:
1. Allocation size (structure size) - `24 bytes`
2. payload field size (slice header size) - `24 bytes`
3. slice within payload slice (slice header size) - `24 bytes`
4. single element - `1 byte`

Total: 73 bytes

Impact of this overhead grows linearly to amount of allocations. 
For example 
1. Allocation of single 10MB object will take 10MB and 72 bytes
```go
allocation.New(
	allocation.WithCount(1),
	allocation.WithSize(10*allocation.MegaByte),
)
```
2. Allocation of ten 10MB objects will take 100MB and 288 bytes
```go
allocation.New(
	allocation.WithCount(10),
	allocation.WithSize(100*allocation.MegaByte),
)
```

So formula is: 48+(count*24)+size
