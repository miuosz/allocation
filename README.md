# Allocation

Allocation is a simple library used for performing memory allocations.

## Quickstart

```go
package main

import (
	"fmt"
	"runtime"
	"runtime/debug"

	"github.com/miuosz/allocation"
)

func main() {
	allocation.New(1, 10*allocation.MegaByte)
}
```

## Allocations size precision
The library does not directly account for the size of `slice headers` and the size of the `Allocation structure`.
As a result, the reported memory usage may be slightly higher than the requested size, particularly noticeable when allocating small sizes.

```go
type Allocation struct {
	Payload [][]byte
}
```

<br />

For a single allocation of 1 byte, the memory usage will be as follows:

1. Allocation size (structure size) - `24 bytes`
2. Payload field size (slice header size) - `24 bytes`
3. Slice within the payload slice (slice header size) - `24 bytes`
4. Single byte element - `1 byte`

Total: `73 bytes`

<br />

The impact of this overhead grows linearly with the number of allocations.
For example:

1. Allocation of single 10MB object will take 10MB and 72 bytes
```go
allocation.New(1, 10*allocation.MegaByte)
```

2. Allocation of ten 10MB objects will take 10MB and 288 bytes
```go
allocation.New(10, 1*allocation.MegaByte)
```

So formula is: `48 + (amount * 24) + (amount * size)`
