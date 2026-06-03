package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	x := make([]int, 0, 1000)
	for i := 0; i < cap(x); i++ {
		x = append(x, i+1)
	}
	y := append(x, 1001)
	fmt.Println("The Items in X are:", x, "The Length of X is:", len(x), "The Capacity of X is:", cap(x))
	fmt.Println("The Items in Y are:", y, "The Length of Y is:", len(y), "The Capacity of Y is:", cap(y))
}
