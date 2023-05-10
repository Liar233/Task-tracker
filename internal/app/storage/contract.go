package storage

import (
	"github.com/Liar233/Task-tracker/internal/app/model"
)

type TaskRepositoryInterface interface {
	Create(task *model.Task) (*model.Task, error)
	Update(task *model.Task) (*model.Task, error)
	Delete(name, userName string) error
	GetList(userName string) ([]*model.Task, error)
}
