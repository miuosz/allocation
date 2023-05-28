package allocation

import (
	"time"
)

// Size represents a memory size in bytes.
// It is a custom type based on the uint64 data type.
type Size uint64

// Predefined Size constants for common memory sizes.
const (
	Byte Size = 1

	decimalStandard      = 1000
	KiloByte        Size = Byte * decimalStandard
	MegaByte        Size = KiloByte * decimalStandard
	GigaByte        Size = MegaByte * decimalStandard
	TeraByte        Size = GigaByte * decimalStandard
	PetaByte        Size = TeraByte * decimalStandard

	binaryStandard      = 1024
	KibiByte       Size = Byte * binaryStandard
	MebiByte       Size = KibiByte * binaryStandard
	GibiByte       Size = MebiByte * binaryStandard
	TebiByte       Size = GibiByte * binaryStandard
	PebiByte       Size = TebiByte * binaryStandard
)

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
	// TODO: add some size limit.
	payload := make([][]byte, amount)

	for i := 0; i < amount; i++ {
		alloc := make([]byte, size)
		payload[i] = alloc
	}

	if physical {
		for i := range payload {
			useMem(payload[i])
		}
	}

	if duration != nil {
		wait(*duration)
	}

	return payload
}

func wait(d time.Duration) {
	time.Sleep(d)
}
