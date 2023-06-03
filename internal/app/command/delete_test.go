package command

import (
	"testing"

	"github.com/Liar233/Task-tracker/internal/app/model"
	"github.com/Liar233/Task-tracker/internal/app/storage"
)

func TestDeleteCommand_Exec(t *testing.T) {

	testTaskRep := storage.NewTaskMemoryRepository()

	cmd := NewDeleteCommand(testTaskRep)

	task := &model.Task{
		User:   "user1",
		Name:   "task1",
		Status: model.CREATED,
	}

	_, _ = testTaskRep.Create(task)

	reqDto := RequestDto{
		User: "user1",
		Cmd:  DeleteTaskCmd,
		Arg:  "task2",
	}

	if resDto := cmd.Exec(reqDto); resDto.Error() == nil {

		t.Error("failed deleting non-existent task")
	}

	reqDto = RequestDto{
		User: "user2",
		Cmd:  DeleteTaskCmd,
		Arg:  "task1",
	}

	if resDto := cmd.Exec(reqDto); resDto.Error() == nil {

		t.Error("failed deleting someone else's task")
	}

	reqDto = RequestDto{
		User: "user1",
		Cmd:  DeleteTaskCmd,
		Arg:  "task1",
	}

	if resDto := cmd.Exec(reqDto); resDto.Error() == nil {

		t.Error("failed deleting opened task")
	}

	task.Status = model.CLOSED

	_, _ = testTaskRep.Update(task)

	if resDto := cmd.Exec(reqDto); resDto.Error() != nil {

		t.Error("failed deleting closed task")
	}
}
