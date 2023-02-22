/*
	1
	2
	3
	4
	5
	6
	7
	8
	9
	10

	skip_next
	Select the correct base case for function printSquares -

	for input n -> prints squares for n, n-1, n-2, …… -5

	Example:

		Input: n=2; Output: 4 1 0 1 4 9 16 25
*/

package main

import "fmt"

func printSquares(n int) {
	// base case
	if n == -6 {
		return
	}

	fmt.Printf("%d ", n*n)
	printSquares(n - 1)
}

func main() {
	fmt.Println("4 1 0 1 4 9 16 25")
	printSquares(2)
}

/*
	base cases

	if n == 0 { return }

	if n == -5 { return }

	if n == -6 { return }

	if n == 5 { return }
*/
