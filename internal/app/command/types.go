package command

import (
	"errors"

	"github.com/Liar233/Task-tracker/internal/app/model"
)

const (
	CreateTaskCmd = "CREATE_TASK"
	CloseTaskCmd  = "CLOSE_TASK"
	DeleteTaskCmd = "DELETE_TASK"
	ReopenTaskCmd = "REOPEN_TASK"
	ListTaskCmd   = "LIST_TASK"
	DeleteAllCmd  = "__DELETE_ALL"
)

var AccessDeniedError = errors.New("AccessDenied")
var InvalidCommandError = errors.New("InvalidCommand")

type CmdInterface interface {
	Exec(dto RequestDto) ResultDtoInterface
}

type ResultDtoInterface interface {
	Error() error
	Data() []*model.Task
}

type RequestDto struct {
	User string
	Cmd  string
	Arg  string
}

type ResultDto struct {
	err error
}

func (r *ResultDto) Error() error {

	return r.err
}

func (r *ResultDto) Data() []*model.Task {

	return nil
}

type ResultListDto struct {
	ResultDto
	tasks []*model.Task
}

func (rl *ResultListDto) Data() []*model.Task {

	return rl.tasks
}
