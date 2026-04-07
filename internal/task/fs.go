package task

import (
	"encoding/json"
	"os"
	"path"
)

func tasksFilePath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return path.Join(cwd, "tasks.json"), nil
}

func ReadTasks() (*TaskDB, error) {
	filePath, err := tasksFilePath()
	if err != nil {
		return nil, err
	}
	file, err := os.OpenFile(filePath, os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	var taskDB *TaskDB
	err = json.NewDecoder(file).Decode(&taskDB)
	if err != nil {
		return nil, err
	}
	return taskDB, nil
}

func WriteFile(tasks *TaskDB) error {
	filePath, err := tasksFilePath()
	if err != nil {
		return err
	}
	file, err := os.OpenFile(filePath, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	err = json.NewEncoder(file).Encode(*tasks)
	if err != nil {
		return err
	}
	return nil
}
