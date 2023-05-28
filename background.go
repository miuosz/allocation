package allocation

import "time"

// NewBackground creates multiple allocations with specified size and options.
// It creates 'amount' number of allocations, each of size 'size'.
// In contrast to `New“, allocations made by `NewBackground“ are ready to be
// recycled immediately after they are allocated, even before the function returns.
//
// Parameters:
//   - amount (int): The number of allocations to be performed.
//   - size (Size): The size of each allocation.
//   - physical (bool): Flag indicating whether to use physical memory for the allocations.
//   - duration (*time.Duration): Optional duration to wait after performing the allocations.
func NewBackground(amount int, size Size, physical bool, duration *time.Duration) {
	for i := 0; i < amount; i++ {
		allocation := make([]byte, size)
		if physical {
			useMem(allocation)
		}
	}

	if duration != nil {
		wait(*duration)
	}
}

func useMem(payload []byte) {
	allocSize := len(payload)
	// Alter one byte every 4KB
	for j := 0; j < allocSize; j += 4 << 10 {
		payload[j] = 1
	}
}
