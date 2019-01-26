package main

import "fmt"

func test1() {
	fmt.Println("hello world")
}

func sumtwo(a int, b int) int {
	sum := a + b
	return sum
}

func sumsub(a, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub
}

// 队返回值进行命名
func test3(a, b int) (sub, sum int) {
	sub = a - b
	sum = a + b
	return
}

func test4(b ...int) int {
	sum := 0
	for i := 0; i < len(b); i++ {
		sum += b[i]
	}
	return sum
}

func main() {

	// sum := sumtwo(10, 12)
	// fmt.Println(sum)

	// sum, sub := sumsub(20, 13)
	// fmt.Printf("sum=%d sub=%d", sum, sub)

	// sub, sum := test3(20, 13)
	// fmt.Printf("sub=%d sum=%d", sub, sum)
	sum := test4(1, 2, 3, 4)
	fmt.Println(sum)
}
