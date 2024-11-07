package service

import (
	"github.com/Muhfikri12/golang-cms-olshop/model"
	"github.com/Muhfikri12/golang-cms-olshop/repository"
)

type BookService struct {
	RepoBook repository.BookRepoDB
}

func NewBookService(repo repository.BookRepoDB) *BookService {
	return &BookService{RepoBook: repo}
}

func (bs *BookService) CreateBookService(book *model.Books) error {
	if err := bs.RepoBook.CreateBookRepo(book); err != nil {
		return err
	}
	return nil
}

func (bs *BookService) ShowALlList() (*[]model.Books, error) {
	
	items, err := bs.RepoBook.BookList(); 
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (bs *BookService) UpdateBook(id int, item *model.Books) error {
	
	if err := bs.RepoBook.UpdateBook(id, item); err != nil {
		return err
	}

	return nil
}

