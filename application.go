package main

import (
	"fmt"
	"github.com/yaolei313/golang-study/internal/base"
	"reflect"
	"rsc.io/quote"
	"unsafe"
)

func main() {
	iv := []byte{'0', '0', '0', '0', '0', '0', '0', '0'}
	fmt.Printf("%v", iv)

	var b byte = 'a'
	//Print Size, Type and Character
	fmt.Printf("Size: %d\nType: %s\nCharacter: %c\n", unsafe.Sizeof(b), reflect.TypeOf(b), b)

	func1 := base.Fibonacci()
	fmt.Print(func1())
	fmt.Print(func1())
	fmt.Print(func1())
	fmt.Print(func1())
	fmt.Print(func1())
	fmt.Println("================")
	fmt.Println(quote.Hello())

	base.TestStruct()

	fmt.Println("================")
	base.HandleMyTag()
}
