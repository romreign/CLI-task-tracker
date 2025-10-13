package repository

import (
	"task-cli/internal/model"
)

type TaskRepository interface {
	Read(path string)
	Write(path string)
	SetTasks([]model.Task)
	GetTasks() []model.Task
}
