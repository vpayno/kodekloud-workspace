/*
	What would be the output of the following code?

	Given code is written in a Go file called main.go inside /root/code/struct directory.
*/

package main

import "fmt"

type Product struct {
	name     string
	quantity int
	price    float64
}

func printName(p Product) {
	fmt.Println(p.name)
}

func printQuantity(p Product) {
	fmt.Println(p.quantity)
}

func printPrice(p Product) {
	fmt.Println(p.price)
}

func main() {
	var p Product
	p.name = "Chair"
	p.quantity = 5
	p.price = 700

	fmt.Println("Product details:")
	printName(p)
	defer printPrice(p)
	printQuantity(p)
}

/*
	A.
		Product details:
		5
		700
		Chair

	B.
		Product details:
		5
		Chair
		700

	C.
		Product details:
		700
		Chair
		5

	D. *
		Product details:
		Chair
		5
		700
*/
