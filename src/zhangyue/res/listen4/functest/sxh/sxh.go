package main

import "fmt"

func issxh(i int) bool {
	ge := i % 10
	shi := (i / 10) % 10
	bai := i / 100

	sum := ge*ge*ge + shi*shi*shi + bai*bai*bai

	if sum == i {
		return true
	} else {
		return false
	}

}

func exp() {
	for i := 100; i <= 999; i++ {
		if issxh(i) {
			fmt.Printf("%d 是水仙花数\n", i)
		}
	}
}

func main() {
	// fmt.Println(111 / 10)
	exp()
}
