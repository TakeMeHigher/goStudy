package main

import "fmt"

func justfy(i int) bool {
	if i <= 1 {
		return false
	}

	for j := 2; j <= i/2; j++ {
		if i%j == 0 {
			return false
		}
	}

	return true
}

func exp() {
	for i := 1; i <= 100; i++ {
		if justfy(i) {
			fmt.Printf("%d is 质数\n", i)
		}
	}
}

func main() {
	exp()
}
