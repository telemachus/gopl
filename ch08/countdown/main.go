package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan bool)
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- false
	}()
	fmt.Println("Commencing countdown. Press return to abort.")
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Printf("%d…\n", countdown)
		select {
		case <-ticker.C:
			// Do nothing.
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
