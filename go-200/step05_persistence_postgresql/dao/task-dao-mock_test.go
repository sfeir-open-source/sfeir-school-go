package dao_test

import (
	"github.com/google/uuid"
	"github.com/sfeir-open-source/sfeir-school-go/dao"
	"github.com/sfeir-open-source/sfeir-school-go/model"
	"testing"
	"time"
)

func TestDAOMock(t *testing.T) {

	daoMock, err := dao.GetTaskDAO("", "", dao.DAOMock)
	if err != nil {
		t.Error(err)
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

	toSave, err = daoMock.Create(toSave)
	if err != nil {
		t.Error(err)
	}

	t.Log("initial task saved", toSave)

	tasks, err := daoMock.GetAll(dao.NoPaging, dao.NoPaging)
	if err != nil {
		t.Error(err)
	}

	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}

	toSave2 := model.Task{
		ID:           uuid.New().String(),
		Title:        "Use Go Again",
		Description:  "Let's use the Go programming language in a real project.",
		Status:       model.StatusDone,
		Priority:     model.PriorityHigh,
		CreationDate: time.Date(2017, 02, 01, 0, 0, 0, 0, time.UTC),
		DueDate:      time.Date(2017, 02, 02, 0, 0, 0, 0, time.UTC),
	}

	toSave2, err = daoMock.Create(toSave2)
	if err != nil {
		t.Error(err)
	}

	// check indexes search
	tasks, err = daoMock.GetAll(0, 0)
	if err != nil {
		t.Error(err)
	}

	if len(tasks) != 1 {
		t.Errorf("Expected 1 tasks, got %d", len(tasks))
	}

	oneTask, err := daoMock.GetByID(toSave.ID)
	if err != nil {
		t.Error(err)
	}
	if !toSave.Equal(oneTask) {
		t.Error("Got wrong task by ID")
	}

	err = daoMock.Delete(oneTask.ID)
	if err != nil {
		t.Error(err)
	}

	oneTask, err = daoMock.GetByID(oneTask.ID)
	if err == nil {
		t.Error("Task should have been deleted", oneTask)
	}

}
