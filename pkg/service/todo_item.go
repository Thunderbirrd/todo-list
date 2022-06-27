package service

import (
	"github.com/Thunderbirrd/todo-list"
	"github.com/Thunderbirrd/todo-list/pkg/repository"
)

type ToDoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewToDoItemService(repo repository.TodoItem, listRepo repository.TodoList) *ToDoItemService {
	return &ToDoItemService{repo: repo, listRepo: listRepo}
}

func (s *ToDoItemService) Create(userId, listId int, item models.Item) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		// list does not exists or does not belong to user
		return 0, err
	}
	return s.repo.Create(listId, item)
}

func (s *ToDoItemService) GetAll(userId, listId int) ([]models.Item, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *ToDoItemService) GetById(userId, itemId int) (models.Item, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *ToDoItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *ToDoItemService) Update(userid, itemId int, input models.UpdateItemInput) error {
	return s.repo.Update(userid, itemId, input)
}
