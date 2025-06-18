package tasks

import (
	"context"
	"time"
	"workmate/pkg/utils"
)

const (
	TaskSuccessStatus   = "success"
	TaskFailedStatus    = "failed"
	TaskInProcessStatus = "in_process"
)

type TaskMeta struct {
	Total   int `json:"total"`
	Process int `json:"in_process"`
	Failed  int `json:"failed"`
	Success int `json:"success"`
}

type TaskListing struct {
	Meta *TaskMeta `json:"meta"`
	Data []*Task   `json:"data"`
}

type Task struct {
	Id        int                `json:"id"`
	Name      string             `json:"name"`
	Desc      string             `json:"description"`
	Progress  int16              `json:"progress"`
	CreatedAt time.Time          `json:"createdAt"`
	WorkedAt  time.Duration      `json:"workedSeconds"`
	Status    string             `json:"status"`
	cancel    context.CancelFunc `json:"-"`
}

func (t *Task) Process() {

	randCount := utils.GetRandomCount(5)

	timeoutDur := time.Duration(randCount) * time.Minute

	ctx, _ := context.WithTimeout(context.Background(), timeoutDur)

	cancelCtx, cancel := context.WithCancel(context.Background())

	t.cancel = cancel

	go func() {
		for {
			select {
			case <-ctx.Done():
				t.Status = TaskSuccessStatus
				t.WorkedAt = time.Duration(time.Since(t.CreatedAt).Seconds())
				return
			case <-cancelCtx.Done():
				t.Status = TaskFailedStatus
				t.WorkedAt = time.Duration(time.Since(t.CreatedAt).Seconds())
				return

			default:
				time.Sleep(5 * time.Second)
				t.WorkedAt = time.Duration(time.Since(t.CreatedAt).Seconds())
				t.Progress = int16(float64(t.WorkedAt) / timeoutDur.Seconds() * 100)
			}
		}
	}()
}

func NewTask(name, desc string) *Task {

	newTask := &Task{
		Name:      name,
		Desc:      desc,
		CreatedAt: time.Now(),
		Status:    TaskInProcessStatus,

		Progress: 0,
	}

	newTask.Process()

	return newTask

}

type TaskCreateReq struct {
	Name string `json:"name"`
	Desc string `json:"description"`
}
