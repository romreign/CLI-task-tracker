package implRepository

import (
	"encoding/json"
	"log"
	"os"
	"task-cli/internal/model"
)

type JsonTaskRepository struct {
	tasks []model.Task
}

func NewJsonTaskRepository() *JsonTaskRepository {
	return &JsonTaskRepository{make([]model.Task, 0)}
}

func (jtr *JsonTaskRepository) GetTasks() []model.Task {
	return jtr.tasks
}

func (jtr *JsonTaskRepository) SetTasks(tasks []model.Task) {
	jtr.tasks = tasks
}

func (jtr *JsonTaskRepository) AddTask(task ...model.Task) {
	jtr.tasks = append(jtr.tasks, task...)
}

func (jtr *JsonTaskRepository) IsExistFile(path string) (bool, error) {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false, nil
	}
	return err == nil, err
}

func (jtr *JsonTaskRepository) CreateJsonFile(path string) error {
	file, err := os.Create(path)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write([]byte("[]"))
	return err
}

func (jtr *JsonTaskRepository) Read(path string) {
	existFile, err := jtr.IsExistFile(path)
	if err != nil {
		log.Fatal(err)
	}

	if !existFile {
		if err := jtr.CreateJsonFile(path); err != nil {
			log.Fatal(err)
		}
		jtr.SetTasks([]model.Task{})
		return
	}

	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	if len(bytes) == 0 {
		jtr.SetTasks([]model.Task{})
		return
	}

	var tasks []model.Task
	if err := json.Unmarshal(bytes, &tasks); err != nil {
		log.Fatal(err)
	}

	jtr.SetTasks(tasks)
}

func (jtr *JsonTaskRepository) Write(path string) {
	bytes, err := json.MarshalIndent(jtr.GetTasks(), "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(path, bytes, 0644); err != nil {
		log.Fatal(err)
	}
}
