package main

import (
	"fmt"
	"runtime"
)

func stats() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Printf("alloc [%v] \t heapAlloc [%v] \n", mem.Alloc, mem.HeapAlloc)
	fmt.Println("Routines: ", runtime.NumGoroutine())
}
