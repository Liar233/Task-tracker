package action

import (
	"fmt"
	"net/http"

	"github.com/Liar233/Task-tracker/internal/app/command"
	"github.com/Liar233/Task-tracker/internal/app/server"
	"github.com/Liar233/Task-tracker/internal/app/storage"
)

type ExecAction struct {
	taskRep storage.TaskRepositoryInterface
}

func (ea *ExecAction) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	c := request.Context().Value("cmd")

	if c == nil {

		server.RenderResponse(writer, server.WrongFormatErrorResponse)

		return
	}

	reqDto := c.(command.RequestDto)

	cmd, err := command.MakeCommand(ea.taskRep, reqDto)

	if err != nil {

		server.RenderResponse(writer, server.WrongFormatErrorResponse)

		return
	}

	resDto := cmd.Exec(reqDto)

	ea.buildResponse(writer, reqDto, resDto)
}

func (ea *ExecAction) buildResponse(
	writer http.ResponseWriter,
	reqDto command.RequestDto,
	resDto command.ResultDtoInterface) {

	if resDto.Error() != nil {

		status := "ERROR"

		switch resDto.Error() {
		case command.AccessDeniedError:

			status = server.AccessDeniedErrorResponse

		case storage.NotFountDBError,
			storage.InvalidQueryDBError,
			storage.AlreadyExistsDBError,
			command.InvalidCommandError:

			status = server.WrongFormatErrorResponse
		}

		server.RenderResponse(writer, status)

		return
	}

	response := ""

	switch reqDto.Cmd {
	case command.CreateTaskCmd:

		response = server.CreatedResponse

	case command.ReopenTaskCmd:

		response = server.ReopenedResponse

	case command.CloseTaskCmd:

		response = server.ClosedResponse

	case command.DeleteTaskCmd,
		command.DeleteAllCmd:

		response = server.DeletedResponse

	case command.ListTaskCmd:

		buf := ""

		for i, task := range resDto.Data() {

			buf += task.Name

			if i < len(resDto.Data())-1 {

				buf += ", "
			}
		}

		response = fmt.Sprintf(server.GteListResponse, buf)
	}

	server.RenderResponse(writer, response)
}

func NewExecAction(taskRep storage.TaskRepositoryInterface) *ExecAction {

	return &ExecAction{
		taskRep: taskRep,
	}
}
