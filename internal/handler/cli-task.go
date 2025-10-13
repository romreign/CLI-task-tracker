package handler

import (
	"fmt"
	implRepository "task-cli/internal/repository/impl"
	service "task-cli/internal/service"
	util "task-cli/pkg/cli"
)

type CliTaskHandler struct {
	taskService *service.TaskService
}

func NewCliTaskHandler() *CliTaskHandler {
	jtr := implRepository.NewJsonTaskRepository()
	taskService := service.NewTaskService(jtr)
	return &CliTaskHandler{taskService}
}

func (cli *CliTaskHandler) StartHandler() {
	var request []string = util.ParseCli()

	if len(request) > 3 {
		fmt.Println("Слишком много параметров")
		return
	} else if len(request) == 0 {
		fmt.Println("Вы не предоставили параметры")
		return
	}

	command := request[0]

	switch command {
	case "add":
		if len(request) != 2 {
			fmt.Println("Не верное количество параметров")
			return
		}
		cli.taskService.Add(request[1])
	case "update":
		if len(request) != 3 {
			fmt.Println("Не верное количество параметров")
			return
		}
		cli.taskService.Update(request[1], request[2])
	case "delete":
		if len(request) != 2 {
			fmt.Println("Не верное количество параметров")
			return
		}
		cli.taskService.Delete(request[1])
	case "mark":
		if len(request) != 3 {
			fmt.Println("Не верное количество параметров")
		}
		cli.taskService.Mark(request[1], request[2])

	case "list":
		if len(request) == 2 {
			cli.taskService.PrintTasks(request[1])
		} else if len(request) == 1 {
			cli.taskService.PrintTasks("")
		} else {
			fmt.Println("Не верное количество параметров")
		}
	}
}
