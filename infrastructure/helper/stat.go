package helper

import (
	"fmt"
	"runtime"
	"time"

	"github.com/mackerelio/go-osstat/cpu"
)

type HwStats struct {
	NumCpu   int
	TotalMem uint64
	AvailMem uint64
	CpuUsage float64
}

func NewHwStats() HwStats {
	total, avail := getMemoryStat()
	cpuUsage := getCurrentCpuUsage()

	return HwStats{
		NumCpu:   runtime.NumCPU(),
		TotalMem: total,
		AvailMem: avail,
		CpuUsage: cpuUsage,
	}
}

func getMemoryStat() (total uint64, avail uint64) {
	// Get memory statistics
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	total = m.Sys / 1024 / 1024                 // Convert to MB
	avail = (m.Sys - m.HeapAlloc) / 1024 / 1024 // Convert to MB
	return
}

func getCurrentCpuUsage() (usage float64) {
	// Get initial CPU stats
	before, err := cpu.Get()
	if err != nil {
		fmt.Println("Error getting CPU stats:", err)
		return
	}

	// Wait for a short duration to capture CPU usage change
	time.Sleep(time.Second)

	// Get CPU stats again
	after, err := cpu.Get()
	if err != nil {
		fmt.Println("Error getting CPU stats:", err)
		return
	}

	// Calculate CPU usage based on stat differences
	total := float64(after.Total - before.Total)
	idle := float64(after.Idle - before.Idle)
	usage = (total - idle) / total * 100
	fmt.Printf("CPU Usage: %.2f%%\n", usage)
	return
}
