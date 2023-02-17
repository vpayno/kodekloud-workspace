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
