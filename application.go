package main

import (
	"fmt"
	"github.com/yaolei313/golang-study/internal"
	"rsc.io/quote"
)

func main() {
	func1 := internal.Fibonacci()
	fmt.Print(func1())
	fmt.Print(func1())
	fmt.Print(func1())
	fmt.Print(func1())
	fmt.Print(func1())
	fmt.Println()
	fmt.Println(quote.Hello())
}
