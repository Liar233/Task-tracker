package command

import "github.com/Liar233/Task-tracker/internal/app/storage"

type DeleteCommand struct {
	taskRep storage.TaskRepositoryInterface
}

func (dc *DeleteCommand) Exec(dto RequestDto) ResultDtoInterface {

	resDto := &ResultDto{}

	task, err := dc.taskRep.Get(dto.Arg)

	if err != nil {

		resDto.err = err

		return resDto
	}

	if task.User != dto.User {

		resDto.err = AccessDeniedError

		return resDto
	}

	resDto.err = dc.taskRep.Delete(dto.Arg)

	return resDto
}

func NewDeleteCommand(taskRep storage.TaskRepositoryInterface) *DeleteCommand {

	return &DeleteCommand{taskRep: taskRep}
}
