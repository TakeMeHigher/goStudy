package main

import "fmt"


func boolTest(){
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


	/*
		true false
		true true false
		a取反 falseno right
		no right
		a的值为:true
	*/

}


func intTest(){
	var a1 int8 = 124   // -127-128
	var a2 int = 885555555   
	fmt.Println(a1,a2)

	var a3 int 
	// a3 =a1  不同的数据类型不能赋值  这样写 不对 应该 把 a1转为int类型
	a3 = int(a1)
	fmt.Println(a3)

	var a4 uint8 = 100 // 无符号  0-128
	fmt.Println(a4)

	/*
		124 885555555
		124
		100
	*/
}

func floatTest(){
	var i float32 
	var j float64 = 258.336
	var n float64
	fmt.Println(i,j,n)
}

func main(){
	//boolTest()
	// intTest()
	floatTest()
	/*
		0 258.336 0
	*/
}
