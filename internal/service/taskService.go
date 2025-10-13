package service

import (
	"fmt"
	"log"
	"strconv"
	"task-cli/internal/model"
	"task-cli/internal/repository"
	"time"
)

var path string = "../../storage/tasks.json"

type TaskService struct {
	taskRepository repository.TaskRepository
}

func NewTaskService(taskRepository repository.TaskRepository) *TaskService {
	return &TaskService{taskRepository}
}

func (ts *TaskService) Add(description string) {
	ts.taskRepository.Read(path)
	tasks := ts.taskRepository.GetTasks()
	idNewTask := findMaxId(tasks) + 1
	task := model.NewTask(uint(idNewTask), description, model.TODO, model.DateTime(time.Now()), model.DateTime(time.Now()))
	ts.taskRepository.SetTasks(append(tasks, *task))
	ts.taskRepository.Write(path)
}

func (ts *TaskService) Update(idStr string, description string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err)
		return
	}
	if id < 0 {
		fmt.Println("id не может быть отрицательным")
		return
	}

	ts.taskRepository.Read(path)
	tasks := ts.taskRepository.GetTasks()
	task := findTaskById(tasks, uint(id))
	if task == nil {
		fmt.Println("Задача с id", id, "не найдена")
		return
	}

	task.SetDescription(description)
	task.SetUpdatedAt(model.DateTime(time.Now()))
	ts.taskRepository.Write(path)
}

func (ts *TaskService) Delete(idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err)
		return
	}
	if id < 0 {
		fmt.Println("id не может быть отрицательным")
		return
	}

	ts.taskRepository.Read(path)
	tasks := ts.taskRepository.GetTasks()
	index := findIndexTaskById(tasks, uint(id))
	if index < 0 {
		fmt.Println("Задача с id", id, "не найдена")
		return
	}

	tmpTask := tasks[index]
	tasks[index] = tasks[len(tasks)-1]
	tasks[len(tasks)-1] = tmpTask
	newTasks := tasks[:len(tasks)-1]
	ts.taskRepository.SetTasks(newTasks)
	ts.taskRepository.Write(path)
}

func (ts *TaskService) Mark(status string, idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err)
		return
	}
	if id < 0 {
		fmt.Println("id не может быть отрицательным")
		return
	}

	switch status {
	case "in-progress":
		ts.markInProgress(id)
	case "done":
		ts.markDone(id)
	default:
		fmt.Println("Не правильные параметры")
		return
	}

	ts.taskRepository.Write(path)

}

func (ts *TaskService) markInProgress(id int) {
	task := ts.findTargetTask(id)
	if task == nil {
		fmt.Println("Задача с id", id, "не найдена")
		return
	}

	if task.GetStatus() == model.IN_PROGRESS {
		return
	}

	task.SetStatus(model.IN_PROGRESS)
	ts.taskRepository.Write(path)
}

func (ts *TaskService) markDone(id int) {
	task := ts.findTargetTask(id)
	if task == nil {
		fmt.Println("Задача с id", id, "не найдена")
		return
	}

	if task.GetStatus() == model.DONE {
		return
	}

	task.SetStatus(model.DONE)
	ts.taskRepository.Write(path)
}

func (ts *TaskService) PrintTasks(status string) {
	switch status {
	case "todo":
		ts.printTodoTasks()
	case "in-progress":
		ts.printInProgressTasks()
	case "done":
		ts.printDoneTasks()
	case "":
		ts.printAllTasks()
	default:
		fmt.Println("Не правильные параметры")
	}
}

func (ts *TaskService) printAllTasks() {
	ts.taskRepository.Read(path)
	tasks := ts.taskRepository.GetTasks()

	if len(tasks) != 0 {
		fmt.Println("Все задачи: ")
		fmt.Println(tasks)
	} else {
		fmt.Println("Нет задач")
	}
}

func (ts *TaskService) printTodoTasks() {
	ts.taskRepository.Read(path)
	allTasks := ts.taskRepository.GetTasks()
	tasks := filterTask(allTasks, model.TODO)
	fmt.Println("Все не просмотренные задачи: ")
	fmt.Println(tasks)
}

func (ts *TaskService) printInProgressTasks() {
	ts.taskRepository.Read(path)
	allTasks := ts.taskRepository.GetTasks()
	tasks := filterTask(allTasks, model.IN_PROGRESS)

	if len(tasks) != 0 {
		fmt.Println("Все выполняемые задачи: ")
		fmt.Println(tasks)
	} else {
		fmt.Println("Нет выполняемых задач")
	}
}

func (ts *TaskService) printDoneTasks() {
	ts.taskRepository.Read(path)
	allTasks := ts.taskRepository.GetTasks()
	tasks := filterTask(allTasks, model.DONE)

	if len(tasks) != 0 {
		fmt.Println("Все завершенные задачи: ")
		fmt.Println(tasks)
	} else {
		fmt.Println("Нет завершенных задач")
	}
}

func filterTask(allTasks []model.Task, status model.TaskStatus) []model.Task {
	var tasks []model.Task

	for _, v := range allTasks {
		if v.GetStatus() == status {
			tasks = append(tasks, v)
		}
	}
	return tasks
}

func (ts *TaskService) findTargetTask(id int) *model.Task {
	ts.taskRepository.Read(path)
	allTasks := ts.taskRepository.GetTasks()

	for i := 0; i < len(allTasks); i++ {
		if allTasks[i].GetId() == uint(id) {
			return &allTasks[i]
		}
	}

	return nil
}

func findMaxId(list []model.Task) int {
	if len(list) == 0 {
		return 0
	}

	max := 0
	for i := 0; i < len(list); i++ {
		if int(list[i].GetId()) > max {
			max = int(list[i].GetId())
		}
	}
	return max
}

func findTaskById(list []model.Task, id uint) *model.Task {
	for i := 0; i < len(list); i++ {
		if list[i].GetId() == id {
			return &list[i]
		}
	}
	return nil
}

func findIndexTaskById(list []model.Task, id uint) int {
	for i := 0; i < len(list); i++ {
		if list[i].GetId() == id {
			return i
		}
	}
	return -1
}
