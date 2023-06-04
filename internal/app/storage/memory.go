package storage

import (
	"sync"

	"github.com/Liar233/Task-tracker/internal/app/model"
)

type TaskMemoryRepository struct {
	tasks *sync.Map
}

func (tmr *TaskMemoryRepository) Get(name string) (*model.Task, error) {

	task, ok := tmr.tasks.Load(name)

	if !ok {

		return nil, NotFountDBError
	}

	return task.(*model.Task), nil
}

func (tmr *TaskMemoryRepository) Create(task *model.Task) (*model.Task, error) {

	if _, ok := tmr.tasks.Load(task.Name); ok {

		return nil, AlreadyExistsDBError
	}

	tmr.tasks.Store(task.Name, task)

	return task, nil
}

func (tmr *TaskMemoryRepository) Update(task *model.Task) (*model.Task, error) {

	_, ok := tmr.tasks.Load(task.Name)

	if !ok {

		return nil, NotFountDBError
	}

	tmr.tasks.Store(task.Name, task)

	return task, nil
}

func (tmr *TaskMemoryRepository) Delete(name string) error {

	val, ok := tmr.tasks.Load(name)

	if !ok {

		return NotFountDBError
	}

	task := val.(*model.Task)

	if task.Status != model.CLOSED {

		return InvalidQueryDBError
	}

	task.Status = model.DELETED

	tmr.tasks.Delete(name)

	return nil
}

func (tmr *TaskMemoryRepository) DeleteAll() error {

	tmr.tasks.Range(func(key interface{}, value interface{}) bool {

		tmr.tasks.Delete(key)

		return true
	})

	return nil
}

func (tmr *TaskMemoryRepository) GetList(userName string) ([]*model.Task, error) {

	res := make([]*model.Task, 0)

	tmr.tasks.Range(func(key, value any) bool {

		buf := value.(*model.Task)

		if buf.User == userName && buf.Status != model.DELETED {

			res = append(res, buf)
		}

		return true
	})

	return res, nil
}

func (tmr *TaskMemoryRepository) Close() error {

	return nil
}

func NewTaskMemoryRepository() *TaskMemoryRepository {

	return &TaskMemoryRepository{
		tasks: new(sync.Map),
	}
}
