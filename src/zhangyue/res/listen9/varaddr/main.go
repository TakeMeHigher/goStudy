package main

import "fmt"

func getVarAddr() {
	var a int = 10000
	fmt.Printf("a 的内存地址 为%p", &a) //a 的内存地址 为0xc0000580a0
}

func main() {
	getVarAddr()
}
