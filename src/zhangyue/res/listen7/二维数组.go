package main

import "fmt"

func test1() {
	var a [3][2]int
	a[0][0] = 1
	a[0][1] = 2
	a[1][0] = 3
	a[1][1] = 4
	a[2][0] = 5
	a[2][1] = 6

	fmt.Println(a) // [[1 2] [3 4] [5 6]]
}

func test2() {
	var a [3][2]int = [3][2]int{{1, 2}, {3, 4}, {5, 6}}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
		fmt.Printf("\n")
	}
}

func test3(){
	var a [3][2]int = [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	for index, val := range(a){
		for i,v:= range(val){
			fmt.Printf("a[%d][%d] = %d\n", index, i ,v)
		}
		fmt.Printf("\n")
	}
}

func main() {
	// test1()
	// test2()
	test3()
}
