package main

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func main() {
	logger.Info("tinker benchmark test")

	// Definir a rota de health check
	http.HandleFunc("/health", healthCheck)

	// Iniciar o servidor na porta 8080
	port := ":8080"
	logger.Infof("Servidor iniciado na porta %s", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		logger.Fatal("Erro ao iniciar o servidor: ", err)
	}
}

// Função para manipular a rota de health check
func healthCheck(w http.ResponseWriter, r *http.Request) {
	logger.Info("Rota de health check acessada")
	// Responder com status OK
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Servidor está saudável")
}
