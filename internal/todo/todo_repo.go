package todo

import (
	"context"
	"fmt"
	"simx-go-todo/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Todo struct {
	ID    string
	Title string
	Done  bool
}

type todoRepo struct {
    db *pgxpool.Pool
}

type TodoRepository interface {
	GetTodos() ([]Todo, error)
	CreateTodo(todo Todo) error
	UpdateTodo(id string, todo Todo) error
	DeleteTodo(id string) error
}

// GetPool exposes the underlying pgxpool.Pool for migration and admin tasks
func (r *todoRepo) GetPool() *pgxpool.Pool {
	return r.db
}


func NewTodoRepository() (TodoRepository, error) {
	if config.DB == nil {
		return nil, ErrNoDBConnection
	}
	return &todoRepo{db: config.DB}, nil
}
var ErrNoDBConnection = fmt.Errorf("database connection is not initialized")


func (r *todoRepo) GetTodos() ([]Todo, error) {
	rows, err := r.db.Query(context.Background(), "SELECT id, title, done FROM todos ORDER BY id")
	if err != nil {
		return []Todo{}, err
	}
	defer rows.Close()
	var todos []Todo
	for rows.Next() {
		var t Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Done); err != nil {
			return []Todo{}, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}

func (r *todoRepo) CreateTodo(todo Todo) error {
	_, err := r.db.Exec(context.Background(), "INSERT INTO todos (id, title, done) VALUES ($1, $2, $3)", todo.ID, todo.Title, todo.Done)
	return err
}

func (r *todoRepo) UpdateTodo(id string, todo Todo) error {
	_, err := r.db.Exec(context.Background(), "UPDATE todos SET title=$1, done=$2 WHERE id=$3", todo.Title, todo.Done, id)
	return err
}

func (r *todoRepo) DeleteTodo(id string) error {
	_, err := r.db.Exec(context.Background(), "DELETE FROM todos WHERE id=$1", id)
	return err
}
