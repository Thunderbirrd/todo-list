package repository

import (
	"github.com/Thunderbirrd/todo-list"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type TodoList interface {
	Create(userId int, list models.List) (int, error)
	GetAll(userId int) ([]models.List, error)
	GetById(userId, listId int) (models.List, error)
	Delete(userId, listId int) error
	Update(userid, listId int, input models.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item models.Item) (int, error)
	GetAll(userId, listId int) ([]models.Item, error)
	GetById(userId, itemId int) (models.Item, error)
	Delete(userId, itemId int) error
	Update(userid, itemId int, input models.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewToDoItemPostgres(db),
	}
}
