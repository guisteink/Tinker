package main

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/guisteink/tinker/config"
	"github.com/guisteink/tinker/infraestructure/concurrency"
)

const (
	healthCheckRoute = "/health"
	serverStartMsg   = "Starting server on port %s"
	serverErrorMsg   = "Error starting server: %s"
)

var logger = logrus.New()

func main() {
	logger.Info("[tinker-benchmark]")

	numWorkers := config.NUM_WORKERS
	pool := initializePool(numWorkers)
	defer pool.Close()

	initializeHTTPServer(pool)
}

func initializePool(numWorkers int) *concurrency.PoolService {
	return concurrency.Create(numWorkers)
}

func initializeHTTPServer(pool *concurrency.PoolService) {
	port := config.PORT
	logger.Infof(serverStartMsg, port)

	pool.Submit(func() {
		assignRoutes()
	})

	err := http.ListenAndServe(port, nil)
	if err != nil {
		logger.Fatalf(serverErrorMsg, err)
	}
}

// assign all routes that will be available as tasks
func assignRoutes() {
	http.HandleFunc(healthCheckRoute, handleHealthCheck)
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	logger.Info("Handling health check request")
	fmt.Fprintf(w, "Server is OK")
}
