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
	pool model.Pool
}

func newPoolService(pool model.Pool) *PoolService {
	return &PoolService{pool}
}

func executeTask(worker *model.Worker, task func()) {
	task()
	logger.Infof(workerReceivedTaskMsg, worker.ID)
}

func (p *PoolService) Start(worker *model.Worker) {
	go func() {
		logger.Infof(workerStartedMsg, worker.ID)
		for {
			select {
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
		TaskCh: make(chan func()),
	}

	poolService := newPoolService(*pool)

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
	for _, worker := range p.pool.Workers {
		if worker.Active {
			logger.Infof("Task submitted to worker %d", worker.ID)
			worker.Task <- task
			return
		}
	}
	p.pool.TaskCh <- task
}

func (p *PoolService) Close() {
	logger.Info("Closing pool")
	close(p.pool.TaskCh)
}
