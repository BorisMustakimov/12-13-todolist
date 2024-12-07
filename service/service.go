package service

import (
	"fmt"
	"log"
	"time"

	"github.com/BorisMustakimov/12-13-todolist/nextdate"
	"github.com/BorisMustakimov/12-13-todolist/repository"
	"github.com/BorisMustakimov/12-13-todolist/task"
)

type TaskService interface {
	AddTask(task *task.Task) (int64, error)
	TaskDone(id string, now time.Time) error
}

type TaskServiceImpl struct {
	Repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskServiceImpl {
	return &TaskServiceImpl{
		Repo: repo,
	}
}

func (s *TaskServiceImpl) AddTask(task *task.Task) (int64, error) {
	now := time.Now()
	var taskDate time.Time

	if task.Date == "" || task.Date == now.Format(nextdate.DateFormat) {
		taskDate = now
		task.Date = now.Format(nextdate.DateFormat)
	} else {
		var err error
		taskDate, err = time.Parse(nextdate.DateFormat, task.Date)
		if err != nil {
			log.Printf("неправильный формат даты: %v", err)
			return 0, fmt.Errorf("неправильный формат")
		}
	}

	if taskDate.Before(now) {
		if task.Repeat == "" || task.Repeat == "d 1" {
			task.Date = now.Format(nextdate.DateFormat)
		} else {
			nextDate, err := nextdate.NextDate(now, taskDate.Format(nextdate.DateFormat), task.Repeat)
			if err != nil {
				return 0, fmt.Errorf("невозможно расчитать дату: %v", err)
			}
			task.Date = nextDate
		}
	}

	id, err := s.Repo.Create(task)
	if err != nil {
		return 0, fmt.Errorf("ошибка сохранения задачи: %v", err)
	}

	return id, nil
}
