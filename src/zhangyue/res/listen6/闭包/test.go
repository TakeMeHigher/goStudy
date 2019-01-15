package main

import (
	"fmt"
	"strings"
	"time"
)

func test1() func(int) int {
	index := 10
	return func(d int) int {
		index += d
		return index
	}

}

func test2(base int) func(int) int {
	return func(a int) int {
		base += a
		return base
	}
}

func test3(suffix string) func(name string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			// return name+suffix
			return fmt.Sprintf("%s%s", name, suffix)
		}
		return name
	}
}

func test4(a int) (func(int) int, func(int) int) {

	add := func(b int) int {
		a += b
		return a
	}

	sub := func(c int) int {
		a -= c
		return a
	}

	return add, sub
}


func test5(){
	for i:=0; i<=5;i++{
		go func(){
			fmt.Println(a)
		}()
	}
	time.Sleep(time.Second)
}


func test6(){
	for i:=0; i<=5;i++{
		go func(a int){
			fmt.Println(a)
		}(i)
	}
	time.Sleep(time.Second)
}

func main() {
	// f := test1()
	// fmt.Printf("f(1) return is %d\n", f(1))  // 11
	// fmt.Printf("f(2) return is %d\n", f(2))  //13
	// fu := test1()
	// fmt.Printf("fu(1) return is %d\n", fu(5))  // 15
	// fmt.Printf("fu(2) return is %d\n", fu(6))  //21

	// f := test2(5)
	// fmt.Printf("f(1) return is %d\n", f(1)) // 6
	// fmt.Printf("f(2) return is %d\n", f(2)) //8
	// fu := test2(10)
	// fmt.Printf("fu(1) return is %d\n", fu(5)) // 15
	// fmt.Printf("fu(2) return is %d\n", fu(6)) //21

	// f3 := test3(".jpg")
	// fmt.Printf("f3('ctz') return is %s\n", f3("ctz")) // f3('ctz') return is ctz.jpg
	// fmt.Printf("f3('yunnan.jpg') return is %s\n", f3("yunnan.jpg")) //f3('yunnan.jpg') return is yunnan.jpg


	// add, sub := test4(10)

	// fmt.Println(add(5), sub(6)) //15 9
	// fmt.Println(add(3), sub(4)) // 12 8

	test5()


}
