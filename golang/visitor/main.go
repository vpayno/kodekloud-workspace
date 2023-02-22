/*
	We are building a program to keep track of the number of visitors to a website.

	We need to store the number of active visitors in a variable and update it each time a new visitor arrives or an old visitor leaves the website.

	A Go file is located at /root/code/visitor directory.

	Expected Output:

		2
*/

package main

import "fmt"

// Declare variable activeUserCount
var activeUserCount int

func entry() {
	// Hint: you can use the "++" operator to increment a variable by 1
	activeUserCount++
}

func exit() {
	// Hint: you can use the "--" operator to decrement a variable by 1
	activeUserCount--
}

func main() {
	entry()
	entry()
	exit()
	exit()
	entry()
	entry()
	fmt.Println(activeUserCount)
}
