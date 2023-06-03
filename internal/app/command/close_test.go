package command

import (
	"testing"

	"github.com/Liar233/Task-tracker/internal/app/model"
	"github.com/Liar233/Task-tracker/internal/app/storage"
)

func TestCloseCommand_Exec(t *testing.T) {

	testTaskRep := storage.NewTaskMemoryRepository()

	_, _ = testTaskRep.Create(&model.Task{
		User:   "user2",
		Name:   "task1",
		Status: model.CLOSED,
	})

	cmd := NewCloseCommand(testTaskRep)

	reqDto := RequestDto{
		User: "user1",
		Cmd:  CloseTaskCmd,
		Arg:  "task1",
	}

	if resDto := cmd.Exec(reqDto); resDto.Error() == nil {

		t.Error("failed close non-existent task")
	}

	if resDto := cmd.Exec(reqDto); resDto.Error() == nil {

		t.Error("failed close someone else's task")
	}

	reqDto = RequestDto{
		User: "user2",
		Cmd:  CloseTaskCmd,
		Arg:  "task1",
	}

	if resDto := cmd.Exec(reqDto); resDto.Error() != nil {

		t.Error("failed close someone else's task")
	}
}
