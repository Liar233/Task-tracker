package command

import (
	"testing"

	"github.com/Liar233/Task-tracker/internal/app/model"
	"github.com/Liar233/Task-tracker/internal/app/storage"
)

func TestNewReopenCommand(t *testing.T) {

	testTaskRep := storage.NewTaskMemoryRepository()

	cmd := NewReopenCommand(testTaskRep)

	_, _ = testTaskRep.Create(&model.Task{
		User:   "user1",
		Name:   "task1",
		Status: model.CLOSED,
	})

	reqDto := RequestDto{
		User: "user2",
		Cmd:  ReopenTaskCmd,
		Arg:  "task1",
	}

	if resDto := cmd.Exec(reqDto); resDto.Error() == nil {

		t.Error("failed reopen someone else's task")
	}

	reqDto = RequestDto{
		User: "user1",
		Cmd:  ReopenTaskCmd,
		Arg:  "task1",
	}

	if resDto := cmd.Exec(reqDto); resDto.Error() != nil {

		t.Error("failed reopen closed task")
	}

	_ = testTaskRep.DeleteAll()

	if resDto := cmd.Exec(reqDto); resDto.Error() == nil {

		t.Error("failed reopen deleted task")
	}
}
