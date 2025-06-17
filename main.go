package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"workmate/internal/tasks"
	"workmate/pkg/logger"

	"github.com/gorilla/mux"
)

func main() {

	logger := logger.GetLogger()

	taskRepo := tasks.NewRepo(logger)
	taskService := tasks.NewService(taskRepo)

	handler := mux.NewRouter()
	tasks.ApplyHandler(handler, taskService)

	address := "127.0.0.1:8080"
	server := http.Server{
		Addr:    address,
		Handler: handler,
	}

	go func() {
		fmt.Println((server.ListenAndServe()))
	}()
	logger.Info("Server was started on address: " + address)
	signs := make(chan os.Signal, 1)

	signal.Notify(signs, syscall.SIGINT, syscall.SIGTERM)

	signal := <-signs

	logger.Error("Server was stopped by signal: " + signal.String())

}
