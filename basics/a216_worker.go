package main

type Task struct {
	i int
}

// 随着任务数量增加，worker 数量也应该相应增加
func Worker(in, out chan *Task) {
	for {
		t := <-in
		process(t)
		out <- t
	}
}
func process(task *Task) {
	println("task running")
}

// 使用一个通道接受需要处理的任务，一个通道接受处理完成的任务（及其结果）
func main() {
	pending, done := make(chan *Task), make(chan *Task)
	go sendWork(pending)           // put tasks with work on the channel
	for i, N := 0, 7; i < N; i++ { // start N goroutines to do work
		go Worker(pending, done)
	}
	consumeWork(done) // continue with the processed tasks
}

func consumeWork(done chan *Task) {
	panic("unimplemented")
}

func sendWork(pending chan *Task) {
	panic("unimplemented")
}
