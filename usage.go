package dbg

import (
	"fmt"
	"runtime"
	"time"
)

func MemUsageRun(interval time.Duration) {
	go func() {
		for {
			PrintMemUsage()
			time.Sleep(interval)
		}
	}()
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	Green(fmt.Sprintf(`	Alloc = %v
	TotalAlloc = %v
	Sys = %v
	NumGC = %v`, m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC))
}
