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

func newPoolService(pool *model.Pool) *PoolService { //
	return &PoolService{pool: pool}
}

func executeTask(worker *model.Worker, task func()) {
	task()
	worker.SetActive(true) // Mark the worker as active again after the task
	logger.Infof(workerReceivedTaskMsg, worker.ID)
}

func (p *PoolService) Start(worker *model.Worker) {
	go func() { // start go-routine
		logger.Infof(workerStartedMsg, worker.ID)
		for { // start an infinite loop, which is when each task waits to be executed
			select { // selects worker channel or shared channel execution
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
	for _, worker := range p.pool.Workers { // iterate all workers
		if worker.IsActive() { // on the first active worker
			worker.Task <- task     // assigns task execution
			worker.SetActive(false) // mark the worker as inactive
			return
		}
	}
	p.pool.TaskCh <- task // if no active worker is found, sends the task to a shared channel
}

func (p *PoolService) Close() {
	logger.Info("Closing pool")
	close(p.pool.TaskCh)
}
