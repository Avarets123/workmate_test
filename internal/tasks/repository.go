package tasks

import (
	"log/slog"
	"net/http"
	"sync"
	"workmate/pkg/cerror"
)

type Repository interface {
	Create(name, desc string) (*Task, *cerror.CError)
	FindMany() []*Task
	GetOne(name string) (*Task, *cerror.CError)
	Cancel(name string) *cerror.CError
}

type repository struct {
	logger    *slog.Logger
	data      []*Task
	dataIndex map[string]int
	mu        *sync.RWMutex
}

func NewRepo(logger *slog.Logger) Repository {
	return &repository{
		logger:    logger,
		mu:        &sync.RWMutex{},
		data:      []*Task{},
		dataIndex: map[string]int{},
	}
}

func (r *repository) Create(name, desc string) (*Task, *cerror.CError) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.hasTask(name) {
		msg := "TASK_BY_PASSED_NAME_EXISTS"
		r.logger.Error(msg)
		return &Task{}, cerror.New(msg, http.StatusBadRequest)
	}

	task := NewTask(name, desc)

	r.data = append(r.data, task)
	r.dataIndex[name] = len(r.data) - 1

	r.logger.Info("New task was created!")

	return task, nil

}

func (r *repository) hasTask(name string) bool {
	_, ok := r.dataIndex[name]
	return ok
}

func (r *repository) FindMany() []*Task {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.data
}

func (r *repository) Cancel(name string) *cerror.CError {
	r.mu.Lock()
	defer r.mu.Unlock()
	if !r.hasTask(name) {
		msg := "TASK_NOT_FOUND"
		r.logger.Info(msg)
		return cerror.New(msg, http.StatusNotFound)
	}

	taskIdx := r.dataIndex[name]
	r.data[taskIdx].cancel()
	r.logger.Info("Task was canceled!")

	return nil
}

func (r *repository) GetOne(name string) (*Task, *cerror.CError) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if !r.hasTask(name) {
		return &Task{}, cerror.New("TASK_NOT_FOUND", http.StatusNotFound)
	}

	taskIdx := r.dataIndex[name]

	return r.data[taskIdx], nil

}
