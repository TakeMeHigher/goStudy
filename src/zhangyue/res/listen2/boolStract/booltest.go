package main

import "fmt"

func main(){
	var a bool = true

	var b bool  //bool类型不给定值 默认false

	fmt.Println(a,b)

    // == 和 !=
	fmt.Println(a == true, b== false, a != true)

	// 取反 ！

	fmt.Printf("a取反 %t", !a)

	/* 
	&&(都相同是才为true) 
	|| (有一个为true则结果就为true)  
   */

   if a == true && b == true{
	   fmt.Println("right")
   }else{
	fmt.Println("no right")
   }

   
   if a == true && b == true{
		fmt.Println("right")
   }else{
		fmt.Println("no right")
	}
   
	// 格式化输出 用Printf  %t

	fmt.Printf("a的值为:%t", a)
}