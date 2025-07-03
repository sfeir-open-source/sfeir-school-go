package dao_test

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sfeir-open-source/sfeir-school-go/dao"
	"github.com/sfeir-open-source/sfeir-school-go/model"
	"os"
	"testing"
	"time"
)

func TestDAOMongo(t *testing.T) {
	// get host IP
	dbHost := os.Getenv("DB_HOST")
	db := fmt.Sprintf("mongodb://%s/todolist", dbHost)

	daoMongo, err := dao.GetTaskDAO(db, "", dao.DAOMongo)
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

	toSave, err = daoMongo.Create(toSave)
	if err != nil {
		t.Error(err)
	}

	t.Log("initial task saved", toSave)

	tasks, err := daoMongo.GetAll(dao.NoPaging, dao.NoPaging)
	if err != nil {
		t.Error(err)
	}

	// Check if persistence ok, stops here if not
	if len(tasks) != 1 {
		t.Fatal("Expected 2 tasks, got ", len(tasks))
	}

	t.Log("initial task found all", tasks[0])

	oneTask, err := daoMongo.GetByID(tasks[0].ID)
	if err != nil {
		t.Error(err)
	}

	t.Log("initial task found one", oneTask)

	oneTask.Title = "Use Go(lang)"
	oneTask.Description = "Let's build a REST service in Go !"
	oneTask, err = daoMongo.Update(oneTask)
	if err != nil {
		t.Error(err)
	}

	t.Log("initial task modified", oneTask)

	oneTask, err = daoMongo.GetByID(oneTask.ID)
	if err != nil {
		t.Error(err)
	}

	t.Log("initial task found one modified", oneTask)

	tasks, err = daoMongo.GetByStatus(model.StatusTodo)
	if err != nil {
		t.Error(err)
	}

	t.Log("initial task found one by status", tasks)

	err = daoMongo.Delete(oneTask.ID)
	if err != nil {
		t.Error(err)
	}

	oneTask, err = daoMongo.GetByID(oneTask.ID)
	if err == dao.ErrNotFound {
		t.Log("initial task deleted", err, oneTask)
	} else {
		t.Error(err)
	}

}
