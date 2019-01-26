package main

import "fmt"

// 找出数组中和为给定值的两个元素的下标，例如数组:[1,3,5,8,7]，找出两个元素之和等于8的下标分别是(0, 4)和(1,2)

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
