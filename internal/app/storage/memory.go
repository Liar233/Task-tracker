package storage

import (
	"fmt"
	"sync"

	"github.com/Liar233/Task-tracker/internal/app/model"
)

type TaskMemoryRepository struct {
	tasks *sync.Map
}

func (tmr *TaskMemoryRepository) Create(task *model.Task) (*model.Task, error) {

	if _, ok := tmr.tasks.Load(task.Name); ok {

		return nil, fmt.Errorf("task %s already exists", task.Name)
	}

	tmr.tasks.Store(task.Name, task)

	return task, nil
}

func (tmr *TaskMemoryRepository) Update(task *model.Task) (*model.Task, error) {

	old, ok := tmr.tasks.Load(task.Name)

	if !ok {

		return nil, fmt.Errorf("there is no task %s", task.Name)
	}

	if old.(*model.Task).User != task.User {

		return nil, fmt.Errorf("task %s not available for user %s", task.Name, task.User)
	}

	tmr.tasks.Store(task.Name, task)

	return task, nil
}

func (tmr *TaskMemoryRepository) Delete(name, userName string) error {

	val, ok := tmr.tasks.Load(name)

	if !ok {

		return fmt.Errorf("there is no task %s", name)
	}

	if val.(*model.Task).User != userName {

		return fmt.Errorf("task %s not available for user %s", name, userName)
	}

	tmr.tasks.Delete(name)

	return nil
}

func (tmr *TaskMemoryRepository) GetList(userName string) ([]*model.Task, error) {

	res := make([]*model.Task, 0)

	tmr.tasks.Range(func(key, value any) bool {

		buf := value.(*model.Task)

		if buf.User == userName {

			res = append(res, buf)
		}

		return true
	})

	return res, nil
}

func NewTaskMemoryRepository() *TaskMemoryRepository {

	return &TaskMemoryRepository{
		tasks: new(sync.Map),
	}
}
