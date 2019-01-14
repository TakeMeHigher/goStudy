package main

import "fmt"

func add(a, b int) int {
	return a + b

}

// 函数名也可以赋值给变量
func test1() {
	ff := add
	fmt.Printf("ff的类型是%T\n", ff)
	sum := ff(4, 5)
	fmt.Printf("sum is %d \n", sum)
}

/*
ff的类型是func(int, int) int
sum is 9
*/

// 匿名函数
func test2() {
	f1 := func(a, b int) int {
		return a + b
	}
	fmt.Printf("f1的类型是%T\n", f1)
	sum := f1(4, 5)
	fmt.Printf("sum is %d \n", sum)
}

/*
f1的类型是func(int, int) int
sum is 9
*/

// defer
func test3() int {
	i := 10
	defer fmt.Printf("defer i is %d\n", i)
	i = 100
	fmt.Printf("i is %d\n", i)
	return i
}

/*
i is 100
defer i is 10
return i is 100
*/

// defer 和匿名函数
func test4 () int {
	i:= 10
	defer func() {
		fmt.Printf("defer i is %d \n", i)
	}()

	i= 100
	fmt.Printf("i is %d\n", i)

	return i
}


func main() {
	// test1()

	// test2()

	// i := test3()
	// fmt.Printf("return i is %d", i)

	i := test4()
	fmt.Printf("return i is %d", i)
}
