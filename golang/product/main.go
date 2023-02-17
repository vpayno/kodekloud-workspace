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
