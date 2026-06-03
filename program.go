package main

import "fmt"

func main() {
	// fmt.Println("Hello, World!")
	// x := make([]int, 0, 1000)
	// for i := 0; i < cap(x); i++ {
	// 	x = append(x, i+1)
	// }
	// y := append(x, 1001)
	// fmt.Println("The Items in X are:", x, "The Length of X is:", len(x), "The Capacity of X is:", cap(x))
	// fmt.Println("The Items in Y are:", y, "The Length of Y is:", len(y), "The Capacity of Y is:", cap(y))
	emojeTxt()
	numbers()
}

func emojeTxt() {
	// var s string = "Hello 😀"
	// var s2 string = s[6:7]
	// var s3 string = s[:5]
	// var s4 string = s[6:]
	// fmt.Println(s2)
	// fmt.Println(s3)
	// fmt.Println(s4)

	// var s string = "H2"
	// fmt.Println(len(s))
}

func numbers() {
	// x := []int{1, 2, 3}
	// b := append(x, 4)
	// b = append(b, 5, 6, 7, 8)
	// x[2] = 10
	// fmt.Println(x)
	// fmt.Println(b)

	s := []int{10, 20, 30}
	s2 := make([]int, len(s))
	copy(s2, s)
	fmt.Println(s)
	fmt.Println(s2)
}
