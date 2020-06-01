package main

import (
	"fmt"
	"github.com/yaolei313/golang-study/internal/base"
	"rsc.io/quote"
)

func main() {
	func1 := base.Fibonacci()
	fmt.Print(func1())
	fmt.Print(func1())
	fmt.Print(func1())
	fmt.Print(func1())
	fmt.Print(func1())
	fmt.Println()
	fmt.Println(quote.Hello())
}
