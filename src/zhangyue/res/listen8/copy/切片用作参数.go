package main

import "fmt"

func testSum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func testChange(a []int) {
	a[0] = 1000
}

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6}
	// sum := testSum(a[:])
	// fmt.Printf("sum = %d", sum) // 21

	fmt.Printf("a=%v\n", a)
	testChange(a[:])
	fmt.Printf("a=%v", a)

	/*
	a=[1 2 3 4 5 6]
    a=[1000 2 3 4 5 6]
	*/
}
