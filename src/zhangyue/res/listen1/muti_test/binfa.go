package main

import "fmt"
import "time"

func aa(){
	for i:=0; i<=10;i++{
		time.Sleep(1*time.Second)
        fmt.Println("执行", i , "次") 
	}
	fmt.Println("aa is ok")
}

func main(){
	go aa()  // 只需要加上 go就可以了
	fmt.Println("主函数执行完了")
	time.Sleep(15*time.Second)
}