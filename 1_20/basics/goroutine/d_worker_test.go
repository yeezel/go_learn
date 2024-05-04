package goroutine

import "testing"

type Task1 struct {
	i int
}

// 随着任务数量增加，worker 数量也应该相应增加
func Worker1(in, out chan *Task1) {
	for {
		t := <-in
		process1(t)
		out <- t
	}
}
func process1(task *Task1) {
	println("task running")
}

/*
使用通道的情景：
	与异步操作的结果进行交互
	分发任务
	传递数据所有权
*/
// 使用一个通道接受需要处理的任务，一个通道接受处理完成的任务（及其结果）
func TestWorkerDemo(t *testing.T) {
	pending, done := make(chan *Task1), make(chan *Task1)
	go sendWork(pending)           // put tasks with work on the channel
	for i, N := 0, 7; i < N; i++ { // start N goroutines to do work
		go Worker1(pending, done)
	}
	consumeWork(done) // continue with the processed tasks
}

func consumeWork(done chan *Task1) {
	panic("unimplemented")
}

func sendWork(pending chan *Task1) {
	panic("unimplemented")
}
