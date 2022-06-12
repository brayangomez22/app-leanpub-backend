package domain

import "leanpub-app/domain/model"

type DatabaseGateway interface {
	SaveUser(user *model.User) (*model.User, error)
	ValidateUser(registeredUser *model.RegisteredUser, user *model.User) (*model.User, error)
	GetUsers() (*[]model.User, error)
	GetUserById(id string) (*model.User, error)
	DeleteUser(id string) error
	UpdateUser(user *model.User) (*model.User, error)
	SaveBook(book *model.Book) (*model.Book, error)
	GetBooks() (*[]model.Book, error)
	GetBookById(id string) (*model.Book, error)
	GetBookByAuthor(authorId string) (*model.Book, error)
	GetBookByCategory(category string) (*model.Book, error)
	DeleteBook(id string) error
	UpdateBook(book *model.Book) (*model.Book, error)
	Setup()
}
