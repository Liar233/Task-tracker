package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/Liar233/Task-tracker/internal/app/model"
)

type TaskPostgresRepository struct {
	db *sql.DB
}

func (tpr *TaskPostgresRepository) Get(name string) (*model.Task, error) {

	task := &model.Task{}

	err := tpr.db.QueryRow(
		"SELECT name, owner, status FROM tasks WHERE name = $1 AND status != $2;",
		name,
		model.DELETED,
	).
		Scan(&task.Name, &task.User, &task.Status)

	if err == sql.ErrNoRows {

		return nil, NotFountDBError
	}

	if err != nil {

		return nil, err
	}

	return task, nil
}

func (tpr *TaskPostgresRepository) Create(task *model.Task) (*model.Task, error) {

	var count int

	err := tpr.db.QueryRow(
		"SELECT count(*) FROM tasks WHERE name = $1 AND status != $2;",
		task.Name,
		model.DELETED,
	).Scan(&count)

	if err != nil {

		return nil, err
	}

	if count > 0 {

		return nil, AlreadyExistsDBError
	}

	_, err = tpr.db.Exec(
		"INSERT INTO tasks(name, owner, status) VALUES ($1, $2, $3)",
		task.Name,
		task.User,
		task.Status,
	)

	if err != nil {

		return nil, err
	}

	return task, nil
}

func (tpr *TaskPostgresRepository) Update(task *model.Task) (*model.Task, error) {

	tx, err := tpr.db.Begin()

	if err != nil {

		return nil, err
	}

	var count int

	txErr := tx.QueryRow(
		"SELECT count(*) FROM tasks WHERE name = $1 AND status != $2;",
		task.Name,
		model.DELETED,
	).Scan(&count)

	if err != nil {

		return nil, err
	}

	if count == 0 {

		return nil, NotFountDBError
	}

	if txErr != nil {

		err = tx.Rollback()

		if txErr = tx.Rollback(); txErr != nil {

			return nil, txErr
		}
	}

	_, err = tx.Exec(
		"UPDATE tasks SET owner = $1, status=$2 WHERE name=$3;",
		task.User,
		task.Status,
		task.Name,
	)

	if txErr = tx.Commit(); txErr != nil {

		return nil, txErr
	}

	return task, nil
}

func (tpr *TaskPostgresRepository) Delete(name string) error {

	_, err := tpr.db.Exec(
		"UPDATE tasks SET status=$1 WHERE name=$2 AND status != $3;",
		model.DELETED,
		name,
		model.DELETED,
	)

	if err != nil {

		return err
	}

	return err
}

func (tpr *TaskPostgresRepository) DeleteAll() error {

	_, err := tpr.db.Exec("DELETE FROM tasks;")

	return err
}

func (tpr *TaskPostgresRepository) GetList(userName string) ([]*model.Task, error) {

	rows, err := tpr.db.Query(
		"SELECT name, owner, status FROM tasks WHERE owner = $1 AND status != $2;",
		userName,
		model.DELETED,
	)

	if err != nil {

		return nil, err
	}

	tasks := make([]*model.Task, 0)

	var scanErr error

	for rows.Next() {

		task := &model.Task{}

		if scanErr = rows.Scan(&task.Name, &task.User, &task.Status); scanErr != nil {

			break
		}

		tasks = append(tasks, task)
	}

	if closeErr := rows.Close(); err != nil {

		return nil, closeErr
	}

	if scanErr != nil {

		return nil, scanErr
	}

	if rows.Err() != nil {

		return nil, rows.Err()
	}

	return tasks, nil
}

func (tpr *TaskPostgresRepository) Close() error {

	return tpr.db.Close()
}

func NewTaskPostgresRepository(host, dbname, username, password string, port uint64) (*TaskPostgresRepository, error) {

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", username, password, host, port, dbname)

	db, err := sql.Open("postgres", dsn)

	if err != nil {

		return nil, err
	}

	return &TaskPostgresRepository{
		db: db,
	}, nil
}
