package util

import (
	"fmt"
	"runtime"
)

func PrintMemory() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Printf("Allocated Memory: %.1f Mb\n", float64(mem.Alloc) / 1024 / 1024)
}
