package taskmgr

import (
	"context"
	"sync"
)

type GoroutineTaskManager struct {
	taskMap sync.Map
}

// NewGoroutineTaskManager creates a new TaskManager
func NewGoroutineTaskManager() TaskManager {
	return &GoroutineTaskManager{
		taskMap: sync.Map{},
	}
}

// Spawn will spawn task in a new goroutine
func (taskMgr *GoroutineTaskManager) Spawn(ctx context.Context, task *Task) error {
	ctx, cancel := context.WithCancel(ctx)
	task.ctx = ctx
	task.cancel = cancel
	taskMgr.taskMap.Store(task.Id, task)
	go func() {
		defer func() {
			taskMgr.taskMap.Delete(task.Id)
			task.Callback(task)
		}()
		task.setStatus(TaskStatusRunning)
		if err := task.Run(ctx); err != nil {
			task.err = err
			task.setStatus(TaskStatusError)
			return
		}
		task.setStatus(TaskStatusSuccess)
	}()
	return nil
}

// Cancel will cancel a task execution by id
func (taskMgr *GoroutineTaskManager) Cancel(id string) error {
	if v, ok := taskMgr.taskMap.Load(id); ok {
		task := v.(*Task)
		task.cancel()
		task.setStatus(TaskStatusCanceled)
		task.Callback(task)
	}
	return nil
}
