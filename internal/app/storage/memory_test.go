package storage

import (
	"testing"

	"github.com/Liar233/Task-tracker/internal/app/model"
)

func TestTaskMemoryRepository_Create(t *testing.T) {

	taskRep := NewTaskMemoryRepository()

	task1 := model.Task{
		User:   "vasya",
		Name:   "task1",
		Status: model.CREATED,
	}

	if t1, err := taskRep.Create(&task1); err != nil || t1 == nil {

		t.Error("failed creating task")
	}

	task2 := model.Task{
		User:   "vasya",
		Name:   "task2",
		Status: model.CREATED,
	}

	if t2, err := taskRep.Create(&task2); err != nil || t2 == nil {

		t.Error("failed creating task")
	}

	task3 := model.Task{
		User:   "vasya",
		Name:   "task1",
		Status: model.CREATED,
	}

	if t3, err := taskRep.Create(&task3); err == nil || t3 != nil {

		t.Error("failed creating existent task")
	}
}

func TestTaskMemoryRepository_Update(t *testing.T) {

	taskRep := NewTaskMemoryRepository()

	task1 := model.Task{
		User:   "vasya",
		Name:   "task1",
		Status: model.CREATED,
	}

	if t1, err := taskRep.Create(&task1); err != nil || t1 == nil {

		t.Error("failed creating task")
	}

	task2 := model.Task{
		User:   "vasya",
		Name:   "task2",
		Status: model.CREATED,
	}

	if t1, err := taskRep.Update(&task2); err == nil || t1 != nil {

		t.Error("failed trying update nonexistent task")
	}

	task3 := model.Task{
		User:   "vasya",
		Name:   "task1",
		Status: model.CLOSED,
	}

	if t1, err := taskRep.Update(&task3); err != nil || t1 == nil {

		t.Error("failed trying update existent task")
	}

	if val, ok := taskRep.tasks.Load("task1"); !ok || val.(*model.Task).Status != model.CLOSED {

		t.Error("failed")
	}
}

func TestTaskMemoryRepository_Delete(t *testing.T) {

	taskRep := NewTaskMemoryRepository()

	task1 := model.Task{
		User:   "vasya",
		Name:   "task1",
		Status: model.CREATED,
	}

	if t1, err := taskRep.Create(&task1); err != nil || t1 == nil {

		t.Error("failed creating task")
	}

	task2 := model.Task{
		User:   "vasya",
		Name:   "task2",
		Status: model.CLOSED,
	}

	if t2, err := taskRep.Create(&task2); err != nil || t2 == nil {

		t.Error("failed creating task")
	}

	if err := taskRep.Delete("task1"); err == nil {

		t.Error("failed trying delete not closed task")
	}

	if err := taskRep.Delete("task2"); err != nil {

		t.Error("failed trying delete closed task")
	}
}

func TestTaskMemoryRepository_GetList(t *testing.T) {

	taskRep := NewTaskMemoryRepository()

	task1 := model.Task{
		User:   "vasya",
		Name:   "task1",
		Status: model.CREATED,
	}

	if t1, err := taskRep.Create(&task1); err != nil || t1 == nil {

		t.Error("failed creating task")
	}

	task2 := model.Task{
		User:   "vasya",
		Name:   "task2",
		Status: model.CREATED,
	}

	if t2, err := taskRep.Create(&task2); err != nil || t2 == nil {

		t.Error("failed creating task")
	}

	task3 := model.Task{
		User:   "julia",
		Name:   "task3",
		Status: model.CREATED,
	}

	if t3, err := taskRep.Create(&task3); err != nil || t3 == nil {

		t.Error("failed creating task")
	}

	tList, err := taskRep.GetList("vasya")

	if err != nil {

		t.Error("failed get task list")
	}

	if len(tList) != 2 {

		t.Error("failed get task list invalid tasks count")
	}

	for _, task := range tList {

		if task.User != "vasya" {

			t.Error("failed get task list invalid tasks by user")

			break
		}
	}
}
