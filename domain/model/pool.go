package model

import "sync"

type Pool struct {
	Workers []*Worker
	TaskCh  chan func() // canal de tasks compartilhadas
}

type Worker struct {
	ID     int
	Task   chan func() // canal de tasks de cada worker
	Active bool
	Mutex  sync.Mutex
}

// SetActive configura o estado ativo do worker
func (w *Worker) SetActive(active bool) {
	w.Mutex.Lock()
	w.Active = active
	w.Mutex.Unlock()
}

// IsActive retorna o estado ativo do worker
func (w *Worker) IsActive() bool {
	w.Mutex.Lock()
	defer w.Mutex.Unlock() // Usando defer para garantir que o desbloqueio ocorra mesmo se ocorrer um panic
	return w.Active
}
