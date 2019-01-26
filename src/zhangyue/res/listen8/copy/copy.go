package main

import "fmt"

func testCopy() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	copy(a, b)
	fmt.Printf("a = %v\n", a) // [4,5,6]
	fmt.Printf("b = %v\n", b) // [4,5,6]

	c := []int{1, 2}
	d := []int{4, 5, 6}

	e := copy(c, d)
	fmt.Printf("c = %v\n", c) // [4,5]
	fmt.Printf("d = %v\n", d) // [4,5,6]
	fmt.Printf("e = %v\n", e) // 2

	b[0] = 100
	fmt.Printf("a = %v\n", a) // [4,5,6]
	fmt.Printf("b = %v\n", b) // [100 5 6]

	aa := [...]int{1, 2, 3, 4, 5, 6}
	bb := aa[:]
	cc := []int{10, 20, 30}
	copy(bb, cc)
	fmt.Printf("aa = %v\n", aa) // [10 20 30 4 5 6]
	fmt.Printf("bb = %v\n", bb) // [10 20 30 4 5 6]
	fmt.Printf("cc = %v\n", cc) // [10 20 30]

}

func main() {
	testCopy()
}
