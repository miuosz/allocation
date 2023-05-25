package main

import (
	"fmt"
	"runtime"

	"github.com/miuosz/allocation"
)

func main() {
	allocation.New(
		allocation.WithCount(1),
		allocation.WithSize(10*allocation.MegaByte),
	)

	runtime.GC()
	printMemStats()
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Alloc = %v MiB\n", byteToMB(m.Alloc))
	fmt.Printf("TotalAlloc = %v MiB\n", byteToMB(m.TotalAlloc))
	fmt.Printf("Sys = %v MiB\n", byteToMB(m.Sys))
}

func byteToMB(b uint64) uint64 {
	return b / 1024 / 1024
}
