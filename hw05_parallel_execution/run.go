package hw05parallelexecution

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type CountParallelGo struct {
	mu sync.Mutex
	n  int32
}

type CountErr struct {
	max     int32
	mu      sync.Mutex
	current int32
}

// RunHw5 starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func RunHw5(tasks []Task, n, m int) error {
	if len(tasks) == 0 {
		return errors.New("no tasks to run")
	}
	countErr := CountErr{
		int32(m),
		sync.Mutex{},
		0,
	}
	countParallelGo := CountParallelGo{
		sync.Mutex{},
		int32(n),
	}

	taskChan := make(chan Task, len(tasks))
	for _, task := range tasks {
		taskChan <- task
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg := sync.WaitGroup{}

	for {
		select {
		case <-ctx.Done():
			return ErrErrorsLimitExceeded
		case <-taskChan:
			countParallelGo.mu.Lock()
			if countParallelGo.n <= 0 {
				countParallelGo.mu.Unlock()
				fmt.Println("not finished")
				continue
			}
			countParallelGo.mu.Unlock()
			if task, ok := <-taskChan; ok {
				//was sleep
				wg.Add(1)
				countParallelGo.mu.Lock()
				countParallelGo.n++
				countParallelGo.mu.Unlock()
				countErr.mu.Lock()
				curIterationCountErrs := countErr.current
				countErr.mu.Unlock()
				go func() {
					defer wg.Done()
					if curIterationCountErrs >= int32(m) {
						cancel()
					}
					err := task()
					if err != nil {
						countErr.mu.Lock()
						countErr.current++
						countErr.mu.Unlock()
					}
					countParallelGo.mu.Lock()
					countParallelGo.n--
					countParallelGo.mu.Unlock()
				}()
			} else {
				return nil
			}
		case <-time.After(5 * time.Second):
			return errors.New("timeout")
		}

		wg.Wait()
	}
}
