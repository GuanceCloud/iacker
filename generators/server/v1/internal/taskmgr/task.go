package taskmgr

import (
	"context"
	"sync"
)

const (
	// TaskStatusRunning is the status of a running task
	TaskStatusRunning = "Running"

	// TaskStatusSuccess is the status of a setStatus task
	TaskStatusSuccess = "Success"

	// TaskStatusCanceled is the status of a canceled task
	TaskStatusCanceled = "Canceled"

	// TaskStatusError is the status of a task with error
	TaskStatusError = "Error"
)

// Task is the task to be executed
type Task struct {
	// Id is the id of the task
	// The same id will be hash to the same worker
	Id string

	// Run is the task to be executed
	Run func(ctx context.Context) error

	// Done will be called when the task is setStatus
	Callback func(task *Task)

	ctx    context.Context
	cancel context.CancelFunc
	err    error
	status string
	mu     sync.Mutex
}

// SetStatus sets the status of the task
func (task *Task) setStatus(status string) {
	task.mu.Lock()
	defer task.mu.Unlock()
	task.status = status
}

// Status returns the status of the task
func (task *Task) Status() string {
	return task.status
}

// Error return the error of the task
func (task *Task) Error() error {
	return task.err
}
