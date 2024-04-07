package concurrency

import (
	"github.com/sirupsen/logrus"

	"github.com/guisteink/tinker/domain/model"
)

var logger = logrus.New()

const (
	workerStartedMsg      = "Worker %d started"
	workerReceivedTaskMsg = "Worker %d received task from pool"
)

type PoolService struct {
	pool *model.Pool
}

func newPoolService(pool *model.Pool) *PoolService { // Modifique para aceitar um ponteiro para model.Pool
	return &PoolService{pool: pool}
}

func executeTask(worker *model.Worker, task func()) {
	task()
	worker.SetActive(true) // Marca o worker como ativo novamente após a tarefa
	logger.Infof(workerReceivedTaskMsg, worker.ID)
}

func (p *PoolService) Start(worker *model.Worker) {
	go func() { // start go-routine
		logger.Infof(workerStartedMsg, worker.ID)
		for { // start infinite loop, its when worker waits tasks to execute
			select { // wait tasks between two possible channels worker.Task or p.pool.TaskCh
			case task := <-worker.Task:
				executeTask(worker, task)
			case task := <-p.pool.TaskCh:
				executeTask(worker, task)
			}
		}
	}()
}

func Create(numWorkers int) *PoolService {
	logger.Info("Creating a new pool")

	pool := &model.Pool{
		TaskCh:  make(chan func()),
		Workers: make([]*model.Worker, 0),
	}

	poolService := newPoolService(pool)

	for i := 0; i < numWorkers; i++ {
		worker := &model.Worker{
			ID:     i,
			Task:   make(chan func()),
			Active: true,
		}

		pool.Workers = append(pool.Workers, worker)
		poolService.Start(worker)
	}

	logger.Infof("Pool created with %d workers", numWorkers)
	return poolService
}

func (p *PoolService) Submit(task func()) {
	for _, worker := range p.pool.Workers { // passa por todos os workers da pool
		if worker.IsActive() { // no primeiro worker ativo
			worker.Task <- task
			worker.SetActive(false) // Marca o worker como inativo
			return
		}
	}
	p.pool.TaskCh <- task // caso nenhum worker ativo seja encontrado, envia a task para um canal compartilhado
}

func (p *PoolService) Close() {
	logger.Info("Closing pool")
	close(p.pool.TaskCh)
}
