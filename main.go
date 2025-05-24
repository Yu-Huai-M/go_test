package main

import (
	_ "./hello"
	"fmt"
)

func main() {
	slice1 := make([]int, 5, 5)
	fmt.Println(slice1)
	fmt.Println(cap(slice1))
	slice1 = append(slice1, 10, 20, 30)
	fmt.Println(cap(slice1))
}
