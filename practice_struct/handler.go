package practice_struct

import "fmt"

func test() {
	a := [...]struct {
		name string
		age  int
	}{
		{"Alice", 20},
		{"Bob", 30},
	}
	fmt.Println(a)
}
