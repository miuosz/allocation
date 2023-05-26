package allocation

import (
	"time"
)

// Size represents a memory size in bytes.
// It is a custom type based on the uint64 data type.
type Size uint64

// Predefined Size constants for common memory sizes.
const (
	standard = 1024

	Byte     Size = 1
	KiloByte Size = Byte * standard
	MegaByte Size = KiloByte * standard
	GigaByte Size = MegaByte * standard
	TeraByte Size = GigaByte * standard
	PetaByte Size = TeraByte * standard
)

type settings struct {
	size     Size
	amount   int
	duration *time.Duration
	physical bool
}

// New creates multiple allocations with specified size and options.
// It creates 'amount' number of allocations, each of size 'size'.
//
// Parameters:
//   - amount: The number of allocations to create.
//   - size: The size of each allocation. Use the 'Size' type to specify the size in bytes, kilobytes, megabytes, etc.
//   - physical: Specifies whether to use physical memory for the allocations.
//   - duration: Specifies the duration for which the allocations should run. Use nil for no duration.
//
// Returns:
//   - A 2D slice of bytes representing the allocated memory.
//
// Example:
//
//	allocation.New(10, 10*allocation.MegaByte, false, nil)
//	// Creates 10 allocations, each of 10 megabytes in size.
//
//	allocation.New(5, 1*allocation.KiloByte, true, nil)
//	// Creates 5 allocations, each of 1 kilobyte in size, using physical memory.
func New(amount int, size Size, physical bool, duration *time.Duration) [][]byte {
	s := settings{
		size:     size,
		amount:   amount,
		physical: physical,
		duration: duration,
	}

	return allocate(s)
}

func allocate(s settings) [][]byte {
	payload := make([][]byte, s.amount)

	for i := 0; i < s.amount; i++ {
		// TODO: add some size limit.
		alloc := make([]byte, s.size)
		payload[i] = alloc
	}

	if s.physical {
		useMem(payload)
	}

	if s.duration != nil {
		wait(*s.duration)
	}

	return payload
}

func useMem(p [][]byte) {
	for i := range p {
		allocSize := len(p[i])
		// Alter one byte every 4KB
		for j := 0; j < allocSize; j += 4 << 10 {
			p[i][j] = 1
		}
	}
}

func wait(d time.Duration) {
	time.Sleep(d)
}
