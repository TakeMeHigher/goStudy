package main

import "fmt"

// 定义指针
func test1() {
	var a int = 100
	var b *int
	fmt.Printf("a 的值为: %v  a的内存地址是:%p  a的类型是%t\n", a, a, a)
	fmt.Printf("b 的值为: %v  b的内存地址是:%p  b的类型是%t\n", b, b, b)

	b = &a
	fmt.Printf("b 的值为: %v  b的内存地址是:%p  b的类型是%t\n", b, b, b)
}

/*
a 的值为: 100  a的内存地址是:%!p(int=100)  a的类型是%!t(int=100)
b 的值为: <nil>  b的内存地址是:0x0  b的类型是%!t(*int=<nil>)
b 的值为: 0xc0000100a8  b的内存地址是:0xc0000100a8  b的类型是%!t(*int=0xc0000100a8)
*/

//控指针
func test2() {
	var a int = 100
	var b *int
	fmt.Printf("b 的值为: %v  b的内存地址是:%p  b的类型是%t\n", b, b, b)
	if b == nil {
		fmt.Println("b是一个空指针")
		b = &a
	}
	fmt.Printf("b 的值为: %v  b的内存地址是:%p  b的类型是%t\n", b, b, b)
}

/*
b 的值为: <nil>  b的内存地址是:0x0  b的类型是%!t(*int=<nil>)
b是一个空指针
b 的值为: 0xc000058080  b的内存地址是:0xc000058080  b的类型是%!t(*int=0xc000058080)
*/

// 获取指针对应地址的值
func test3() {
	var a int = 100
	var b *int
	b = &a
	fmt.Printf("b 的值为: %v  b的内存地址是:%p  b的类型是%t\n", b, b, b)
	fmt.Printf("b 中存储的内存地址对应的值为%v\n", *b)

}

/*
b 的值为: 0xc000058080  b的内存地址是:0xc000058080  b的类型是%!t(*int=0xc000058080)
b 中存储的内存地址对应的值为100
*/

// 改变指针存储地址对应的值
func test4() {
	var a int = 100
	var b *int
	b = &a
	fmt.Printf("b 中存储的内存地址对应的值为%v\n", *b)
	fmt.Printf("a的值为%v\n", a)
	*b = 500
	fmt.Printf("b 中存储的内存地址对应的值为%v\n", *b)
	fmt.Printf("a的值为%v\n", a)

}

/*
b 中存储的内存地址对应的值为100
a的值为100
b 中存储的内存地址对应的值为500
a的值为500
*/

func main() {
	// test1()
	// test2()
	// test3()
	test4()
}
