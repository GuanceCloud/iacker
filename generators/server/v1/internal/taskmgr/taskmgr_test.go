package taskmgr

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTaskManager(t *testing.T) {
	taskManagers := []TaskManager{
		NewGoroutineTaskManager(),
	}
	for _, mgr := range taskManagers {
		modified := false

		wg := sync.WaitGroup{}

		// Test cancel
		wg.Add(1)
		err := mgr.Spawn(context.Background(), &Task{
			Run: func(ctx context.Context) error {
				<-time.After(10 * time.Second)
				modified = true
				return nil
			},
			Callback: func(task *Task) {
				assert.NoError(t, task.Error())
				assert.Equal(t, TaskStatusCanceled, task.Status())
				assert.False(t, modified)
				wg.Done()
			},
			Id: "canceled",
		})
		<-time.After(100 * time.Millisecond)
		assert.NoError(t, mgr.Cancel("canceled"))
		assert.NoError(t, err)

		// Test Spawn
		wg.Add(1)
		err = mgr.Spawn(context.Background(), &Task{
			Run: func(ctx context.Context) error {
				modified = true
				return nil
			},
			Callback: func(task *Task) {
				assert.NoError(t, task.Error())
				assert.Equal(t, TaskStatusSuccess, task.Status())
				assert.True(t, modified)
				wg.Done()
			},
			Id: "spawn",
		})
		assert.NoError(t, err)

		// TestError
		wg.Add(1)
		err = mgr.Spawn(context.Background(), &Task{
			Run: func(ctx context.Context) error {
				return fmt.Errorf("test error")
			},
			Callback: func(task *Task) {
				assert.Error(t, task.Error())
				assert.Equal(t, TaskStatusError, task.Status())
				wg.Done()
			},
			Id: "error",
		})
		assert.NoError(t, err)

		wg.Wait()
	}
}
