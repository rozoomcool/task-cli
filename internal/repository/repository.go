package repository

import (
	"errors"

	"github.com/rozoomcool/task-cli/internal/model"
	"github.com/rozoomcool/task-cli/internal/storage"
)

type Repository[T storage.Identifiable] interface {
	Add(T) (int, error)
	List() ([]T, error)
	Update(int, T) error
	Delete(int) error
}

type TaskRepository[T storage.Identifiable] struct {
	storage storage.Storage[T]
}

var (
	taskRepository Repository[*model.Task]
)

func GetTaskRepository() Repository[*model.Task] {
	if taskRepository == nil {
		storage := storage.GetTaskStorage()
		taskRepository = &TaskRepository[*model.Task]{
			storage: storage,
		}
	}

	return taskRepository
}

// Add implements Repository.
func (t *TaskRepository[T]) Add(entity T) (int, error) {
	db, err := t.storage.ReadOrCreate()
	if err != nil {
		return 0, err
	}

	db.MaxId++
	entity.SetID(db.MaxId)

	db.Data = append(db.Data, entity)

	err = t.storage.Commit(db)
	if err != nil {
		return 0, err
	}

	return entity.GetID(), nil
}

// Delete implements Repository.
func (t *TaskRepository[T]) Delete(id int) error {
	db, err := t.storage.ReadOrCreate()
	if err != nil {
		return err
	}

	newArray := []T{}
	founded := false
	for _, v := range db.Data {
		if v.GetID() == id {
			founded = true
			continue
		}
		newArray = append(newArray, v)
	}
	if !founded {
		return errors.New("Task not founded")
	}
	db.Data = newArray

	err = t.storage.Commit(db)
	if err != nil {
		return err
	}

	return nil
}

// List implements Repository.
func (t *TaskRepository[T]) List() ([]T, error) {
	db, err := t.storage.ReadOrCreate()
	if err != nil {
		return nil, err
	}
	return db.Data, nil
}

// Update implements Repository.
func (t *TaskRepository[T]) Update(id int, newEntity T) error {
	db, err := t.storage.ReadOrCreate()
	if err != nil {
		return err
	}

	newArray := []T{}
	founded := false
	for _, v := range db.Data {
		if v.GetID() == id {
			founded = true
			newArray = append(newArray, newEntity)
			continue
		}
		newArray = append(newArray, v)
	}
	if !founded {
		return errors.New("Task not founded")
	}
	db.Data = newArray

	err = t.storage.Commit(db)
	if err != nil {
		return err
	}

	return nil
}
