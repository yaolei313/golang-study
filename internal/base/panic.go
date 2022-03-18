package base

import (
	"fmt"
	"time"
)

func PanicStudy()  {

	fmt.Println("hi")

	go func() {
		time.Sleep(time.Second * 5)
		panic("panic 123")
	}()

	for {
		time.Sleep(time.Second)
	}
}