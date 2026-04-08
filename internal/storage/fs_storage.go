package storage

import (
	"encoding/json"
	"io"
	"os"
	"path"

	"github.com/rozoomcool/task-cli/internal/model"
)

type FsStorage[T any] interface {
	GetFilePath() (string, error)
	Read(*T) error
	Write(*T) error
}

type TaskFsStorage struct {
	FileName string
}

var (
	taskFsStorage FsStorage[StorageWrapper[*model.Task]]
)

func GetTaskFsStorage() FsStorage[StorageWrapper[*model.Task]] {
	if taskFsStorage == nil {
		taskFsStorage = &TaskFsStorage{FileName: "tasks.json"}
	}

	return taskFsStorage
}

func (ts *TaskFsStorage) GetFilePath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return path.Join(cwd, ts.FileName), nil
}

func (ts *TaskFsStorage) Read(data *StorageWrapper[*model.Task]) error {
	filePath, err := ts.GetFilePath()
	if err != nil {
		return err
	}
	file, err := os.OpenFile(filePath, os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(data)
	if err != nil && err != io.EOF {
		return err
	}
	return nil
}

func (ts *TaskFsStorage) Write(data *StorageWrapper[*model.Task]) error {
	filePath, err := ts.GetFilePath()
	if err != nil {
		return err
	}
	file, err := os.OpenFile(filePath, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		return err
	}
	return nil
}
