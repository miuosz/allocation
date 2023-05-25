package allocation

import (
	"fmt"
	"math"
	"time"
)

type Size uint64

const (
	standard = 1024

	Byte     Size = 1
	KiloByte Size = Byte * standard
	MegaByte Size = KiloByte * standard
	GigaByte Size = MegaByte * standard
	TeraByte Size = GigaByte * standard
	PetaByte Size = TeraByte * standard
)

type Settings struct {
	size     Size           // 1MB by default
	count    int            // 1 by default
	duration *time.Duration // nil by default
	physical bool
}

type Allocation struct {
	Payload [][]byte
}

func (a Allocation) String() string {
	return fmt.Sprintf("allocated: %d elements", len(a.Payload))
}

type Option func(s *Settings)

// WithDuration specifies time to run allocations.
func WithDuration(duration time.Duration) func(s *Settings) {
	return func(s *Settings) {
		s.duration = &duration
	}
}

// WithSize is used to specify the total amount of memory that should be allocated.
// The specified size will be divided among the number of allocations specified by WithCount.
// For example, if you specify 10MB as the size and 10 as the count,
// it means that the total 10MB will be divided into ten 1MB objects.
func WithSize(size Size) func(s *Settings) {
	return func(s *Settings) {
		s.size = size
	}
}

// WithCount specifies how many allocations should be made.
func WithCount(count int) func(s *Settings) {
	return func(s *Settings) {
		s.count = count
	}
}

// WithPhysical makes allocation use physical memory.
func WithPhysical(physical bool) func(s *Settings) {
	return func(s *Settings) {
		s.physical = physical
	}
}

// New causes allocation with given parameters.
func New(opts ...Option) Allocation {
	s := &Settings{
		size:     MegaByte,
		count:    1,
		duration: nil,
	}

	for _, opt := range opts {
		opt(s)
	}

	return allocate(s)
}

func allocate(s *Settings) Allocation {
	tmpSize := uint64(s.size) / uint64(s.count)
	if tmpSize >= math.MaxInt64 {
		panic(fmt.Sprintf("too big alocSize(%v), increase count", tmpSize))
	}

	allocSize := int(tmpSize)

	payload := make([][]byte, s.count)

	for i := 0; i < s.count; i++ {
		alloc := make([]byte, allocSize)
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
