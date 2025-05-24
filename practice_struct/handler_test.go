package practice_struct

import (
	"fmt"
	"testing"
)

func Test_test(t *testing.T) {
	tests := [...]struct {
		name string
		age  int
	}{
		{"Alice", 20},
		{"Bob", 30},
	}
	for _, tt := range tests {

		fmt.Println(tt)
	}
}
