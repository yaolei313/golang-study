package main

import (
	"fmt"
	"github.com/yaolei313/golang-study/internal/base"
	"reflect"
	"rsc.io/quote"
)

func main() {
	base.TypeStudy()

	func1 := base.Fibonacci()
	fmt.Println("fibonacci:", func1(), func1(), func1(), func1(), func1(), func1())
	fmt.Println("================")
	fmt.Println(quote.Hello())

	base.TestStruct()

	fmt.Println("=======reflect.tag=========")
	base.HandleMyTag()

	val1 := string(65)
	fmt.Printf("%v %v\n", val1, reflect.TypeOf(val1))
	val2 := int32('A')
	fmt.Printf("%v %v\n", val2, reflect.TypeOf(val2))

	base.StudySync()

	fmt.Println("=====defer===========")

	defer fmt.Println("3rd message")
	defer fmt.Println("2th message")
	fmt.Println("1st message")

	fmt.Println("----------------")
	func() {
		// 延迟函数的实参是在此调用被推入延迟调用堆栈的时确定的
		for i := 0; i < 3; i++ {
			defer fmt.Println("val is :", i)
		}
		// 压栈0，1，2；打印2，1，0
	}()
	fmt.Println("----------------")
	func() {
		for i := 0; i < 3; i++ {
			defer func() {
				fmt.Println("val is :", i)
			}()
		}
		// 压栈3，3，3；打印3，3，3
	}()
	fmt.Println("================")

	ps0 := &[]string{"Go", "C"}[0]
	fmt.Println(*ps0) // Go

	ps1 := &[]string{"Go", "C"}[1]
	fmt.Println(*ps1) // C

	fmt.Println("================")

	base.StudyLog()

	base.StudyArray()
}
