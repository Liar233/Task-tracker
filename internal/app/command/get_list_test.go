package command

import (
	"testing"

	"github.com/Liar233/Task-tracker/internal/app/model"
	"github.com/Liar233/Task-tracker/internal/app/storage"
)

func TestGetListCommand_Exec(t *testing.T) {

	testTaskRep := storage.NewTaskMemoryRepository()

	_, _ = testTaskRep.Create(&model.Task{
		User:   "user1",
		Name:   "task1",
		Status: model.CREATED,
	})

	_, _ = testTaskRep.Create(&model.Task{
		User:   "user1",
		Name:   "task2",
		Status: model.CREATED,
	})

	_, _ = testTaskRep.Create(&model.Task{
		User:   "user1",
		Name:   "task3",
		Status: model.CREATED,
	})

	cmd := NewGetListCommand(testTaskRep)

	reqDto := RequestDto{
		User: "user1",
		Cmd:  ListTaskCmd,
		Arg:  "",
	}

	if resDto := cmd.Exec(reqDto); resDto.Error() != nil || len(resDto.Data()) != 3 {

		t.Error("failed get list tasks")
	}

}
