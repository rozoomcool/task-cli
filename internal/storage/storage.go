package storage

import "github.com/rozoomcool/task-cli/internal/model"

type Identifiable interface {
	GetID() int
	SetID(int)
}

type StorageWrapper[T Identifiable] struct {
	MaxId int
	Data  []T
}

type Storage[T Identifiable] interface {
	ReadOrCreate() (*StorageWrapper[T], error)
	Commit(*StorageWrapper[T]) error
}

type TaskStorage struct {
	taskFsStorage FsStorage[StorageWrapper[*model.Task]]
}

var (
	taskStorage Storage[*model.Task]
)

func GetTaskStorage() Storage[*model.Task] {
	if taskStorage == nil {
		taskFsStorage := GetTaskFsStorage()
		taskStorage = &TaskStorage{taskFsStorage: taskFsStorage}
	}
	return taskStorage
}

// Commit implements Storage.
func (t *TaskStorage) Commit(data *StorageWrapper[*model.Task]) error {
	return t.taskFsStorage.Write(data)
}

// ReadOrCreate implements Storage.
func (t *TaskStorage) ReadOrCreate() (*StorageWrapper[*model.Task], error) {
	newTaskDb := &StorageWrapper[*model.Task]{
		MaxId: 0,
		Data:  []*model.Task{},
	}
	if err := t.taskFsStorage.Read(newTaskDb); err != nil {
		return nil, err
	}
	return newTaskDb, nil
}
