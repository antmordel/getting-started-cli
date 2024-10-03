package internal

import (
	"fmt"
	"time"
)

type Todo struct {
	ID     int
	Text   string
	Done   bool
	DoneBy time.Time
}

type TodoDB interface {
	New(text string)
	Get(id int) (*Todo, error)
	All() []*Todo
}

type InMemoryTodoDB struct {
	todos []*Todo
}

func NewInMemoryTodoDB() *InMemoryTodoDB {
	return &InMemoryTodoDB{
		todos: make([]*Todo, 0),
	}
}

func (db *InMemoryTodoDB) New(text string) {
	todo := &Todo{
		ID:     db.generateId(),
		Text:   text,
		Done:   false,
		DoneBy: time.Time{},
	}
	db.todos = append(db.todos, todo)
}

func (db *InMemoryTodoDB) Get(id int) (*Todo, error) {
	for _, todo := range db.todos {
		if todo.ID == id {
			return todo, nil
		}
	}
	return nil, fmt.Errorf("todo with id %d not found", id)
}

func (db *InMemoryTodoDB) All() []*Todo {
	return db.todos
}

func (db *InMemoryTodoDB) generateId() int {
	max := 0
	for _, todo := range db.todos {
		if todo.ID > max {
			max = todo.ID
		}
	}
	return max + 1
}
