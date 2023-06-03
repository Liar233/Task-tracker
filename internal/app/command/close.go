package command

import (
	"github.com/Liar233/Task-tracker/internal/app/model"
	"github.com/Liar233/Task-tracker/internal/app/storage"
)

type CloseCommand struct {
	taskRep storage.TaskRepositoryInterface
}

func (clc *CloseCommand) Exec(dto RequestDto) ResultDtoInterface {

	resDto := &ResultDto{}

	task, err := clc.taskRep.Get(dto.Arg)

	if err != nil {

		resDto.err = err

		return resDto
	}

	if task.User != dto.User {

		resDto.err = AccessDeniedError

		return resDto
	}

	task.Status = model.CLOSED

	if _, err = clc.taskRep.Update(task); err != nil {

		resDto.err = err

		return resDto
	}

	return resDto
}

func NewCloseCommand(taskRep storage.TaskRepositoryInterface) *CloseCommand {

	return &CloseCommand{taskRep: taskRep}
}
