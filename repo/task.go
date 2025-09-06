package repo

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type TaskRepo interface {
	Create(task CreateTask) error
	Update(id string, task UpdatedTask) error
	List() ([]Task, error)
	Get(id string) (Task, error)
	Delete(id string) error
}

type taskRepo struct {
	db *sql.DB
}

type Task struct {
	Id          string
	Title       string
	Description string
	Status      string
	CreatedAt   int
	UpdatedAt   int
}

type CreateTask struct {
	Title       string
	Description string
	Status      string
	CreatedAt   int
	UpdatedAt   int
}

type UpdatedTask struct {
	Title       *string
	Description *string
	Status      *string
	CreatedAt   *int
	UpdatedAt   *int
}

func NewTask(db *sql.DB) TaskRepo {
	return &taskRepo{db: db}
}

func (r *taskRepo) Create(task CreateTask) error {
	query := `
		INSERT INTO tasks (id, title, description, status, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	id := uuid.NewString()
	_, err := r.db.Exec(query, id, task.Title, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
	if err != nil {
		return fmt.Errorf("create task: %w", err)
	}

	return nil
}

func (r *taskRepo) Update(id string, task UpdatedTask) error {
	query := `
		UPDATE tasks
		SET
			title = COALESCE(?, title),
			description = COALESCE(?, description),
			status = COALESCE(?, status),
			created_at = COALESCE(?, created_at),
			updated_at = COALESCE(?, updated_at)
		WHERE id = ?
	`

	_, err := r.db.Exec(query,
		task.Title,
		task.Description,
		task.Status,
		task.CreatedAt,
		task.UpdatedAt,
		id,
	)
	if err != nil {
		return fmt.Errorf("update task: %w", err)
	}

	return nil
}

func (r *taskRepo) List() ([]Task, error) {
	rows, err := r.db.Query(`
		SELECT id, title, description, status, created_at, updated_at
		FROM tasks
	`)
	if err != nil {
		return nil, fmt.Errorf("list tasks: %w", err)
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *taskRepo) Get(id string) (Task, error) {
	var task Task
	query := `
		SELECT id, title, description, status, created_at, updated_at
		FROM tasks
		WHERE id = ?
	`

	err := r.db.QueryRow(query, id).Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return Task{}, fmt.Errorf("get task: %w", err)
	}

	return task, nil
}

func (r *taskRepo) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM tasks WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("delete task: %w", err)
	}

	return nil
}
