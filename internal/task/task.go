package task

import (
	"io"
	"time"
)

type TaskStatus string

const (
	TaskStatusToDo       TaskStatus = "todo"
	TaskStatusInProgress TaskStatus = "in-progress"
	TaskStatusDone       TaskStatus = "done"
)

type Task struct {
	Id          int        `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

type TaskDB struct {
	MaxId int
	Tasks []Task
}

func NewTaskDb() *TaskDB {
	return &TaskDB{
		MaxId: 0,
		Tasks: []Task{},
	}
}

func (t *TaskDB) AddTask(task Task) int {
	t.MaxId++
	task.Id = t.MaxId
	t.Tasks = append(t.Tasks, task)
	return task.Id
}

func AddTask(description string) (int, error) {
	var newTask Task
	db, err := ReadTasks()
	if err != nil && err != io.EOF {
		return 0, err
	}
	if db == nil {
		db = NewTaskDb()
	}

	newTask = Task{
		Id:          0,
		Description: description,
		Status:      TaskStatusToDo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	newId := db.AddTask(newTask)

	err = WriteFile(db)
	if err != nil {
		return 0, err
	}

	return newId, nil
}

func ListTask(status TaskStatus) ([]Task, error) {
	db, err := ReadTasks()
	if err != nil && err != io.EOF {
		return nil, err
	}
	if db == nil {
		db = NewTaskDb()
	}
	return db.Tasks, nil
}
