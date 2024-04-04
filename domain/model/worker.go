package model

type Worker struct {
	ID     int
	Task   chan func()
	Active bool
}
