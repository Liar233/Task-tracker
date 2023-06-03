package command

import (
	"testing"

	"github.com/Liar233/Task-tracker/internal/app/model"
	"github.com/Liar233/Task-tracker/internal/app/storage"
)

func TestCreateCommand_Exec(t *testing.T) {

	testTaskRep := storage.NewTaskMemoryRepository()

	cmd := NewCreateCommand(testTaskRep)

	_, _ = testTaskRep.Create(&model.Task{
		User:   "user1",
		Name:   "task1",
		Status: model.CREATED,
	})

	reqDto := RequestDto{
		User: "user1",
		Cmd:  CreateTaskCmd,
		Arg:  "task1",
	}

	if resDto := cmd.Exec(reqDto); resDto.Error() == nil {

		t.Error("failed while creating existent task")
	}

	reqDto = RequestDto{
		User: "user1",
		Cmd:  CreateTaskCmd,
		Arg:  "task2",
	}

	if resDto := cmd.Exec(reqDto); resDto.Error() != nil {

		t.Error("failed create task")
	}
}
