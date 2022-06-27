package service

import (
	"github.com/Thunderbirrd/todo-list"
	"github.com/Thunderbirrd/todo-list/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list models.List) (int, error)
	GetAll(userId int) ([]models.List, error)
	GetById(userId, listId int) (models.List, error)
	Delete(userId, listId int) error
	Update(userid, listId int, input models.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item models.Item) (int, error)
	GetAll(userId, listId int) ([]models.Item, error)
	GetById(userId, itemId int) (models.Item, error)
	Delete(userId, itemId int) error
	Update(userid, itemId int, input models.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewToDoItemService(repos.TodoItem, repos.TodoList),
	}
}
