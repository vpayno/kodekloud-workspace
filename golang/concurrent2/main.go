package main

import (
	"fmt"
	"time"
)

func start() {
	fmt.Println("entering start()")
	go process()
	fmt.Println("exiting start()")
}

func process() {
	fmt.Println("entering process()")
	fmt.Println("exiting process()")
}

func main() {
	fmt.Println("entering main()")
	go start()
	time.Sleep(1 * time.Second)
	fmt.Println("exiting main()")
}
