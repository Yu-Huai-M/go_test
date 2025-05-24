package practice_map

import "fmt"

func Main() {
	map1 := make(map[int]string, 10)
	map1[1] = "maohonglin"
	map1[2] = "gaoyilin"
	fmt.Println(map1)
	fmt.Println(map1[1])

	for k, v := range map1 {
		fmt.Println(k, v)
	}
}
