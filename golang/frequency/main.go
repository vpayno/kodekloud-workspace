/*
	In this question, we are building a tool that analyzes the frequency of words in a given text.
	You need to implement a function wordFrequency that receives a string and returns a map with the frequency of each word in the string.

	A Go file is located at /root/code/frequency directory.

	Expected output:

		map[The:1 brown:1 dog:1 fox:1 jumps:1 lazy:1 over:1 quick:1 the:1]
*/

package main

import (
	"fmt"
	"strings"
)

func wordFrequency(text string) map[string]int {
	frequency := map[string]int{}

	for _, word := range strings.Fields(text) {
		frequency[word] += 1
	}

	return frequency
}

func main() {
	text := "The quick brown fox jumps over the lazy dog"
	fmt.Println(wordFrequency(text))
}
