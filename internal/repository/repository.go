package repository

import (
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
func (t *TaskRepository[T]) Delete(int) error {
	panic("unimplemented")
}

// List implements Repository.
func (t *TaskRepository[T]) List() ([]T, error) {
	panic("unimplemented")
}

// Update implements Repository.
func (t *TaskRepository[T]) Update(int, T) error {
	panic("unimplemented")
}
