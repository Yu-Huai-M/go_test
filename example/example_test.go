package example

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := add(1, 2)
	if result != 3 {
		t.Errorf("add(1, 2)=%d; want %d", result, 3)
	}
	t.Logf("test add success")

}

func TestSub(t *testing.T) {
	result := sub(1, 2)
	if result != -1 {
		t.Errorf("sub(1, 2)=%d; want %d", result, -1)
	}
	t.Logf("test sub success")
}
