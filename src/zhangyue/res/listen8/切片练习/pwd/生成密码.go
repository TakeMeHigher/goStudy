package main

import (
	"flag"
	"fmt"
	"math/rand"
)

var (
	length  int
	charset string
)

const (
	NumStr  = "0123456789"
	CharStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	SpecStr = "+=-@#~,.[]()!%^*$"
)

func parseCommand() {
	flag.IntVar(&length, "l", 16, "-l 生成密码的长度")
	flag.StringVar(&charset, "t", "num",
		`-t 制定密码生成的字符集, 
		num:只使用数字[0-9], 
		char:只使用英文字母[a-zA-Z], 
		mix: 使用数字和字母， 
		advance:使用数字、字母以及特殊字符`)
	flag.Parse()
}

func getPwd() string {
	var pwd []byte = make([]byte, length, length)
	var sourceStr string
	if charset == "num" {
		sourceStr = NumStr
	} else if charset == "char" {
		sourceStr = CharStr
	} else if charset == "mix" {
		sourceStr = fmt.Sprintf("%s%s", NumStr, CharStr)

	} else if charset == "advance" {
		sourceStr = fmt.Sprintf("%s%s%s", NumStr, CharStr, SpecStr)
	} else {
		sourceStr = NumStr
	}

	for i := 0; i < length; i++ {
		index := rand.Intn(len(sourceStr))
		pwd[i] = sourceStr[index]
	}
	fmt.Println(string(pwd))
	return string(pwd)
}

func main() {
	parseCommand()
	pwd := getPwd()
	fmt.Println(pwd)
}
