package base

import (
	"errors"
	"fmt"
	"go/types"
	"time"
)

type NameAware interface {
	GetName() string
}
type Student struct {
	name string
}

func (s Student) GetName() string {
	return s.name
}

func (s *Student) SetName(str string) {
	s.name = str
}

func PrintName(aware NameAware) {
	fmt.Println("name is ", aware.GetName())
}

func Swap2(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func Swap1(a, b *int) {
	*a, *b = *b, *a
}

func PrintType(obj interface{}) {
	switch i := obj.(type) {
	case types.Nil:
		fmt.Printf("obj is %T\n", i)
	case int:
		fmt.Println("obj is int")
	case bool, string:
		fmt.Println("obj is bool or string")
	case Student:
		fmt.Println("obj is Student")
	case *Student:
		fmt.Println("obj is Student Pointer")
	case uintptr:
		fmt.Println("obj is point")
	default:
		fmt.Println("unknown type")
	}
}

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		return n * Factorial(n-1)
	}
	return 1
}

type Handler func(s string) int

func (h Handler) Add(param1 string, param2 int) int {
	return h(param1) + param2
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("b is zero")
	} else {
		return a / b, nil
	}
}

// defer
func sdefer1() int {
	i := 0
	// 1.defer在return之后被调用，但其参数为实时解析
	// 2.多个defer后定义的先被执行
	// 类似finally吧
	defer fmt.Println("second call", i)
	i++
	defer fmt.Println("first call")
	return i
}



func sayHello(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s)
	}
}
