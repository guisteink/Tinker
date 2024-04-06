package model

type Pool struct {
	Workers []*Worker
	TaskCh  chan func()
}

type Worker struct {
	ID     int
	Task   chan func()
	Active bool
}
