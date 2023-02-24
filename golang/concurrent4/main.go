package main

import (
	"fmt"
	"sync"
	"time"
)

func calculateSquare(i int, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println(i * i)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	startTime := time.Now()

	for i := 1; i <= 10_000; i++ {
		wg.Add(1)
		go calculateSquare(i, &wg)
	}

	wg.Wait()

	elapsedTime := time.Since(startTime)

	fmt.Printf("Function took %s seconds to run.\n", elapsedTime)
}
