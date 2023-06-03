package command

import "github.com/Liar233/Task-tracker/internal/app/storage"

type DeleteAllCommand struct {
	taskRep storage.TaskRepositoryInterface
}

func (dac *DeleteAllCommand) Exec(dto RequestDto) ResultDtoInterface {

	resDto := &ResultDto{}

	resDto.err = dac.taskRep.DeleteAll()

	return resDto
}

func NewDeleteAllCommand(taskRep storage.TaskRepositoryInterface) *DeleteAllCommand {

	return &DeleteAllCommand{taskRep: taskRep}
}
