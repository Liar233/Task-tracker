package service

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

const (
	CREATE_TASK_COMMAND = "CREATE_TASK"
	CLOSE_TASK_COMMAND  = "CLOSE_TASK"
	DELETE_TASK_COMMAND = "DELETE_TASK"
	REOPEN_TASK_COMMAND = "REOPEN_TASK"
	LIST_TASK_COMMAND   = "LIST_TASK"
	DELETE_ALL_COMMAND  = "__DELETE_ALL"
)

var (
	CREATED_RESPONSE       = []byte("CREATED")
	DELETED_RESPONSE       = []byte("DELETED")
	CLOSED_RESPONSE        = []byte("CLOSED")
	REOPENED_RESPONSE      = []byte("REOPENED")
	TASKS_RESPONSE         = []byte("TASKS")
	WRONG_FORMAT_RESPONSE  = []byte("WRONG_FORMAT")
	ACCESS_DENIED_RESPONSE = []byte("ACCESS_DENIED_FORMAT")
	ERROR_RESPONSE         = []byte("ERROR")
)

var ThreeArgsCommands = []string{
	CREATE_TASK_COMMAND,
	CLOSE_TASK_COMMAND,
	DELETE_TASK_COMMAND,
	REOPEN_TASK_COMMAND,
	LIST_TASK_COMMAND,
}

type CommandDto struct {
	user string
	cmd  string
	arg  string
}

func BuildCommand(query []byte) (*CommandDto, error) {

	str := bytes.NewBuffer(query).String()

	entries := strings.Split(str, " ")

	if len(entries) < 2 || entries[0] == "" || entries[1] == "" {

		return nil, fmt.Errorf(string(WRONG_FORMAT_RESPONSE))
	}

	command := &CommandDto{
		user: entries[0],
	}

	if entries[1] == DELETE_ALL_COMMAND {

		command.cmd = DELETE_ALL_COMMAND

		return command, nil
	}

	if entries[2] == "" {

		return nil, fmt.Errorf(string(WRONG_FORMAT_RESPONSE))
	}

	command.arg = entries[2]

	for _, cmd := range ThreeArgsCommands {

		if entries[1] == cmd {

			command.cmd = entries[1]

			return command, nil
		}
	}

	return nil, fmt.Errorf(string(WRONG_FORMAT_RESPONSE))
}

func RenderResponse(writer http.ResponseWriter, data []byte) {

	if _, err := writer.Write(data); err != nil {

		// ToDo: log this shit
		_, _ = writer.Write(ERROR_RESPONSE)
	}
}
