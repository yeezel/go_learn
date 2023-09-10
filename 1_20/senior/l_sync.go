package senior

import "sync"

type Task struct {
	i int
}
type Pool struct {
	Mu    sync.Mutex
	Tasks []*Task
}

func Worker(pool *Pool) {
	for len(pool.Tasks) > 0 {
		pool.Mu.Lock()
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
