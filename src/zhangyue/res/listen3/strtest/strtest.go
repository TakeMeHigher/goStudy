package main

import "fmt"

// 写一个程序，对英问字符串进行逆序。
func ewreverse() string {
	var aa = "hello world"
	var bb []byte
	bb = []byte(aa)
	for i := 0; i < len(aa)/2; i++ {
		tem := bb[len(aa)-i-1]
		bb[len(bb)-i-1] = bb[i]
		bb[i] = tem
	}

	aa = string(bb)
	return aa
}

// 写一个程序，对包含中问的字符串进行逆序。

func zwreverse() string {
	var aa = "hello 世界"
	var bb []rune
	bb = []rune(aa)

	for i := 0; i < len(bb)/2; i++ {
		tem := bb[len(bb)-i-1]
		bb[len(bb)-i-1] = bb[i]
		bb[i] = tem
	}

	aa = string(bb)
	return aa
}

//写一个程序，判断一个字符串是否是回文  (正序和倒序一样)
func testHuiwen() bool {
	var aa = "上海自来水来自海上"
	var bb []rune
	bb = []rune(aa)

	for i := 0; i < len(bb)/2; i++ {
		tem := bb[len(bb)-i-1]
		bb[len(bb)-i-1] = bb[i]
		bb[i] = tem
	}

	cc := string(bb)
	return cc == aa

}

func main() {
	// aa := ewreverse()
	// aa := zwreverse()
	aa := testHuiwen()
	fmt.Printf("aa=%t", aa)

}
