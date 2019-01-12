package main

import "fmt"

func defer1() {
	defer fmt.Println("----------------1")
	fmt.Println("-----------------2")
	fmt.Println("------------------3")
}



func testdefer1(){
	for i:=0; i<10; i++ {
		defer fmt.Printf("this is %d\n", i)
	}
	fmt.Println("go-----------")
}


func testdefer2(){
	i:=5
	defer fmt.Printf("defer test i=%d\n",i)
	i=1000
	fmt.Println("over -----------")
}


func main() {
	// defer1()
	// testdefer1()
	testdefer2()
}
