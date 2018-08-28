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
	Id    string
	Type1 string
}

// Exec - specific task or one work unit
func (j JobType1) Exec() {
	fmt.Printf("Performing job waiting... %+v\n", j)
	time.Sleep(3 * time.Second)
	fmt.Printf("Performing job finish %+v\n", j)
}

// JobType2 struct for Worker to process
type JobType2 struct {
	Id    string
	Type2 string
}

// Exec - specific task or one work unit
func (j JobType2) Exec() {
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
		i = i + 1
		d.Exec(&JobType1{Id: fmt.Sprintf("my-job-%d-1", i), Type1: fmt.Sprintf("message %d-1", i)}) // Exec new job requests
		d.Exec(&JobType1{Id: fmt.Sprintf("my-job-%d-2", i), Type1: fmt.Sprintf("message %d-2", i)}) // Exec new job requests
		d.Exec(&JobType1{Id: fmt.Sprintf("my-job-%d-3", i), Type1: fmt.Sprintf("message %d-3", i)}) // Exec new job requests
		d.Exec(&JobType1{Id: fmt.Sprintf("my-job-%d-4", i), Type1: fmt.Sprintf("message %d-4", i)}) // Exec new job requests
		d.Exec(&JobType1{Id: fmt.Sprintf("my-job-%d-5", i), Type1: fmt.Sprintf("message %d-5", i)}) // Exec new job requests

		d.Exec(&JobType2{Id: fmt.Sprintf("my-job-%d-1", i), Type2: fmt.Sprintf("message %d-1", i)}) // Exec new job requests
		d.Exec(&JobType2{Id: fmt.Sprintf("my-job-%d-2", i), Type2: fmt.Sprintf("message %d-2", i)}) // Exec new job requests
		d.Exec(&JobType2{Id: fmt.Sprintf("my-job-%d-3", i), Type2: fmt.Sprintf("message %d-3", i)}) // Exec new job requests
		d.Exec(&JobType2{Id: fmt.Sprintf("my-job-%d-4", i), Type2: fmt.Sprintf("message %d-4", i)}) // Exec new job requests
		d.Exec(&JobType2{Id: fmt.Sprintf("my-job-%d-5", i), Type2: fmt.Sprintf("message %d-5", i)}) // Exec new job requests
	}
}
