package main

import "fmt"

func test() {
	var a []int

	fmt.Printf("addr = %p len=%d  cap =%d\n", a, len(a), cap(a))
	// a[0]=100  报错

	if a == nil {
		fmt.Printf("a is nil\n")
	}

	a = append(a, 100)
	fmt.Printf("addr = %p len=%d  cap =%d\n", a, len(a), cap(a))
	a = append(a, 200)
	fmt.Printf("addr = %p len=%d  cap =%d\n", a, len(a), cap(a))
	a = append(a, 300)
	fmt.Printf("addr = %p len=%d  cap =%d\n", a, len(a), cap(a))
	a = append(a, 400)
	fmt.Printf("addr = %p len=%d  cap =%d\n", a, len(a), cap(a))

}

func main() {
	test()
}
