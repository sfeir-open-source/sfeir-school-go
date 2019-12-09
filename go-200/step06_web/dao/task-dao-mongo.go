package dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/sfeir-open-source/sfeir-school-go/model"
	logger "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// compilation time interface check
var _ TaskDAO = (*TaskDAOMongo)(nil)

const (
	todolist = "todolist"
	tasks    = "tasks"
	index    = "id"
)

// TaskDAOMongo is the mongo implementation of the TaskDAO
type TaskDAOMongo struct {
	client *mongo.Client
}

// NewTaskDAOMongo creates a new TaskDAO mongo implementation
func NewTaskDAOMongo(client *mongo.Client) TaskDAO {

	// build a unique index on the id
	mod := mongo.IndexModel{
		Keys: bson.M{
			index: 1, // index in ascending order
		},
		Options: options.Index().SetUnique(true).SetBackground(true).SetSparse(true),
	}

	// create index
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	idx, err := client.Database(todolist).Collection(tasks).Indexes().CreateOne(ctx, mod)

	if err != nil {
		logger.WithField("error", err).WithField("index", idx).Warn("index creation error")
	} else {
		logger.WithField("index", idx).Debug("index created")
	}

	return &TaskDAOMongo{
		client: client,
	}
}

// GetByID returns a task by its ID
func (s *TaskDAOMongo) GetByID(ID string) (*model.Task, error) {

	// check ID
	if _, err := uuid.Parse(ID); err != nil {
		return nil, ErrInvalidUUID
	}

	task := model.Task{}
	c := s.client.Database(todolist).Collection(tasks)
	result := c.FindOne(context.TODO(), bson.M{"id": ID})

	// check error for not found task
	err := result.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrNotFound
		}
		return nil, err
	}

	// decode result and handle error
	err = result.Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, err
}

// getAllTasksByQuery returns tasks by query and paging capability
func (s *TaskDAOMongo) getAllTasksByQuery(query interface{}, start, end int) ([]model.Task, error) {

	// check param
	hasPaging := start > NoPaging && end > NoPaging && end >= start

	c := s.client.Database(todolist).Collection(tasks)

	// perform request
	var tasks []model.Task
	var opt *options.FindOptions

	// check pagination option
	if hasPaging {
		opt = options.Find().SetSkip(int64(start)).SetLimit(int64(end - start + 1))
	}

	// find query
	cur, err := c.Find(context.TODO(), query, opt)
	if err != nil {
		return nil, err
	}

	// defer close cursor
	defer func() {
		err := cur.Close(context.TODO())
		if err != nil {
			logger.WithField("error", err).WithField("query", query).Warn("error closing cursor")
		}
	}()

	// marshall results
	err = cur.All(context.TODO(), &tasks)
	if err != nil {
		return nil, err
	}

	// if no result found wrap with dao known error
	if len(tasks) == 0 {
		return nil, ErrNotFound
	}

	return tasks, err
}

// GetAll returns all tasks with paging capability
func (s *TaskDAOMongo) GetAll(start, end int) ([]model.Task, error) {
	return s.getAllTasksByQuery(bson.M{}, start, end)
}

// GetByTitle returns all tasks by title
func (s *TaskDAOMongo) GetByTitle(title string) ([]model.Task, error) {
	return s.getAllTasksByQuery(bson.M{"title": title}, NoPaging, NoPaging)
}

// GetByStatus returns all tasks by status
func (s *TaskDAOMongo) GetByStatus(status model.TaskStatus) ([]model.Task, error) {
	return s.getAllTasksByQuery(bson.M{"status": status}, NoPaging, NoPaging)
}

// GetByStatusAndPriority returns all tasks by status and priority
func (s *TaskDAOMongo) GetByStatusAndPriority(status model.TaskStatus, priority model.TaskPriority) ([]model.Task, error) {
	return s.getAllTasksByQuery(bson.M{"status": status, "priority": priority}, NoPaging, NoPaging)
}

// Create saves the task
func (s *TaskDAOMongo) Create(task *model.Task) error {

	// check task has an ID, if not create one
	if len(task.ID) == 0 {
		task.ID = uuid.New().String()
	}

	// insert task
	c := s.client.Database(todolist).Collection(tasks)
	_, err := c.InsertOne(context.TODO(), task)

	return err
}

// Update updates a task
func (s *TaskDAOMongo) Update(task *model.Task) error {

	// check ID
	// check task has an ID, if not create one
	if len(task.ID) == 0 {
		task.ID = uuid.New().String()
	}

	// update task
	c := s.client.Database(todolist).Collection(tasks)
	_, err := c.ReplaceOne(context.TODO(), bson.M{"id": task.ID}, task)

	return err
}

// Delete deletes a tasks by its ID
func (s *TaskDAOMongo) Delete(ID string) error {

	// check ID
	if _, err := uuid.Parse(ID); err != nil {
		return ErrInvalidUUID
	}

	// delete task
	c := s.client.Database(todolist).Collection(tasks)
	_, err := c.DeleteOne(context.TODO(), bson.M{"id": ID})

	return err
}
