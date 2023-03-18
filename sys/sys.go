package sys

import (
	"fmt"
	"runtime"
)

func MemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("Alloc = %v", m.Alloc)
	fmt.Println("TotalAlloc = %v", m.TotalAlloc)
	fmt.Println("Sys = %v", m.Sys)
	fmt.Println("NumGC = %v", m.NumGC)

}
