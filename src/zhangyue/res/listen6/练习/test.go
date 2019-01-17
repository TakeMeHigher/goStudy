package main

import "fmt"

func insert_sort(li [8]int) [8]int {
	for i := 1; i < len(li); i++ {
		for j := i; j > 0; j-- {
			if li[j] < li[j-1] {
				li[j], li[j-1] = li[j-1], li[j]
			} else {
				break
			}
		}
	}

	return li
}

func select_sort(li [8]int) [8]int {
	for i := 0; i < len(li); i++ {
		for j := i + 1; j < len(li); j++ {
			if li[j] < li[i] {
				li[j], li[i] = li[i], li[j]
			}
		}
	}
	return li
}

func mp_sort(li [8]int) [8]int {
	for i := 0; i < len(li); i++ {
		for j := 0; j < len(li)-i-1; j++ {
			if li[j] > li[j+1] {
				li[j], li[j+1] = li[j+1], li[j]
			}
		}
	}
	return li
}

func main() {
	var a [8]int = [8]int{5, 6, 3, 2, 4, 8, 1, 3}
	// li := insert_sort(a)
	// li := select_sort(a)
	li := mp_sort(a)
	fmt.Println(li)
}
