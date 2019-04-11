package main

import (
	"time"
	"fmt"
)

func main() {
	go spinner(100 * time.Millisecond)
	res := fib(45)
	fmt.Printf("\rFibonacci result:%d\n", res)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
