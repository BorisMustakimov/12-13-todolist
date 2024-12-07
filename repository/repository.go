package repository

import (
	"github.com/BorisMustakimov/12-13-todolist/task"
	"github.com/jmoiron/sqlx"
)

type TaskRepository interface {
	Create(task *task.Task) (int64, error)
}

type TaskRepo struct {
	db *sqlx.DB
}

func NewTaskRepo(db *sqlx.DB) TaskRepository {
	return &TaskRepo{db: db}
}

func (r *TaskRepo) Create(task *task.Task) (int64, error) {
	res, err := r.db.Exec(
		`INSERT INTO scheduler (date, title, comment, repeat) VALUES (?,?,?,?)`,
		task.Date, task.Title, task.Comment, task.Repeat,
	)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
