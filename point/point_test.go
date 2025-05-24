package point

import (
	"testing"
)

func Test_modity1(t *testing.T) {
	a := 10
	Modity1(a)
	if a != 10 {
		t.Errorf("modity1(a)=%d; want %d", a, 10)
	}
	t.Logf("test modify1 success")
}

func Test_modity2(t *testing.T) {
	a := 10
	Modity2(&a)
	if a != 100 {
		t.Errorf("modity2(a)=%d; want %d", a, 100)
	}
	t.Logf("test modify2 success")
}
