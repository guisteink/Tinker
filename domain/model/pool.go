package model

import (
	"sync"

	"github.com/guisteink/tinker/domain/model"
)

type Pool struct {
	Workers []*model.Worker
	TaskCh  chan func()
	wg      sync.WaitGroup
}
