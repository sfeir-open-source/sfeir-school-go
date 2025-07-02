package model

import (
	"encoding/json"
	"testing"
	"time"
)

func TestEqual(t *testing.T) {
	// load NY time zone EST
	ny, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Error("error loading time zone")
	}

	// load LA time zone PST
	sf, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		t.Error("error loading time zone")
	}

	// build a task with dates in NY
	task1 := Task{
		ID:           "07b0a105-086e-4a66-8642-a6546a747434",
		Title:        "test",
		Description:  "test equal",
		Status:       StatusTodo,
		Priority:     PriorityHigh,
		CreationDate: time.Date(2019, 12, 24, 10, 0, 0, 0, ny),
		DueDate:      time.Date(2019, 12, 25, 10, 0, 0, 0, ny),
	}

	// copy the task
	task2 := task1

	// update the date in LA time zone BUT representing the same moment
	// at 10 AM in NY, it's 7 AM in SF, but both dates represent the same moment
	task2.CreationDate = time.Date(2019, 12, 24, 7, 0, 0, 0, sf)
	task2.DueDate = time.Date(2019, 12, 25, 7, 0, 0, 0, sf)

	// test that the equal is checking correctly the dates with Equal and not ==
	if !task1.Equal(task2) {
		t.Error("wrong equal implementation")
	}

}

func TestMarshalling(t *testing.T) {

	goldenMaster := "{\"id\":\"07b0a105-086e-4a66-8642-a6546a747434\",\"title\":\"test\",\"description\":\"test marshall\",\"status\":0,\"priority\":2,\"creation_date\":\"2019-12-24T10:00:00Z\",\"due_date\":\"2019-12-25T10:00:00Z\"}"

	tks := Task{
		ID:           "07b0a105-086e-4a66-8642-a6546a747434",
		Title:        "test",
		Description:  "test marshall",
		Status:       StatusTodo,
		Priority:     PriorityHigh,
		CreationDate: time.Date(2019, 12, 24, 10, 0, 0, 0, time.UTC),
		DueDate:      time.Date(2019, 12, 25, 10, 0, 0, 0, time.UTC),
	}

	// marshalling
	bin, err := json.Marshal(tks)
	if err != nil {
		t.Error(err)
	}

	// testing marshalling with golden master
	if string(bin) != goldenMaster {
		t.Error("marshalling is different from expected")
	}

	// unmarshalling
	tke := Task{}
	json.Unmarshal(bin, &tke)

	// testing marshalling/unmarshalling equality
	if !tks.Equal(tke) {
		t.Error("marshalling/unmarshalling broke the equality")
	}

}
