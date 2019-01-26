package main

import "fmt"

func reslice() {
	a := [...]int{1, 2, 3, 4, 5, 6, 4, 7, 8}
	b := a[2:4]
	fmt.Printf("b = %d len(b) = %d cap(b)=%d\n", b, len(b), cap(b))
	c := b[:cap(b)]
	fmt.Printf("c = %d len(c) = %d cap(c)=%d\n", c, len(c), cap(c))
}

func main() {
	reslice()
}
