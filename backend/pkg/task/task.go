package task

import (
	"errors"
	"time"
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

// to-do list item.
type Task struct {
	ID        string
	Title     string
	Completed bool
	Deadline  time.Time
}

// managing tasks.
type Service struct {
	db *sql.DB
}

// new task service.
func NewService() *Service {

	const file string = "todo.db"
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		panic(err)
	}

	const create string = `
  CREATE TABLE IF NOT EXISTS todos (
  ID VARCHAR(256) NOT NULL PRIMARY KEY,
  Title VARCHAR(256) NOT NULL,
  Completed BOOLEAN NOT NULL,
	Deadline INT NOT NULL
  );`

	_, err = db.Exec(create);
	if err != nil {
		panic(err);
	}

	return &Service{
		db: db,
	}
}

func (s *Service) CreateTask(title string, deadline time.Time) (*Task, error) {
	id := generateID()
	task := &Task{
		ID:        id,
		Title:     title,
		Completed: false,
		Deadline:  deadline,
	}

	_, err := s.db.Exec("INSERT INTO todos VALUES(?,?,?,?);", id, title, 0, deadline.Unix())

	if(err != nil){
		return nil, err;
	}

	return task, nil
}

func (s *Service) DeleteTask(id string) error {
  const dele string = "DELETE FROM todos WHERE ID = ?;"

  _, err := s.db.Exec(dele, id)

	if err != nil {
		return errors.New("task not found")
	}
	return nil
}

func (s *Service) MarkTaskComplete(id string, completed bool) error {
	var update string;	
	if(completed){
		update = "UPDATE todos SET Completed = 1 WHERE ID = ?;"
	} else {
		update = "UPDATE todos SET Completed = 0 WHERE ID = ?;"
	}


	_, err := s.db.Exec(update, id);

	if err != nil {
		return errors.New("task not found")
	}
	return nil
}

func (s *Service) GetTask(id string) (*Task, error) {
	const getTaski string = "SELECT * FROM todos WHERE ID = ?;"
	row := s.db.QueryRow(getTaski, id);

	var idi string;
	var title string;
	var completed bool;
	var deadlineEpoch int64;
	var deadline time.Time;

	err := row.Scan(&idi, &title, &completed, &deadlineEpoch);
	if(err != nil){
		return nil, err
	}
	deadline = time.Unix(deadlineEpoch, 0);

	return &Task{
		ID: idi,
		Title: title,
		Completed: completed,
		Deadline: deadline,
	}, nil
}

func (s *Service) ListTasks() []*Task {
	// tasks := make([]*Task, 0, len(s.tasks))

	const list string = "SELECT * FROM todos;"
	rows, err := s.db.Query(list);

	if err != nil {
		return nil
	}

	var tasks []*Task;
	for rows.Next() {
		var id string;
		var title string;
		var completed bool;
		var deadlineEpoch int64;
		var deadline time.Time;

		err := rows.Scan(&id, &title, &completed, &deadlineEpoch);
		if(err != nil){
			return nil
		}
		deadline = time.Unix(deadlineEpoch, 0);
		tasks = append(tasks, &Task{
			ID: id,
			Title: title,
			Completed: completed,
			Deadline: deadline,
		})

	}
	return tasks
}

func generateID() string {
	return uuid.New().String()
}
