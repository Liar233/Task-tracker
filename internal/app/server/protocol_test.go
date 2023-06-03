package server

import (
	"testing"
)

func TestBuildCommand(t *testing.T) {

	query1 := []byte(" CREATE_TASK ")
	query2 := []byte("vasya CREATE_TASK ")
	query3 := []byte(" CREATE_TASK task1")
	query4 := []byte(" __DELETE_ALL task1")
	query5 := []byte(" CREATE_TASK task1")
	query6 := []byte("vasya __DELETE_ALL")
	query7 := []byte("vasya CREATE_TASK task1")

	if _, err := BuildCommand(query1); err == nil {

		t.Errorf("fail to build %s query", query1)
	}

	if _, err := BuildCommand(query2); err == nil {

		t.Errorf("fail to build %s query", query2)
	}

	if _, err := BuildCommand(query3); err == nil {

		t.Errorf("fail to build %s query", query3)
	}

	if _, err := BuildCommand(query4); err == nil {

		t.Errorf("fail to build %s query", query4)
	}

	if _, err := BuildCommand(query5); err == nil {

		t.Errorf("fail to build %s query", query5)
	}

	cmd6, err := BuildCommand(query6)

	if err != nil {

		t.Errorf("fail to build %s query", query6)
	}

	if cmd6.User != "vasya" || cmd6.Cmd != DeleteAllCmd {

		t.Errorf("fail to build %s query", query6)
	}

	cmd7, err := BuildCommand(query7)

	if cmd7.User != "vasya" || cmd7.Cmd != CreateTaskCmd || cmd7.Arg != "task1" {

		t.Errorf("fail to build %s query", query7)
	}

	if err != nil {

		t.Errorf("fail to build %s query", query7)
	}

}
