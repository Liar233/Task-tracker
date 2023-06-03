package server

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/Liar233/Task-tracker/internal/app/command"
)

const (
	CreateTaskCmd = "CREATE_TASK"
	CloseTaskCmd  = "CLOSE_TASK"
	DeleteTaskCmd = "DELETE_TASK"
	ReopenTaskCmd = "REOPEN_TASK"
	ListTaskCmd   = "LIST_TASK"
	DeleteAllCmd  = "__DELETE_ALL"
)

var (
	CreatedResponse           = "CREATED"
	ClosedResponse            = "CLOSED"
	ReopenedResponse          = "REOPENED"
	DeletedResponse           = "DELETED"
	GteListResponse           = "TASKS [%s]"
	WrongFormatErrorResponse  = "WRONG_FORMAT"
	AccessDeniedErrorResponse = "ACCESS_DENIED"
	ErrorResponse             = "ERROR"
)

var ThreeArgsCommands = []string{
	CreateTaskCmd,
	CloseTaskCmd,
	DeleteTaskCmd,
	ReopenTaskCmd,
	ListTaskCmd,
}

func BuildCommand(query []byte) (*command.RequestDto, error) {

	str := bytes.NewBuffer(query).String()

	entries := strings.Split(str, " ")

	if len(entries) < 2 || entries[0] == "" || entries[1] == "" {

		return nil, fmt.Errorf(string(WrongFormatErrorResponse))
	}

	cmdDto := &command.RequestDto{
		User: entries[0],
	}

	if entries[1] == DeleteAllCmd {

		cmdDto.Cmd = DeleteAllCmd

		return cmdDto, nil
	}

	if entries[2] == "" {

		return nil, fmt.Errorf(string(WrongFormatErrorResponse))
	}

	cmdDto.Arg = entries[2]

	for _, cmd := range ThreeArgsCommands {

		if entries[1] == cmd {

			cmdDto.Cmd = entries[1]

			return cmdDto, nil
		}
	}

	return nil, fmt.Errorf(string(WrongFormatErrorResponse))
}

func RenderResponse(writer http.ResponseWriter, response string) {

	if _, err := writer.Write([]byte(response)); err != nil {

		// ToDo: log this shit
		_, _ = writer.Write([]byte(ErrorResponse))
	}
}
