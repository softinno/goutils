package goutils

import (
	"fmt"
    "testing"
    "time"
    "github.com/softinno/goutil"
)

// Test for pool of workers
func TestRunWaitNWorkers(t *testing.T) {
	_ = goutil.WorkerT (goutil.WorkerExample)
	_ = goutil.WorkerT (WorkerDummy)
	
	w := goutil.WorkerExample
	goutil.RunWaitNWorkers(w, "Example", 5, 3)
	goutil.RunWaitNWorkers(WorkerDummy, "Dummy", 5, 3)
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
