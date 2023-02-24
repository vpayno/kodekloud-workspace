package main

import (
	"fmt"
	"time"
)

func calculateSquare(i int) {
	time.Sleep(1 * time.Second)
	fmt.Println(i * i)
}

func main() {
	startTime := time.Now()

	for i := 1; i <= 10_000; i++ {
		go calculateSquare(i)
	}

	// don't do this in reall programs
	time.Sleep(2 * time.Second)

	elapsedTime := time.Since(startTime)

	fmt.Printf("Function took %s seconds to run.\n", elapsedTime)
}
