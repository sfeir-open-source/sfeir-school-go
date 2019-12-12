package dao

import (
	"errors"
	"github.com/sfeir-open-source/sfeir-school-go/model"
)

const (
	// NoPaging used with skip, limit parameters
	NoPaging = -1
)

var (
	// ErrInvalidUUID is used on invalid UUID number
	ErrInvalidUUID = errors.New("invalid input to UUID")

	// ErrNotFound is used when no result are found for the given parameters
	ErrNotFound = errors.New("no result found")
)

// TaskDAO is the DAO interface to work with tasks
type TaskDAO interface {

	// GetByID returns a task by its ID
	GetByID(ID string) (model.Task, error)

	// GetAll returns all tasks with paging capability
	GetAll(start, end int) ([]model.Task, error)

	// GetByTitle returns all tasks by title
	GetByTitle(title string) ([]model.Task, error)

	// GetByStatus returns all tasks by status
	GetByStatus(status model.TaskStatus) ([]model.Task, error)

	// GetByStatusAndPriority returns all tasks by status and priority
	GetByStatusAndPriority(status model.TaskStatus, priority model.TaskPriority) ([]model.Task, error)

	// Create creates a new task
	Create(task model.Task) (model.Task, error)

	// Update updates a task
	Update(task model.Task) (model.Task, error)

	// Delete deletes a tasks by its ID
	Delete(ID string) error
}
