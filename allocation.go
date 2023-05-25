package allocation

import (
	"fmt"
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
}

type Allocation struct {
	payload [][]byte
}

func (a Allocation) String() string {
	return fmt.Sprintf("allocated: %d elements", len(a.payload))
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
	allocSize := uint64(s.size) / uint64(s.count)
	payload := make([][]byte, 0, s.count)

	for i := 0; i < s.count; i++ {
		alloc := make([]byte, 0, allocSize)
		payload = append(payload, alloc)
	}

	if s.duration != nil {
		wait(*s.duration)
	}

	return Allocation{
		payload: payload,
	}
}

func wait(d time.Duration) {
	time.Sleep(d)
}
