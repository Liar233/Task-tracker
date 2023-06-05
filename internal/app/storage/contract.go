package storage

import (
	"errors"

	"github.com/Liar233/Task-tracker/internal/app/model"
)

var (
	AlreadyExistsDBError = errors.New("task already exists")
	NotFountDBError      = errors.New("task not found")
	InvalidQueryDBError  = errors.New("invalid query")
)

type TaskRepositoryInterface interface {
	Get(name string) (*model.Task, error)
	Create(task *model.Task) (*model.Task, error)
	Update(task *model.Task) (*model.Task, error)
	Delete(name string) error
	DeleteAll() error
	GetList(userName string) ([]*model.Task, error)
	Open() error
	Close() error
}
