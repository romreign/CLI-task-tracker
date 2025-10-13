package main

import (
	"task-cli/internal/handler"
)

func main() {
	handler := handler.NewCliTaskHandler()
	handler.StartHandler()
}
