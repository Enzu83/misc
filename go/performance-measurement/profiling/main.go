package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func sumFunction(size int64) {
	slice := make([]int64, 0, size)
	for i := int64(0); i < size; i++ {
		slice = append(slice, i)
	}
	// Simulate some CPU work
	var sum int64 = 0
	for _, v := range slice {
		sum += v
	}
	fmt.Println("Sum:", sum)
}

func prodFunction(size int64) {
	slice := make([]int64, 0, size)
	for i := int64(1); i < size; i++ {
		slice = append(slice, i)
	}
	// Simulate some CPU work
	var prod int64 = 1
	for _, v := range slice {
		prod = (prod * v) % int64(1844674407370955161)
	}
	fmt.Println("Prod:", prod)
}

func main() {
	// --- CPU Profiling ---
	cpuFile, err := os.Create("cpu.prof")
	if err != nil {
		return
	}
	defer cpuFile.Close()

	if err := pprof.StartCPUProfile(cpuFile); err != nil {
		return
	}
	defer pprof.StopCPUProfile()

	// Run the function you want to profile
	start := time.Now()
	sumFunction(100_000_000)
	elapsed := time.Since(start)
	fmt.Printf("Sum Execution time: %s\n", elapsed)

	// Run the function you want to profile
	start = time.Now()
	prodFunction(100_000_000)
	elapsed = time.Since(start)
	fmt.Printf("Prod Execution time: %s\n", elapsed)

	run_c_leaking_function()

	// --- Memory Profiling ---
	memFile, err := os.Create("mem.prof")
	if err != nil {
		return
	}
	defer memFile.Close()

	runtime.GC() // force garbage collection to get up-to-date memory stats
	if err := pprof.WriteHeapProfile(memFile); err != nil {
		return
	}

	fmt.Println("Profiles written: cpu.prof and mem.prof")
}
