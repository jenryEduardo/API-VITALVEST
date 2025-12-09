package workerpool

import (
	"context"
	"sync"
)

// Task es una función que retorna (data, error)
type Task func() (interface{}, error)

type WorkerPool struct {
	workers   int
	taskQueue chan taskWrapper
	wg        sync.WaitGroup
	ctx       context.Context
	cancel    context.CancelFunc
}

type taskWrapper struct {
	task       Task
	resultChan chan taskResult
}

type taskResult struct {
	Data interface{}
	Err  error
}

// New crea un worker pool con n workers
func New(workers int) *WorkerPool {
	ctx, cancel := context.WithCancel(context.Background())
	
	pool := &WorkerPool{
		workers:   workers,
		taskQueue: make(chan taskWrapper, workers*2),
		ctx:       ctx,
		cancel:    cancel,
	}
	
	pool.start()
	return pool
}

func (wp *WorkerPool) start() {
	for i := 0; i < wp.workers; i++ {
		wp.wg.Add(1)
		go wp.worker()
	}
}

func (wp *WorkerPool) worker() {
	defer wp.wg.Done()
	
	for {
		select {
		case <-wp.ctx.Done():
			return
		case wrapper, ok := <-wp.taskQueue:
			if !ok {
				return
			}
			
			// Ejecutar tarea y enviar resultado
			data, err := wrapper.task()
			wrapper.resultChan <- taskResult{Data: data, Err: err}
			close(wrapper.resultChan)
		}
	}
}

// Submit envía una tarea y retorna un canal para recibir el resultado
func (wp *WorkerPool) Submit(task Task) <-chan taskResult {
	resultChan := make(chan taskResult, 1)
	
	wrapper := taskWrapper{
		task:       task,
		resultChan: resultChan,
	}
	
	select {
	case wp.taskQueue <- wrapper:
		// Tarea encolada exitosamente
	case <-wp.ctx.Done():
		// Pool cerrado, enviar error
		resultChan <- taskResult{Err: context.Canceled}
		close(resultChan)
	}
	
	return resultChan
}

// Stop detiene el pool gracefully
func (wp *WorkerPool) Stop() {
	wp.cancel()
	close(wp.taskQueue)
	wp.wg.Wait()
}

// Stats retorna estadísticas del pool
func (wp *WorkerPool) Stats() map[string]interface{} {
	return map[string]interface{}{
		"workers":        wp.workers,
		"queued_tasks":   len(wp.taskQueue),
		"queue_capacity": cap(wp.taskQueue),
	}
}