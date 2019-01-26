package main

import "fmt"

func testCap() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7}
	b := a[1:3]
	fmt.Print("b= %v  len(b) = %d, cap(b) = %d\n", b, len(b), cap(b))
}

func main() {
	testCap()
}
