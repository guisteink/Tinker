package model

import "sync"

type Pool struct {
	Workers []*Worker
	TaskCh  chan func() // shared tasks channel
}

type Worker struct {
	ID     int
	Task   chan func() // task channel for each worker
	Active bool
	Mutex  sync.Mutex
}

func (w *Worker) SetActive(active bool) {
	w.Mutex.Lock()
	w.Active = active
	w.Mutex.Unlock()
}

func (w *Worker) IsActive() bool {
	w.Mutex.Lock()
	defer w.Mutex.Unlock() // Using defer to ensure unlocking occurs even if a panic occurs
	return w.Active
}
