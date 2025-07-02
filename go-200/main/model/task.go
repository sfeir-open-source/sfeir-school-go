package model

import (
	"github.com/google/uuid"
	"time"
)

// TaskStatus is the current processing status of a task
type TaskStatus int

const (
	// StatusTodo is used for incomplete tasks
	StatusTodo TaskStatus = iota
	// StatusInProgress is used for tasks in progress
	StatusInProgress
	// StatusDone is used for completed tasks
	StatusDone
)

// TaskPriority is the priority of a task
type TaskPriority int

const (
	// PriorityMinor is used for task with a lower priority
	PriorityMinor TaskPriority = iota
	// PriorityMedium is used for task with medium priority
	PriorityMedium
	// PriorityHigh is used for task with high priority
	PriorityHigh
)

// Task is the structure to define a task to be done
type Task struct {
	ID           string       `json:"id,omitempty" bson:"id"`
	Title        string       `json:"title" bson:"title"`
	Description  string       `json:"description" bson:"description"`
	Status       TaskStatus   `json:"status" bson:"status"`
	Priority     TaskPriority `json:"priority" bson:"priority"`
	CreationDate time.Time    `json:"creation_date" bson:"creation_date"`
	DueDate      time.Time    `json:"due_date" bson:"due_date"`
}

// NewTask sets a new ID of the Task as a string
func NewTask(title, description string) *Task {
	return &Task{
		ID:           uuid.New().String(),
		CreationDate: time.Now(),
		DueDate:      time.Now().Add(24 * time.Hour),
		Status:       StatusTodo,
		Priority:     PriorityMedium,
		Title:        title,
		Description:  description,
	}
}

// Equal compares a Task to another and returns true if they are equal false otherwise
func (t Task) Equal(task Task) bool {
	return t.ID == task.ID &&
		t.Title == task.Title &&
		t.Description == task.Description &&
		t.Status == task.Status &&
		t.Priority == task.Priority &&
		t.CreationDate.Equal(task.CreationDate) &&
		t.DueDate.Equal(task.DueDate)
}
