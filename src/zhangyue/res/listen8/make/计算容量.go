// package main

// import "fmt"

// func testCap() {
// 	a := [...]int{1, 2, 3, 4, 5, 6, 7}
// 	b := a[1:3]
// 	fmt.Print("b= %v  len(b) = %d, cap(b) = %d\n", b, len(b), cap(b)) // 切片的容量 等于 原始素组的长度 减去 切片开始的索引
// }

// func main() {
// 	testCap()
// }

package main

import (
	"fmt"
)

func main() {
	cars := []string{"Ferrari", "Honda", "Ford"}
	//⻓度和容量都等于3
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars))
	cars = append(cars, "Toyota")
	//容量等于6
	fmt.Println("cars:", cars, "has new length",
		len(cars), "and capacity", cap(cars))
}
