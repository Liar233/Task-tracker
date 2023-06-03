package command

import (
	"testing"

	"github.com/Liar233/Task-tracker/internal/app/model"
	"github.com/Liar233/Task-tracker/internal/app/storage"
)

func TestDeleteAllCommand_Exec(t *testing.T) {

	testTaskRep := storage.NewTaskMemoryRepository()

	cmd := NewDeleteAllCommand(testTaskRep)

	task1 := &model.Task{
		User:   "user1",
		Name:   "task1",
		Status: model.CREATED,
	}

	_, _ = testTaskRep.Create(task1)

	task2 := &model.Task{
		User:   "user1",
		Name:   "task1",
		Status: model.CREATED,
	}

	_, _ = testTaskRep.Create(task2)

	reqDto := RequestDto{
		User: "user1",
		Cmd:  DeleteAllCmd,
		Arg:  "",
	}

	if resDto := cmd.Exec(reqDto); resDto.Error() != nil {

		t.Error("failed deleting all task")
	}

	if _, err := testTaskRep.Get("task1"); err == nil {

		t.Error("failed deleting all task")
	}

	if _, err := testTaskRep.Get("task2"); err == nil {

		t.Error("failed deleting all task")
	}
}
