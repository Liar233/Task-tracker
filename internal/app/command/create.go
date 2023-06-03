package command

import (
	"github.com/Liar233/Task-tracker/internal/app/model"
	"github.com/Liar233/Task-tracker/internal/app/storage"
)

type CreateCommand struct {
	taskRep storage.TaskRepositoryInterface
}

func (crc *CreateCommand) Exec(dto RequestDto) ResultDtoInterface {

	task := &model.Task{
		User:   dto.User,
		Name:   dto.Arg,
		Status: model.CREATED,
	}

	_, err := crc.taskRep.Create(task)

	return &ResultDto{err: err}
}

func NewCreateCommand(taskRep storage.TaskRepositoryInterface) *CreateCommand {

	return &CreateCommand{taskRep: taskRep}
}
