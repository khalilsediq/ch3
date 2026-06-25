package main

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

	// s := []int{10, 20, 30}
	// s2 := make([]int, len(s))
	// copy(s2, s)
	// fmt.Println(s)
	// fmt.Println(s2)

	s1 := []int{1, 2, 3}
	s2 := s1[0:2]             // Shared view. s2 is now [1,2]
	s3 := append(s2, 4, 5, 6) // Appending that triggers reallocation. s3 is [1,2,4,5,6]
	s3[0] = 99                //s3 is [99,2,4,5,6], s2[1,2,4] and it completely exhausted on it's own capacity, so s3 created it's own completely new slice and added these [99, 2,4,5,6]

	s4 := make([]int, 2) // so s4 is [0,0]
	copy(s4, s1)         // Independent copy so now s4 is [1,2]
	s4[1] = 88           // and the value of s4's 1st indes is 88 which is [1 , 88].
	//So in the conclusion is that s2 becomes [1,2] and then the s3 appends into s2 and the s2 still curently holds one capacity so it drops 4 into himself, which he becomes s2[1,2,4]. And s3 becomes his own independent slice, so he holds s3[1,2,4,5,6]. after that we do s3[0] = 99 which the s3 0 index value becomes 99, so now s3 is [99, 2, 4,5,6]. and finaly we come to s4 which holds 2 values of [0, 0], and then we copy s1 into it so now s4 is [1,2], and then we do s4[1] = 88, which changes the s4 to [2,88]. that's it.



	var age = 12
	var name = "John"
	var isStudent = true

	if age > 18 && isStudent {
		println(name + " is an adult student.")
	} else if age > 18 {
		println(name + " is an adult.")
	} else {
		println(name + " is not an adult. and is infact a minor and should be banned")
	}
}
