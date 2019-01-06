package main

import "fmt"


const (
	a = iota
	b
	c
	d
	e = 8
	f
	g = iota
	h
	k
)


const (
	A = iota
	B
	C
)


func main(){
	fmt.Printf("a= %d\n", a)
	fmt.Printf("b= %d\n", b)
	fmt.Printf("c= %d\n", c)
	fmt.Printf("d= %d\n", d)
	fmt.Printf("e= %d\n", e)
	fmt.Printf("f= %d\n", f)
	fmt.Printf("g= %d\n", f)
	fmt.Printf("h= %d\n", h)
	fmt.Printf("k= %d\n", k)
	fmt.Printf("A= %d\n", A)
	fmt.Printf("B= %d\n", B)
    fmt.Printf("C= %d\n", C)
}
/*
a= 0
b= 1
c= 2
d= 3
e= 8
f= 8
g= 8
h= 7
k= 8
A= 0
B= 1
C= 2
*/

/*
const() 中 如果 定义了 常量 不给定默认值 那么 该常量的表达式 和 其上一个 表达式是相同的
iota 在const 中 没经过一层 自增1
*/
