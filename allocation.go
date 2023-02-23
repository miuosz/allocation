package allocation

import (
	"fmt"
	"time"
)

const (
	standard = 1024

	Byte     uint64 = 1024
	KiloByte uint64 = Byte * standard
	MegaByte uint64 = KiloByte * standard
	GigaByte uint64 = MegaByte * standard
	TeraByte uint64 = GigaByte * standard
	PetaByte uint64 = TeraByte * standard
)

type Settings struct {
	size     uint64         // 1MB by default
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

// WithSize specifies how much memory should be allocated.
func WithSize(size uint64) func(s *Settings) {
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
	allocSize := s.size / uint64(s.count)
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
