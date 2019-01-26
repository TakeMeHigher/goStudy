package main

import "fmt"

func defer1() {
	defer fmt.Println("----------------1")
	fmt.Println("-----------------2")
	fmt.Println("------------------3")
}

/*
------------------2
-----------------3
----------------1
*/



func testdefer1(){
	for i:=0; i<10; i++ {
		defer fmt.Printf("this is %d\n", i)
	}
	fmt.Println("go-----------")
}
/*
go-----------
this is 9
this is 8
this is 7
this is 6
this is 5
this is 4
this is 3
this is 2
this is 1
this is 0
*/

func testdefer2(){
	i:=5
	defer fmt.Printf("defer test i=%d\n",i)
	i=1000
	fmt.Println("over -----------")
}
/*
over -----------
defer test i=5
*/

func main() {
	// defer1()
	testdefer1()
	// testdefer2()
}
