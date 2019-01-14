package main

import (
	"fmt"
	"zhangyue/res/listen6/可见性"
)

func main() {
	sum := 可见性.Add(100, 200)
	fmt.Printf("cacl Add return %d\n", sum)

	// sub := 可见性.sub(200,100)   // 报错 说明 小写是私有的 只能在包内使用 不能在保外使用
	// fmt.Printf("cacl sub return %d\n", sub)

	fmt.Printf(" cacl A is %d \n", 可见性.A)
	// fmt.Printf(" cacl a is %d \n", 可见性.a)  // 报错

	// 如果想得到小a  可用下面这种方法
	a := 可见性.Test()
	fmt.Printf(" cacl a is %d \n", a)
}
