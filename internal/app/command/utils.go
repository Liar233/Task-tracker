package command

import (
	"github.com/Liar233/Task-tracker/internal/app/storage"
)

func MakeCommand(taskRep storage.TaskRepositoryInterface, dto RequestDto) (CmdInterface, error) {

	switch dto.Cmd {
	case CreateTaskCmd:

		return NewCreateCommand(taskRep), nil

	case ReopenTaskCmd:

		return NewReopenCommand(taskRep), nil

	case CloseTaskCmd:

		return NewCloseCommand(taskRep), nil

	case DeleteTaskCmd:

		return NewDeleteCommand(taskRep), nil

	case ListTaskCmd:

		return NewGetListCommand(taskRep), nil

	case DeleteAllCmd:

		return NewDeleteAllCommand(taskRep), nil

	default:

		return nil, InvalidCommandError
	}
}
