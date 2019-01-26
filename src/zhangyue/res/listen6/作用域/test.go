package main

import "fmt"

var a int = 100  // a 是全局变量  整个程序 


func test(){
	fmt.Println(a)
	var b = 123 // b为局部变量 在整个test函数中 有用 即作用域为test函数的范围  test函数被执行完以后被销毁  在test函数外不能被调用
	if b >0{
		var c = 125 // b为局部变量 作用域范围为他所在的if语句块 if语句块执行玩以后 被销毁 在所在if语句块 之外不能调用
	}

	if d:=10; if d>5{
		fmt.Println(d)
	}else{
		fmt.Println(d)
	}   // 这种 d的作用域 在整个 if 语句中  包括 if   else if   else


	for i：=0; i<100; i++{
		fmt.Println(i)   // i 的作用域为整个for循坏的范围 for 循坏执行完以后 i被销毁  在for 循环外 不能调用i
	}
}

func main(){

}