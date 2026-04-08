package service

import (
	"errors"
	"time"

	"github.com/rozoomcool/task-cli/internal/model"
	"github.com/rozoomcool/task-cli/internal/repository"
)

type TaskService interface {
	AddTask(string) (int, error)
	ListTasks(model.TaskStatus) ([]*model.Task, error)
	DeleteTask(int) error
	UpdatedTask(int, string, model.TaskStatus) error
}

type taskService struct {
	Repo repository.Repository[*model.Task]
}

var (
	ts TaskService
)

func NewTaskService() TaskService {
	if ts == nil {
		repo := repository.GetTaskRepository()
		ts = &taskService{Repo: repo}
	}
	return ts
}

// AddTask implements TaskService.
func (t *taskService) AddTask(description string) (int, error) {
	newTask := model.Task{
		Description: description,
		Status:      model.TaskStatusToDo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return t.Repo.Add(&newTask)
}

// DeleteTask implements TaskService.
func (t *taskService) DeleteTask(id int) error {
	return t.Repo.Delete(id)
}

// ListTasks implements TaskService.
func (t *taskService) ListTasks(status model.TaskStatus) ([]*model.Task, error) {
	tasks, err := t.Repo.List()
	if err != nil {
		return nil, err
	}

	if status == "todo" || status == "in-progress" || status == "done" {
		response := []*model.Task{}
		for _, v := range tasks {
			if v.Status == status {
				response = append(response, v)
			}
		}
		return response, nil
	}

	return tasks, nil

}

// UpdatedTask implements TaskService.
func (t *taskService) UpdatedTask(id int, description string, status model.TaskStatus) error {
	tasks, err := t.Repo.List()
	if err != nil {
		return err
	}

	var entity *model.Task
	for _, v := range tasks {
		if v.GetID() == id {
			entity = v
			break
		}
	}
	if entity == nil {
		return errors.New("Task not found")
	}

	if description != "" {
		entity.Description = description
	}

	if status != "" {
		entity.Status = status
	}

	return t.Repo.Update(id, entity)
}
