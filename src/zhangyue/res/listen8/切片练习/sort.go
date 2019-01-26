package main

import (
	"fmt"
	"sort"
)

func sortTest() {
	a := [...]int{1, 5, 6, 3, 4, 9}
	sort.Ints(a[:])
	fmt.Println("a:", a)

	var b [5]string = [5]string{"ac", "ec", "be", "fa", "ii"}
	sort.Strings(b[:])
	fmt.Println("b:", b)

	var c [5]float64 = [5]float64{29.38, 22.32, 0.8, 99191.2}
	sort.Float64s(c[:])

	fmt.Println("c:", c)

}
func main() {
	sortTest()
}
