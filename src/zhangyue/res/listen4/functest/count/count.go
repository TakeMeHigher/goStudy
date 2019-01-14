package main

import "fmt"

// 求一个字符串中英文个数  空格个数 数字个数 其他个数
// 字符串用双引号 字符 用单引号
func counttest(str string) (charCount, spaceCount, numCount, otherCount int) {
	utf8_str := []rune(str)
	for i := 0; i < len(utf8_str); i++ {
		if (utf8_str[i] >= 'a' && utf8_str[i] <= 'z') || (utf8_str[i] >= 'A' && utf8_str[i] <= 'Z') {
			charCount++
			continue
		} else if utf8_str[i] == ' ' {
			spaceCount++
			continue
		} else if utf8_str[i] >= '0' && utf8_str[i] <= '9' {
			numCount++
			continue
		} else {
			otherCount++
			continue
		}
	}
	return
}

func main() {
	str := "yunnan  is beautiful 云南欢迎你 123"
	charCount, spaceCount, numCount, otherCount := counttest(str)
	fmt.Printf("charCount = %d \n spaceCount=%d \n numCount=%d \n otherCount=%d \n", charCount, spaceCount, numCount, otherCount)
}
