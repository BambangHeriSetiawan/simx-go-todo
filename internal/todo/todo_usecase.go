package todo

type TodoUsecase interface {
	GetTodos() ([]Todo, error)
	CreateTodo(todo Todo) error
	UpdateTodo(id string, todo Todo) error
	DeleteTodo(id string) error
}

type todoUsecase struct {
	repo TodoRepository
}

func NewTodoUsecase(r TodoRepository) TodoUsecase {
	return &todoUsecase{repo: r}
}

// Implement methods here

func (u *todoUsecase) GetTodos() ([]Todo, error) {
	return u.repo.GetTodos()
}

func (u *todoUsecase) CreateTodo(todo Todo) error {
	return u.repo.CreateTodo(todo)
}

func (u *todoUsecase) UpdateTodo(id string, todo Todo) error {
	return u.repo.UpdateTodo(id, todo)
}

func (u *todoUsecase) DeleteTodo(id string) error {
	return u.repo.DeleteTodo(id)
}
