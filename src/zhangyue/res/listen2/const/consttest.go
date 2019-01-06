package main

import  "fmt"

func main(){
	/*
	常量 定义时一定要给定值 不给值 报错
    常量一旦定义 不可修改
	*/

	// 方式一
	// const a int = 10
	// const b = "str"
	
	// 方式二
	const(
		a int =10
		b = "str"
	)
	 
	fmt.Printf("a=%d b=%s\n", a , b)

	// 下面这种写法如果后面不给定值 则后面常量的值 和前面的一样
	const (
		c int = 200
		d
		e = 300
		f
	)

	fmt.Printf("c=%d d=%d e=%d f=%d \n", c, d, e, f)
	//=200 d=200 e=300 f=300
   
	// iota 开始为0 依次向下递增1
	const(
		a1 = iota
		a2
		a3
		a4
	)

	fmt.Printf("a1=%d a2=%d a3=%d a4=%d \n", a1, a2, a3, a4)
	// a1=0 a2=1 a3=2 a4=3

	const(
		b1 = iota
		b2 = iota
		b3 = iota
		b4 = iota
	)
	fmt.Printf("b1=%d b2=%d b3=%d b4=%d \n", b1, b2, b3, b4)
	// b1=0 b2=1 b3=2 b4=3


	const(
		c1 = 1 << iota
		c2
		c3
		c4
	)
	fmt.Printf("c1=%d c2=%d c3=%d c4=%d \n", c1, c2, c3, c4)
	// c1=1 c2=2 c3=4 c4=8
	const(
		d1 = 1 << iota
		d2 = 1 << iota
 		d3 = 1 << iota
		d4 = 1 << iota
	)
	fmt.Printf("d1=%d d2=%d d3=%d d4=%d \n", d1, d2, d3, d4)
	//d1=1 d2=2 d3=4 d4=8
}