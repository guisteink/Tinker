package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	NUM_WORKERS int
	PORT        string
)

var logger = logrus.New()

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		logger.Warn("Error loading .env file")
	}

	NUM_WORKERS = getNumWorkers()
	PORT = getAPIPort()
}

func getNumWorkers() int {
	numWorkersStr := os.Getenv("NUM_WORKERS")
	if numWorkersStr == "" {
		return 5
	}
	numWorkers, err := strconv.Atoi(numWorkersStr)
	if err != nil {
		logger.Warn("Error converting NUM_WORKERS to integer:", err)
		return 5 // Use default value if conversion fails
	}
	return numWorkers
}

func getAPIPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return ":3000" // Default port if not specified in environment variable
	}
	return ":" + port
}
