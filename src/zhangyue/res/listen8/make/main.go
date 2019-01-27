package main

import "fmt"

func testMake1() {
	var a []int
	a = make([]int, 5, 10)
	a[0] = 10
	//a[1] = 20
	fmt.Printf("a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))
	a = append(a, 11)
	fmt.Printf("a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))

	for i := 0; i < 8; i++ {
		a = append(a, i)
		fmt.Printf("a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))
	}
	//观察切片的扩容操作，扩容的策略是翻倍扩容
	a = append(a, 1000)
	fmt.Printf("扩容之后的地址：a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))
}

/*
a=[10 0 0 0 0] addr:0xc0000860a0 len:5 cap:10
a=[10 0 0 0 0 11] addr:0xc0000860a0 len:6 cap:10
a=[10 0 0 0 0 11 0] addr:0xc0000860a0 len:7 cap:10
a=[10 0 0 0 0 11 0 1] addr:0xc0000860a0 len:8 cap:10
a=[10 0 0 0 0 11 0 1 2] addr:0xc0000860a0 len:9 cap:10
a=[10 0 0 0 0 11 0 1 2 3] addr:0xc0000860a0 len:10 cap:10
a=[10 0 0 0 0 11 0 1 2 3 4] addr:0xc000098000 len:11 cap:20
a=[10 0 0 0 0 11 0 1 2 3 4 5] addr:0xc000098000 len:12 cap:20
a=[10 0 0 0 0 11 0 1 2 3 4 5 6] addr:0xc000098000 len:13 cap:20
a=[10 0 0 0 0 11 0 1 2 3 4 5 6 7] addr:0xc000098000 len:14 cap:20
扩容之后的地址：a=[10 0 0 0 0 11 0 1 2 3 4 5 6 7 1000] addr:0xc000098000 len:15 cap:20
*/

func testMake2() {
	var a []int
	a = make([]int, 5, 10)
	//a[5] = 100
	a = append(a, 10)
	fmt.Printf("a=%v\n", a)

	b := make([]int, 0, 10)
	fmt.Printf("b=%v len:%d cap:%d\n", b, len(b), cap(b))
	b = append(b, 100)
	fmt.Printf("b=%v len:%d cap:%d\n", b, len(b), cap(b))
}

func main() {
	// testMake1()
	testMake2()
}
