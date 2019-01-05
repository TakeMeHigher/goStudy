package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func aa(a int, b int) (int, int) {
	return a +b , a-b
}

func main(){
	sum := add(3, 6)
    he, cha := aa(6, 9)
	fmt.Println(sum)
	fmt.Println(he, cha)
}