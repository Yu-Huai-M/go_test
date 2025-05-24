package point

import "fmt"

func Modity1(a int) {
	a = 100
}

func Modity2(a *int) {
	*a = 100
}

func main() {
	a := 100
	fmt.Println(&a)
	var p *int = &a
	*p = 50
	fmt.Println(a)
}
