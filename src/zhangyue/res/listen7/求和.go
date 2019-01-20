package main

import (
	"fmt"
	"math/rand"
	"time"
)

func test1(a [5]int) int {
	var sum int
	// for i := 0; i < len(a); i++ {
	// 	sum += a[i]
	// }

	for _, v := range a {
		sum += v
	}
	return sum
}

func run() int {
	rand.Seed(time.Now().Unix())  // 不设置 这个的话 rand.Intn()  取到的数据 都是 一样的  
	var a [5]int
	for i := 0; i < len(a); i++ {
		a[i] = rand.Intn(100)  // rand.Intn(n)  随机取出 0 - n-1 的整数
		// a[i] = rand.Int()  // rand.Intn()  随机取出 0- int最大范围之间的 的整数
	}

	sum := test1(a)
	return sum
}

func main() {
	sum := run()
	fmt.Println(sum)
}
