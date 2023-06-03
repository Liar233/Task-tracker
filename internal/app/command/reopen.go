package command

import (
	"github.com/Liar233/Task-tracker/internal/app/model"
	"github.com/Liar233/Task-tracker/internal/app/storage"
)

type ReopenCommand struct {
	taskRep storage.TaskRepositoryInterface
}

func (rc *ReopenCommand) Exec(dto RequestDto) ResultDtoInterface {

	resDto := &ResultDto{}

	task, err := rc.taskRep.Get(dto.Arg)

	if err != nil {

		resDto.err = err

		return resDto
	}

	if task.User != dto.User {

		resDto.err = AccessDeniedError

		return resDto
	}

	task.Status = model.CREATED

	if _, err = rc.taskRep.Update(task); err != nil {

		resDto.err = err

		return resDto
	}

	return resDto
}

func NewReopenCommand(taskRep storage.TaskRepositoryInterface) *ReopenCommand {

	return &ReopenCommand{taskRep: taskRep}
}
