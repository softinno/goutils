package goutils

import (
	"fmt"
    "testing"
    "time"
    "github.com/softinno/goutils"
)

// Test for pool of workers
func TestRunWaitNWorkers(t *testing.T) {
	_ = goutils.WorkerT (goutils.WorkerExample)
	_ = goutils.WorkerT (WorkerDummy)
	
	w := goutils.WorkerExample
	goutils.RunWaitNWorkers(w, "Example", 5, 3)
	goutils.RunWaitNWorkers(WorkerDummy, "Dummy", 5, 3)
}

// Dummy worker doing some job
// Dummy job - print the step and wait for a while
func WorkerDummy (name string, count int) int {
	for i:=0; i< count; i++ {
		fmt.Println(fmt.Sprintf("%s %d (%d)", name, i, count))
		time.Sleep(100*time.Millisecond)
	}
	return 0;
}
