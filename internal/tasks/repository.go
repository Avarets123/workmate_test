package tasks

import (
	"log/slog"
	"net/http"
	"sync"
	"workmate/pkg/cerror"
)

type Repository interface {
	Create(name, desc string) *Task
	FindMany() []*Task
	GetOne(id int) (*Task, *cerror.CError)
	Cancel(id int) *cerror.CError
}

type repository struct {
	logger *slog.Logger
	data   []*Task
	mu     *sync.RWMutex
}

func NewRepo(logger *slog.Logger) Repository {
	return &repository{
		logger: logger,
		mu:     &sync.RWMutex{},
		data:   []*Task{},
	}
}

func (r *repository) Create(name, desc string) *Task {
	r.mu.Lock()
	defer r.mu.Unlock()

	task := NewTask(name, desc)
	task.Id = len(r.data) + 1

	r.data = append(r.data, task)

	r.logger.Info("New task was created!")

	return task

}

func (r *repository) hasTask(id int) bool {

	if id < 1 {
		return false
	}

	if len(r.data) < id {
		return false
	}

	return true
}

func (r *repository) FindMany() []*Task {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.data
}

func (r *repository) Cancel(id int) *cerror.CError {
	r.mu.Lock()
	defer r.mu.Unlock()
	if !r.hasTask(id) {
		msg := "TASK_NOT_FOUND"
		r.logger.Error(msg)
		return cerror.New(msg, http.StatusNotFound)
	}

	r.data[id-1].cancel()
	r.logger.Info("Task was canceled!")

	return nil
}

func (r *repository) GetOne(id int) (*Task, *cerror.CError) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if !r.hasTask(id) {
		return &Task{}, cerror.New("TASK_NOT_FOUND", http.StatusNotFound)
	}

	return r.data[id-1], nil

}
