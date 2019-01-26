package main

import "fmt"

func main() {
	var sa = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		sa = append(sa, fmt.Sprintf("%v", i))
	}
	fmt.Println(sa)
	fmt.Printf("sa长度= %d sa 容量=%d \n", len(sa), cap(sa))
}
