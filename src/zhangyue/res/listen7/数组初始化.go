package main

import "fmt"

// 数组初始化

func test1() {
	var a [5]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a) // [1 2 3 0 0]

}

func test2() {
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(a) // [1 2 3 4 5]
}

func test3() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a) // [1 2 3 4 5]
}

func test4() {
	a := [...]int{1, 2, 3, 4, 5}
	fmt.Println(a) // [1 2 3 4 5]
}

func test5() {
	a := [3]int{5}
	fmt.Println(a) //[5 0 0]
}

func test6(){
	a := [5]int{3:30, 2:20}
	fmt.Println(a)  //[0 0 20 30 0]
}

func main() {
	// test1()
	// test2()
	// test3()
	// test4()
	// test5()
	test6()

}
