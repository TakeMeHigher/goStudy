package main

import "fmt"

func test1() {
	var a int = 100
	var b *int
	b = &a
	fmt.Printf("a 的值为: %v  a的内存地址是:%p  a的类型是%t\n", a, a, a)
	fmt.Printf("b 中存储的内存地址对应的值为%v\n", *b)
	fmt.Printf("b 的值为: %v  b的内存地址是:%p  b的类型是%t\n", b, b, b)

	a = 200
	fmt.Printf("a 的值为: %v  a的内存地址是:%p  a的类型是%t\n", a, a, a)
	fmt.Printf("b 中存储的内存地址对应的值为%v\n", *b)
	fmt.Printf("b 的值为: %v  b的内存地址是:%p  b的类型是%t\n", b, b, b)

}

/*
a 的值为: 100  a的内存地址是:%!p(int=100)  a的类型是%!t(int=100)
b 中存储的内存地址对应的值为100
b 的值为: 0xc000058080  b的内存地址是:0xc000058080  b的类型是%!t(*int=0xc000058080)
a 的值为: 200  a的内存地址是:%!p(int=200)  a的类型是%!t(int=200)
b 中存储的内存地址对应的值为200
b 的值为: 0xc000058080  b的内存地址是:0xc000058080  b的类型是%!t(*int=0xc000058080)
*/

func test2(a *int) {
	*a = 200
}

func test3(a *[3]int) {
	(*a)[0] = 1000
}

func test4(a []int){
	a[0] = 123

}

func modifytest() {
	// var a int = 2
	// test2(&a)
	// fmt.Printf("修改后a的值为 %d", a) //修改后a的值为 200

	// a := [3]int{1, 2, 3}
	// test3(&a)
	// fmt.Printf("修改后a的值为 %v", a) // 修改后a的值为 [1000 2 3]


	a:= [5]int{1,5,6,4,5}
	test4(a[:])
	fmt.Printf("修改后a的值为 %v", a) // 修改后a的值为 [123 5 6 4 5]
}

func main() {
	// test1()

	modifytest()
}
