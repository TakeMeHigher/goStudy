package main

import "fmt"

func sumTwo(a [5]int, target int) {
	for i := 0; i < len(a); i++ {
		b := target - a[i]
		for j := i + 1; j < len(a); j++ {
			if b == a[j] {
				fmt.Printf("(%d, %d)\n", i, j)
			}
		}
	}
}

func main() {
	var a [5]int = [5]int{1, 2, 3, 7, 5}
	sumTwo(a, 8)

}
