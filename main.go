package main

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/guisteink/tinker/config"
	"github.com/guisteink/tinker/infraestructure/concurrency"
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

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("")
		logger.Info("Handling health check request")

		pool.Submit(func() {
			fmt.Fprintf(w, "Server is OK")
		})
	})

	logger.Infof("Starting server on port %s", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		logger.Fatalf("Error starting server: %s", err)
	}
}
