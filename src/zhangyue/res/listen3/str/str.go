package main

import "fmt"

func main() {
	var aa = "hello 世界"

	fmt.Printf("aa[0]=%c  len(aa)=%d\n", aa[0], len(aa)) // aa[0]=h  len(aa)=12  字节的长度 一个中文占3个字节

	for index, str := range aa {
		fmt.Printf("aa[%d]=%c\n", index, str)
	}
	/*
		aa[0]=h
		aa[1]=e
		aa[2]=l
		aa[3]=l
		aa[4]=o
		aa[5]=
		aa[6]=世
		aa[9]=界
	*/

	// 修改字符串

	var bb []byte
	bb = []byte(aa)
	fmt.Println(bb) // [104 101 108 108 111 32 228 184 150 231 149 140]

	bb[0] = 'H'
	aa = string(bb)
	fmt.Println(aa)                   // Hello 世界
	fmt.Printf("aa的长度:%d\n", len(aa)) // 12

	var cc []rune

	cc = []rune(aa)
	fmt.Println(cc)                   // fmt.Printf("aa的长度:%d", len(aa))
	fmt.Printf("cc的长度:%d\n", len(cc)) // 8
	fmt.Printf("aa的长度:%d\n", len(aa)) // 12
	fmt.Printf("aa[7] = %c\n", aa[7]) // aa[7] = ¸
	fmt.Printf("cc[7] = %c\n", cc[7]) // cc[7] = 界
}
