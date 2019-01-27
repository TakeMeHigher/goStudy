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

/*
addr = 0x0 len=0  cap =0
a is nil
addr = 0xc0000100e0 len=1  cap =1
addr = 0xc000010100 len=2  cap =2
addr = 0xc00000a3c0 len=3  cap =4
addr = 0xc00000a3c0 len=4  cap =4
*/

func main() {
	test()
}
