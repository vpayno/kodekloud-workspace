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

	for i := 1; i <= 10; i++ {
		calculateSquare(i)
	}

	elapsedTime := time.Since(startTime)

	fmt.Printf("Function took %s seconds to run.\n", elapsedTime)
}
