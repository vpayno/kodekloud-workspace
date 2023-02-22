/*
	In this question, we are building a simple calculator that can perform addition, subtraction, multiplication, and division operations.

	Complete the code in the function calculate to return a slice consisting of 4 elements
		[ sum of a and b, difference of a and b, product of a and b, quotient on dividing a by b]

	A Go file is located at /root/code/calculator/ directory.

	Expected Output:

		[30 10 200 2]
		[770 630 49000 10]
*/

package main

import "fmt"

func calculate(a int, b int) []float64 {
	results := []float64{}
	results = append(results, sum(a, b))
	results = append(results, difference(a, b))
	results = append(results, product(a, b))
	results = append(results, quotient(a, b))
	return results
}

func sum(a, b int) float64 {
	return float64(a + b)
}

func difference(a, b int) float64 {
	return float64(a - b)
}

func product(a, b int) float64 {
	return float64(a * b)
}

func quotient(a, b int) float64 {
	return float64(a / b)
}

func main() {
	fmt.Println(calculate(20, 10))
	fmt.Println(calculate(700, 70))
}
