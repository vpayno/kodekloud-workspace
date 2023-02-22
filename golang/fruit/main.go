/*
	We are working on a project for a grocery store.
	The store has a list of items that are currently in stock, and we need to implement a function that will return a list of items that are in a given price range.

	A Go file is located at /root/code/fruit directory.

	Expected Output:

		[{Apple 0.5} {Banana 0.25} {Orange 0.75}]
		[{Apple 0.5} {Orange 0.75}]
		[{Orange 0.75} {Pineapple 1.5}]
*/

package main

import "fmt"

type Item struct {
	Name  string
	Price float64
}

func getItemsInPriceRange(items []Item, minPrice, maxPrice float64) []Item {
	var result []Item

	for _, item := range items {
		if item.Price >= minPrice && item.Price < maxPrice {
			result = append(result, item)
		}
	}

	return result
}

func main() {
	items := []Item{
		{Name: "Apple", Price: 0.5},
		{Name: "Banana", Price: 0.25},
		{Name: "Orange", Price: 0.75},
		{Name: "Pineapple", Price: 1.5},
	}

	fmt.Println(getItemsInPriceRange(items, 0.0, 1.0))
	fmt.Println(getItemsInPriceRange(items, 0.5, 1.0))
	fmt.Println(getItemsInPriceRange(items, 0.75, 1.5))
}
