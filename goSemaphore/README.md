# go-semaphore
go-semaphore is a Go package that implements the semaphore pattern than can be used to manage the number of Goroutine
to run concurrently.

### installation
```shell
go get github.com/michaelwp/goSemaphore
```

### basic of use
```go
package main

import (
	"fmt"
	"github.com/michaelwp/goSemaphore"
	"sync"
	"time"
)

func main() {
	sem := gosemaphore.SemaphoreNew(5) // <-- set up the maximum allowed number of goroutine to run concurrently.
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			sem.Acquire()       // <-- to acquire a seat in buffered channel. This has to be at the beginning of the process.
			defer sem.Release() // <-- to release the seat. Strongly suggested to use defer to have a guarantee that this function will be executed when the process is finishes.

			longRunningProcess(i)
		}(i)
	}

	wg.Wait()
}

func longRunningProcess(taskID int) {
	fmt.Println(time.Now().Format("15:04:05"), "Running task with ID", taskID)
	time.Sleep(2 * time.Second)
}
```
