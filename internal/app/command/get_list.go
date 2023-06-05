package command

import "github.com/Liar233/Task-tracker/internal/app/storage"

type GetListCommand struct {
	taskRep storage.TaskRepositoryInterface
}

func (glc *GetListCommand) Exec(dto RequestDto) ResultDtoInterface {

	resDto := &ResultListDto{}

	tasksList, err := glc.taskRep.GetList(dto.Arg)

	if err != nil {

		resDto.err = err

		return resDto
	}

	resDto.tasks = tasksList

	return resDto
}

func NewGetListCommand(taskRep storage.TaskRepositoryInterface) *GetListCommand {

	return &GetListCommand{taskRep: taskRep}
}
