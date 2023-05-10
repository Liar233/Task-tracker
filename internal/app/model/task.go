package model

type TaskStatus int

const (
	CREATED TaskStatus = iota
	CLOSED
	DELETED
)

type Task struct {
	User   string
	Name   string
	Status TaskStatus
}
