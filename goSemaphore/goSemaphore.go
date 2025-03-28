// Package gosemaphore is a Go package that implements the semaphore pattern than can be used to manage the number of Goroutine
// to run concurrently.
package gosemaphore

// Semaphore interface
type Semaphore interface {
	Acquire()
	Release()
}

type semaphore struct {
	semC chan struct{}
}

// SemaphoreNew create new semaphore instance
func SemaphoreNew(maxConcurrency int) Semaphore {
	return &semaphore{
		semC: make(chan struct{}, maxConcurrency),
	}
}

func (s *semaphore) Acquire() {
	s.semC <- struct{}{}
}

func (s *semaphore) Release() {
	<-s.semC
}
