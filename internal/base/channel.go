package base

import "fmt"

type Task struct {
	Name string
}

func createTasks() []Task {
	return nil
}

func main() {
	//带缓冲的channel
	ch := make(chan Task, 3)

	//启动固定数量的worker
	for i := 0; i < 10; i++ {
		go worker(ch)
	}

	//发送任务给worker
	tasks := createTasks()

	for _, task := range tasks {
		ch <- task
	}

	ch2 := make(chan int)
	<-ch2
}

func worker(ch chan Task) {
	for {
		//接受任务
		task := <-ch
		fmt.Println("handle task %v", task.Name)
	}
}
