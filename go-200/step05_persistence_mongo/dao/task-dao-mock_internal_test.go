package dao

import (
	"github.com/google/uuid"
	"github.com/sfeir-open-source/sfeir-school-go/model"
	"testing"
	"time"
)

func TestDAOMockInternal(t *testing.T) {

	daoMock := &TaskDAOMock{
		storage: make(map[string]model.Task),
	}

	toSave := model.Task{
		ID:           uuid.New().String(),
		Title:        "Use Go",
		Description:  "Let's use the Go programming language in a real project.",
		Status:       model.StatusTodo,
		Priority:     model.PriorityMedium,
		CreationDate: time.Date(2017, 02, 01, 0, 0, 0, 0, time.UTC),
		DueDate:      time.Date(2017, 02, 02, 0, 0, 0, 0, time.UTC),
	}

	daoMock.Create(toSave)

	tasks, err := daoMock.getBy(func(task model.Task) bool {
		return task.Status == model.StatusTodo
	})

	if err != nil {
		t.Errorf("Error retrieving tasks %v", err)
	}

	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}
	if tasks[0] != toSave {
		t.Error("Got wrong task from mocked DAO.")
	}

}
