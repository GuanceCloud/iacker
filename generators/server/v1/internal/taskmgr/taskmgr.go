package taskmgr

import (
	"context"
)

// TaskManager is a manager object to manage the runtime of a task
// TODO: implement M:N worker
type TaskManager interface {
	// Spawn will spawn task in a new goroutine
	Spawn(ctx context.Context, task *Task) error

	// Cancel will cancel a task execution by id
	Cancel(id string) error
}
