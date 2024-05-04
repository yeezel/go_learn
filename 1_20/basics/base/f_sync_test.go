package basics

import "sync"

type Task struct {
	i int
}
type Pool struct {
	Mu    sync.Mutex
	rmu   sync.RWMutex
	Tasks []*Task
}

func Worker(pool *Pool) {
	for len(pool.Tasks) > 0 {
		// pool.rmu.RLock() //只读锁
		// defer pool.rmu.RUnlock()
		pool.Mu.Lock() //互斥锁
		// begin critical section:
		task := pool.Tasks[0]       // take the first task
		pool.Tasks = pool.Tasks[1:] // update the pool of tasks
		// end critical section
		pool.Mu.Unlock()
		process(task)
	}
}

func process(task *Task) {
	println("task running")
}
func SyncDemo() {
	ts := make([]*Task, 5)
	for i := 0; i < 5; i++ {
		ts[i] = &Task{i}
	}
	tasks := Pool{Tasks: ts}
	Worker(&tasks)
}
