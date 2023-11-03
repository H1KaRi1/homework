package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	a := 40
	b := 2
	sum := add(a, b)
	fmt.Println(sum)
}
