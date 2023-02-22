/*
	We are building a software for a store that sells three products: apples, bananas, and oranges.
	We need to write a function that takes the name of a product and its price as arguments and returns the price of the product with a discount applied.
	The discount should be 10% for apples and 20% for bananas. Oranges do not have a discount.

	A Go file is located at /root/code/product directory.

	Expected Output:

		90
		100
		80
		100
		100
*/

package main

import "fmt"

func discountedPrice(product string, price float64) float64 {
	var result float64

	switch product {
	case "apples":
		return (1 - 0.10) * price
	case "bananas":
		return (1 - 0.20) * price
	default:
		return price
	}

	return result
}

func main() {
	fmt.Println(discountedPrice("apples", 100))
	fmt.Println(discountedPrice("orange", 100))
	fmt.Println(discountedPrice("bananas", 100))
	fmt.Println(discountedPrice("banana", 100))
	fmt.Println(discountedPrice("oranges", 100))
}
