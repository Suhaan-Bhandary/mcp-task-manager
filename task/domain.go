package task

import "github.com/Suhaan-Bhandary/mcp-task-manager/repo"

type Status string

const (
	TODO        Status = "todo"
	IN_PROGRESS Status = "in-progress"
	DONE        Status = "done"
)

type Task struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      Status `json:"status"`
	CreatedAt   int    `json:"created_at"`
	UpdatedAt   int    `json:"updated_at"`
}

type CreateTaskRequest struct {
	Title       string
	Description string
	Status      Status
	CreatedAt   int
	UpdatedAt   int
}

type UpdateTaskRequest struct {
	Title       *string
	Description *string
	Status      *Status
	CreatedAt   *int
	UpdatedAt   *int
}

func MapDBToTask(dbTask repo.Task) Task {
	return Task{
		Id:          dbTask.Id,
		Title:       dbTask.Title,
		Description: dbTask.Description,
		Status:      Status(dbTask.Status),
		CreatedAt:   dbTask.CreatedAt,
		UpdatedAt:   dbTask.UpdatedAt,
	}
}
