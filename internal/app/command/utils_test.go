package command

import (
	"reflect"
	"testing"

	"github.com/Liar233/Task-tracker/internal/app/storage"
)

func TestMakeCommand(t *testing.T) {

	testRep := &storage.TaskMemoryRepository{}

	var dto RequestDto
	var cmd CmdInterface
	var err error

	dto = RequestDto{Cmd: ListTaskCmd}

	if cmd, err = MakeCommand(testRep, dto); err != nil {

		t.Error("failed making LIST_TASK")

		if reflect.TypeOf(cmd).Elem().Name() != "GetListCommand" {

			t.Error("failed making LIST_TASK")
		}
	}

	dto = RequestDto{Cmd: CreateTaskCmd}

	if cmd, err = MakeCommand(testRep, dto); err != nil {

		t.Error("failed making CREATE_TASK")

		if reflect.TypeOf(cmd).Elem().Name() != "CreateCommand" {

			t.Error("failed making CREATE_TASK")
		}
	}

	dto = RequestDto{Cmd: CloseTaskCmd}

	if cmd, err = MakeCommand(testRep, dto); err != nil {

		t.Error("failed making CLOSE_TASK")

		if reflect.TypeOf(cmd).Elem().Name() != "CloseCommand" {

			t.Error("failed making CLOSE_TASK")
		}
	}

	dto = RequestDto{Cmd: ReopenTaskCmd}

	if cmd, err = MakeCommand(testRep, dto); err != nil {

		t.Error("failed making REOPEN_TASK")

		if reflect.TypeOf(cmd).Elem().Name() != "ReopenCommand" {

			t.Error("failed making REOPEN_TASK")
		}
	}

	dto = RequestDto{Cmd: DeleteTaskCmd}

	if cmd, err = MakeCommand(testRep, dto); err != nil {

		t.Error("failed making DELETE_TASK")

		if reflect.TypeOf(cmd).Elem().Name() != "DeleteCommand" {

			t.Error("failed making DELETE_TASK")
		}
	}

	dto = RequestDto{Cmd: DeleteAllCmd}

	if cmd, err = MakeCommand(testRep, dto); err != nil {

		t.Error("failed making DELETE_ALL_TASK")

		if reflect.TypeOf(cmd).Elem().Name() != "DeleteAllCommand" {

			t.Error("failed making DELETE_ALL_TASK")
		}
	}

	dto = RequestDto{Cmd: "NoCommand"}

	if cmd, err = MakeCommand(testRep, dto); err == nil {

		t.Error("failed making not existed command")
	}
}
