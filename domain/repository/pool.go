package repository

import (
	"github.com/guisteink/tinker/domain/model"
)

type Pool interface {
	Create(numWorkers int) error
	Start(worker *model.Worker) error
}
