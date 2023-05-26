package main

import (
	"fmt"
	"runtime"

	"github.com/miuosz/allocation"
)

func main() {
	allocation.New(10, allocation.MegaByte, false, nil)

	runtime.GC()
	printMemStats()
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Alloc = %v MiB\n", byteToMiB(m.Alloc))
	fmt.Printf("TotalAlloc = %v MiB\n", byteToMiB(m.TotalAlloc))
	fmt.Printf("Sys = %v MiB\n", byteToMiB(m.Sys))
}

func byteToMiB(b uint64) uint64 {
	return b / 1024 / 1024
}
