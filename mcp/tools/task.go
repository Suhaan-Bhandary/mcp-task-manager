package tools

import (
	"context"
	"time"

	"github.com/Suhaan-Bhandary/mcp-task-manager/task"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type TaskHandler interface {
	Create(ctx context.Context, req *mcp.CallToolRequest, input CreateTaskInput) (*mcp.CallToolResult, MessageOutput, error)
	Update(ctx context.Context, req *mcp.CallToolRequest, input UpdateTaskInput) (*mcp.CallToolResult, MessageOutput, error)
	List(ctx context.Context, req *mcp.CallToolRequest, _ any) (*mcp.CallToolResult, ListTaskOutput, error)
	Get(ctx context.Context, req *mcp.CallToolRequest, input TaskIDInput) (*mcp.CallToolResult, TaskOutput, error)
	Delete(ctx context.Context, req *mcp.CallToolRequest, input TaskIDInput) (*mcp.CallToolResult, MessageOutput, error)
}

type taskHandler struct {
	taskService task.Service
}

func NewTaskHandler(taskService task.Service) TaskHandler {
	return &taskHandler{taskService: taskService}
}

func (h *taskHandler) Create(ctx context.Context, req *mcp.CallToolRequest, input CreateTaskInput) (*mcp.CallToolResult, MessageOutput, error) {
	now := int(time.Now().Unix())
	newTask := task.CreateTaskRequest{
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	err := h.taskService.Create(newTask)
	if err != nil {
		return nil, MessageOutput{}, err
	}

	return nil, MessageOutput{Message: "Task created successfully"}, nil
}

func (h *taskHandler) Update(ctx context.Context, req *mcp.CallToolRequest, input UpdateTaskInput) (*mcp.CallToolResult, MessageOutput, error) {
	now := int(time.Now().Unix())
	updateReq := task.UpdateTaskRequest{
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
		UpdatedAt:   &now,
	}

	err := h.taskService.Update(input.Id, updateReq)
	if err != nil {
		return nil, MessageOutput{}, err
	}

	return nil, MessageOutput{Message: "Task updated successfully"}, nil
}

func (h *taskHandler) List(ctx context.Context, req *mcp.CallToolRequest, _ any) (*mcp.CallToolResult, ListTaskOutput, error) {
	tasks, err := h.taskService.List()
	if err != nil {
		return nil, ListTaskOutput{}, err
	}

	return nil, ListTaskOutput{Tasks: tasks}, nil
}

func (h *taskHandler) Get(ctx context.Context, req *mcp.CallToolRequest, input TaskIDInput) (*mcp.CallToolResult, TaskOutput, error) {
	taskObj, err := h.taskService.Get(input.Id)
	if err != nil {
		return nil, TaskOutput{}, err
	}

	return nil, TaskOutput{Task: taskObj}, nil
}

func (h *taskHandler) Delete(ctx context.Context, req *mcp.CallToolRequest, input TaskIDInput) (*mcp.CallToolResult, MessageOutput, error) {
	err := h.taskService.Delete(input.Id)
	if err != nil {
		return nil, MessageOutput{}, err
	}

	return nil, MessageOutput{Message: "Task deleted successfully"}, nil
}
