package main

import (
	"fmt"
)

func main(){
	// 方式一
	/*  
	var a int 
	var b string
	var c bool
	var d float32
	*/

	// 方式二

	var(
		a int 
		b string 
		c bool
		d float32
	)

	fmt.Printf("a=%d b=%s  c=%t d=%f\n", a, b, c, d)

	a = 12
	b = "ctz"
	c = true
	d = 10.36

	fmt.Printf("a=%d b=%s  c=%t d=%f", a, b, c, d)
}