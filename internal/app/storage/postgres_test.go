package storage

import (
	"testing"

	"github.com/Liar233/Task-tracker/internal/app/model"
)

func TestTaskPostgresRepository(t *testing.T) {

	taskRep, err := NewTaskPostgresRepository(
		"tracker-postgresql",
		"tracker_db",
		"tracker",
		"secret",
		5432,
	)

	if err != nil {

		t.Fatalf("failed connect db with: %s\n", err.Error())
	}

	defer func() {
		if err = taskRep.Close(); err != nil {

			t.Fatalf("failed disconnect db with: %s", err.Error())
		}
	}()

	_, err = taskRep.Get("Zii")

	if err != NotFountDBError {

		t.Logf("failed getting non-existent task: %s\n", err.Error())
	}

	task := &model.Task{
		User:   "user1",
		Name:   "task1",
		Status: model.CREATED,
	}

	task, err = taskRep.Create(task)

	if err != nil {

		t.Logf("failed creating non-existent task: %+v\n", task)
	}

	_, err = taskRep.Get("task1")

	if err != nil {

		t.Logf("failed getting existent task: %s\n", err.Error())
	}

	if taskEx, err := taskRep.Create(task); err == nil {

		t.Logf("failed creating existent task: %+v\n", taskEx)
	}

	//task.Status = model.CLOSED
	//
	//if task, err = taskRep.Update(task); err != nil {
	//
	//	t.Logf("failed updating existent task with: %+v\n", err)
	//}
	//
	//task.Name = "task2"
	//
	//if taskUp, err := taskRep.Update(task); err == nil {
	//
	//	t.Errorf("failed updating non-existent task %+v\n", taskUp)
	//}

	tasks, err := taskRep.GetList("user1")

	if err != nil {

		t.Errorf("failed get task list with %s\n", err.Error())
	}

	if len(tasks) != 1 {

		t.Errorf("failed get task list not valid count tasks %+v\n", tasks)
	}

	if err = taskRep.DeleteAll(); err != nil {

		t.Fatalf("failed deleting all tasks with: %s\n", err.Error())
	}
}
