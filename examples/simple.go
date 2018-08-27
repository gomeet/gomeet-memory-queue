package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/gomeet/gomeet-memory-queue/models"
)

// Implements Job interface
/////////////////////////////////////////

// JobType1 struct for Worker to process
type JobType1 struct {
	Id    int
	Type1 string
}

// Perform - specific task or one work unit
func (j JobType1) Perform() {
	fmt.Printf("Performing job waiting... %+v\n", j)
	time.Sleep(3 * time.Second)
	fmt.Printf("Performing job finish %+v\n", j)
}

// JobType2 struct for Worker to process
type JobType2 struct {
	Id    int
	Type2 string
}

// Perform - specific task or one work unit
func (j JobType2) Perform() {
	fmt.Printf("Performing job waiting... %+v\n", j)
	time.Sleep(time.Second)
	fmt.Printf("Performing job finish %+v\n", j)
}

/////////////////////////////////////////

func main() {

	// NumCPU - spawn number of workers
	// 100 - Buffered asynchronous job queue
	d := models.NewDispatcher(runtime.NumCPU(), 100)
	d.Start() // Initializes workers

	i := 0
	for {
		fmt.Printf("Before Collect : %d\n", i)
		var job models.Job
		fmt.Println(i % 2)

		if (i % 2) == 0 {
			job = &JobType1{Id: i, Type1: fmt.Sprintf("type1-%d", i)}
		} else {
			job = &JobType2{Id: i, Type2: fmt.Sprintf("type2-%d", i)}
		}

		d.Collect(job) // Collects new job requests
		fmt.Printf("After Collect : %d\n", i)
		i = i + 1
	}
}
