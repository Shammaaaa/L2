package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(14*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
		sig(2*time.Second),
	)

	fmt.Printf("fone after %v\n", time.Since(start))
}

func or(channels ...<-chan any) <-chan any {
	out := make(chan any)

	for _, c := range channels {
		go func(c <-chan any) {
			for val := range c {
				out <- val
			}
		}(c)
	}

	return out
}

func sig(d time.Duration) <-chan any {
	out := make(chan any)
	go func() {
		out <- <-time.After(d)
	}()
	return out
}
