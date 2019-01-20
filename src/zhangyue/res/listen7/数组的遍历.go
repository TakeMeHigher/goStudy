package main

import "fmt"

func test1() {
	a := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%d]=%d\n", i, a[i])
	}
}


func test2(){
	a := [5]int{1, 2, 3, 4, 5}
	for index, val := range a {
		fmt.Printf("a[%d]=%d\n", index, val)
	}
}

func main() {
	// test1()
	test2()
}
