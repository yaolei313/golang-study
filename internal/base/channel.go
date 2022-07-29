package base

import (
	"log"
	"time"
)

func kickOffBoll(playerName string, ch chan string) {
	for {
		fromPlayer, ok := <-ch
		if ok {
			log.Println("receive ball", fromPlayer)
			time.Sleep(time.Second)
			SafeSend(ch, playerName)
		}
	}
}

func StudyChannel() {
	//带缓冲的channel
	ch := make(chan string, 3)
	go kickOffBoll("张三", ch)
	go kickOffBoll("李四", ch)
	go kickOffBoll("王五", ch)
	ch <- "裁判"

	time.Sleep(5 * time.Second)
	SafeClose(ch)
	SafeClose(ch)
	log.Println("close channel")

	ch2 := make(chan string, 2)
	trySend := func(v string) {
		select {
		case ch2 <- v:
		default:
			log.Println("send fail")
		}
	}

	tryReceive := func() string {
		select {
		case m, ok := <-ch2:
			if ok {
				return m
			}
			return "-xx-"
		default:
			return "-xxx-"
		}
	}

	log.Println(tryReceive())
	trySend("hello")
	trySend("world")
	trySend("message1-fail")
	close(ch2)
	log.Println(tryReceive())
	log.Println(tryReceive())
	log.Println(tryReceive())
}

func SafeClose(ch chan string) (justClosed bool) {
	defer func() {
		if err := recover(); err != nil {
			// The return result can be altered
			// in a defer function call.
			justClosed = false
		}
	}()

	// assume ch != nil here.
	close(ch)   // panic if ch is closed
	return true // <=> justClosed = true; return
}

func SafeSend(ch chan string, value string) (closed bool) {
	defer func() {
		if err := recover(); err != nil {
			closed = true
		}
	}()

	ch <- value  // panic if ch is closed
	return false // <=> closed = false; return
}
