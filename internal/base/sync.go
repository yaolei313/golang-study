package base

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func SayHello(message string, count int) {
	for i := 0; i < count; i++ {
		fmt.Println(message)
		d := time.Second * time.Duration(rand.Intn(10)) / 2
		time.Sleep(d)
	}
	wg.Done()
}

func StudySync() {
	wg.Add(2)
	go SayHello("hi, moon studio", 3)
	go SayHello("hi, li bai.", 2)
	wg.Wait()
}
