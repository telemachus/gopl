package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := runtime.GOMAXPROCS(0)
	fmt.Printf("number of CPUs = %d\n", n)
}
