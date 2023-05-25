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
	duration *time.Duration // nil by default
	physical bool           // false by default
}

// Allocation is a structure that represents a collection of allocated payloads.
// It contains the allocated payloads in the form of a 2-dimensional byte slice.
//
// Fields:
//   - Payload: The allocated payloads stored as a 2-dimensional byte slice.
//     Each inner byte slice represents an individual allocation.
//
// Example:
//
//	alloc := allocation.New(10, 1*allocation.MegaByte)
//	// Creates an allocation with 10 payloads, each of 1 megabyte in size.
//
//	for _, payload := range alloc.Payload {
//	    // Process each individual payload
//	    fmt.Println(len(payload))
//	}
//	// Prints the size of each allocated payload.
type Allocation struct {
	// Payload represents allocated payloads stored as a 2-dimensional byte slice.
	Payload [][]byte
}

type Option func(s *settings)

// WithDuration specifies time to run allocations.
func WithDuration(duration time.Duration) func(s *settings) {
	return func(s *settings) {
		s.duration = &duration
	}
}

// WithPhysical makes allocation use physical memory.
func WithPhysical(physical bool) func(s *settings) {
	return func(s *settings) {
		s.physical = physical
	}
}

// New creates multiple allocations with specified size.
// It creates 'amount' number of allocations, each of size 'size'.
// Additional options can be provided to customize the allocation behavior.
//
// Parameters:
//   - amount: The number of allocations to create.
//   - size: The size of each allocation. Use the 'Size' type to specify the size in bytes, kilobytes, megabytes, etc.
//   - opts: Optional. Additional options to customize the allocation behavior. See the 'Option' type and the available options for more details.
//
// Returns:
//   - Allocation: The created allocation containing the specified number of allocations of the specified size.
//
// Example:
//
//	allocation.New(10, 10*allocation.MegaByte)
//	// Creates 10 allocations, each of 10 megabytes in size.
//
//	allocation.New(5, 1*allocation.KiloByte, allocation.WithPhysical(true))
//	// Creates 5 allocations, each of 1 kilobyte in size, using physical memory.
func New(amount int, size Size, opts ...Option) Allocation {
	s := &settings{
		size:     size,
		amount:   amount,
		duration: nil,
		physical: false,
	}

	for _, opt := range opts {
		opt(s)
	}

	return allocate(s)
}

func allocate(s *settings) Allocation {
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

	return Allocation{
		Payload: payload,
	}
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
