package tools

import "github.com/Suhaan-Bhandary/mcp-task-manager/task"

type CreateTaskInput struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Status      task.Status `json:"status"`
}

type UpdateTaskInput struct {
	Id          string       `json:"id"`
	Title       *string      `json:"title"`
	Description *string      `json:"description"`
	Status      *task.Status `json:"status"`
}

type TaskIDInput struct {
	Id string `json:"id"`
}

type TaskOutput struct {
	Task task.Task `json:"task"`
}

type ListTaskOutput struct {
	Tasks []task.Task `json:"tasks"`
}
