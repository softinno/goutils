package goutils

import (
	"fmt"
	"sync"
	"time"
)

// Worker function type.
// param - name
// param - count number
// return - status
type WorkerT func(string, int) int

// Pool of workers. Run and wait n workers to finish.
// param - worker function type to be executed inside and do the work
// param - no of workers to start
// param - no of times to cicle in each worker 
func RunWaitNWorkers(f WorkerT, wName string, wNo int, jobsNo int) {
	var wg sync.WaitGroup
	
	for i := 0; i < wNo; i++ {
		wg.Add(1)

		// declare func without name and also call it
		go func(s int) {
			defer wg.Done()
			f(fmt.Sprintf("Worker %s %d", wName, s), jobsNo)		
		}(i)
	}

	wg.Wait()
}

// dummy worker doing some jobs (count times)
// dummy job - print the step
func WorkerExample (name string, count int) int {
	for i:=0; i< count; i++ {
		fmt.Println(fmt.Sprintf("%s %d (%d)", name, i, count))
		time.Sleep(100*time.Millisecond)
	}
	return 0;
}

