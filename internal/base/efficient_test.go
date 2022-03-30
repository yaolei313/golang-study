package base

import (
	"fmt"
	"testing"
)

type Teacher struct {
	Name string
}

func BenchmarkCompareEfficient1(b *testing.B) {

	// method 1
	t1 := new(Teacher)
	t1.Name = "teacher1"
	fmt.Sprintf("%v", t1)
}

func BenchmarkCompareEfficient2(b *testing.B) {
	// method 2
	t2 := &Teacher{
		Name: "teacher2",
	}
	fmt.Sprintf("%v", t2)
}
