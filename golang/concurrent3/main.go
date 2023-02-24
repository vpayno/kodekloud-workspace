package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("entering main()")

	go func() {
		fmt.Println("In anonymous function")
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("exiting main()")
}
