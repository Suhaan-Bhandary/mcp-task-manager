package task

import "github.com/Suhaan-Bhandary/mcp-task-manager/repo"

type Service interface {
	Create(task CreateTaskRequest) error
	Update(id string, task UpdateTaskRequest) error
	List() ([]Task, error)
	Get(id string) (Task, error)
	Delete(id string) error
}

type service struct {
	repo repo.TaskRepo
}

func NewService(repo repo.TaskRepo) Service {
	return &service{repo: repo}
}

func (s *service) Create(task CreateTaskRequest) error {
	dbTask := repo.CreateTask{
		Title:       task.Title,
		Description: task.Description,
		Status:      string(task.Status),
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}

	return s.repo.Create(dbTask)
}

func (s *service) Update(id string, task UpdateTaskRequest) error {
	var status *string
	if task.Status != nil {
		s := string(*task.Status)
		status = &s
	}

	updatedDBTask := repo.UpdatedTask{
		Title:       task.Title,
		Description: task.Description,
		Status:      status,
	}

	return s.repo.Update(id, updatedDBTask)
}

func (s *service) List() ([]Task, error) {
	dbTasks, err := s.repo.List()
	if err != nil {
		return nil, err
	}

	tasks := make([]Task, 0, len(dbTasks))
	for _, dbTask := range dbTasks {
		task := MapDBToTask(dbTask)
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s *service) Get(id string) (Task, error) {
	dbTask, err := s.repo.Get(id)
	if err != nil {
		return Task{}, err
	}

	task := MapDBToTask(dbTask)
	return task, nil
}

func (s *service) Delete(id string) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
