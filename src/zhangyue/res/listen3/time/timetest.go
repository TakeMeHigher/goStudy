package main

import (
	"fmt"
	"time"
)

func timeTest() {
	now := time.Now()
	fmt.Println(now)

	year := now.Year()
	month := now.Month()
	day := now.Day()

	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	// 时间戳
	fmt.Println("当前时间戳", now.Unix())
}

// 将时间戳转换为时间
func timestampToTime(timestamp int64) {
	fmt.Printf("转换的时间为:%v\n", time.Unix(timestamp, 0))
}

// 讲时间转换为时间戳

func timeToStamp(str_time string) {
	t, _ := time.Parse("2006-01-02 15:04:05", str_time)
	fmt.Println("当前时间字符串转换为时间戳是", t.Unix())
	fmt.Println("当前时间戳是", time.Now().Unix())
}


func task(){
	fmt.Print("do task\n")
}

// 定时器 
func timeTick(){
	tick := time.Tick(1*time.Second)
	for i:= range tick{
		fmt.Println(i)
		task()
	}
}

func main() {
	// timeTest()

	// stamp := time.Now().Unix()
	// timestampToTime(stamp)

	// str_time := time.Now().Format("2006-01-02 15:04:05")  // 格式化时间固定写成2006-01-02 15:04:05
	// fmt.Println(str_time, "str_time")
	// timeToStamp(str_time)
	
	timeTick()
}
