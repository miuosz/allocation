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

	fmt.Printf("Alloc = %v MB\n", byteToMB(m.Alloc))
	fmt.Printf("TotalAlloc = %v MB\n", byteToMB(m.TotalAlloc))
	fmt.Printf("Sys = %v MB\n", byteToMB(m.Sys))
}

func byteToMB(b uint64) uint64 {
	return b / 1000 / 1000
}
