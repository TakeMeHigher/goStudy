package main

import (
	"fmt"
	"strings"
)

func jbyf() {
	var a string = "hello"
	var b = "world"

	c := "hello world "

	fmt.Println(a, b, c)

	var d = "an\nc" // 有特殊字符 则转义

	fmt.Println(d)

	d = `an\nc` // 原样输出

	fmt.Println(d)
}

func commonMethod() {
	aa := "lol is bad, but i like play lol"
	// 长度
	aLen := len(aa)
	fmt.Printf("aa的长度:%d\n", aLen) //aa的长度:31

	// 字符串拼接

	aaStr := "hello" + "world"
	fmt.Println(aaStr)                              // helloworld
	bbStr := fmt.Sprintf("%s-%s", "hello", "china") // hello-china
	fmt.Println(bbStr)

	// 字符串分割

	cc := "i-n-j"

	dd := strings.Split(cc, "-")
	fmt.Println(dd)    // [i n j]
	fmt.Println(dd[1]) // n

	// 包含

	if strings.Contains(aa, "lol") {
		fmt.Println("aa is contains lol")
	} else {
		fmt.Println("aa is not contains lol")
	}

	// 前缀和后缀判断

	if strings.HasPrefix(aa, "lol") {
		fmt.Println("aa is start lol")
	} else {
		fmt.Println("aa is not start lol")
	}

	if strings.HasSuffix(aa, "lol") {
		fmt.Println("aa is end lol")
	} else {
		fmt.Println("aa is not end lol")
	}

	// 自字符串出现的位置

	fmt.Println(strings.Index(aa, "lol"))
	fmt.Println(strings.LastIndex(aa, "lol"))

	// Join 操作
	var strArr [] string = [] string{"aa", "bb", "cc"}
	fmt.Println(strings.Join(strArr, "*"))
}

func main() {
	// jbyf()
	commonMethod()
}
