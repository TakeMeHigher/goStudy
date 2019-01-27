package main

import (
	"fmt"
)

func main() {
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	copy(veggies, fruits)
	fmt.Println(veggies, fruits)
}
