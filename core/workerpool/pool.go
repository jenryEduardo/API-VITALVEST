package workerpool

type Task func() any

type WorkerPool struct {
	tasks chan chan Task
	queue chan Task
}

func New(workers int) *WorkerPool {
	pool := &WorkerPool{
		tasks: make(chan chan Task, workers),
		queue: make(chan Task),
	}

	for i := 0; i < workers; i++ {
		go worker(pool.tasks)
	}

	go dispatcher(pool.tasks, pool.queue)
	return pool
}

func worker(taskPool chan chan Task) {
	taskChan := make(chan Task)
	for {
		taskPool <- taskChan
		task := <-taskChan
		task()
	}
}

func dispatcher(taskPool chan chan Task, queue chan Task) {
	for task := range queue {
		taskChan := <-taskPool
		taskChan <- task
	}
}

func (wp *WorkerPool) Submit(task Task) {
	wp.queue <- task
}
