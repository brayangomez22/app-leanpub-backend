package domain

import "leanpub-app/domain/models"

type DatabaseGateway interface {
	SaveUser(user *models.User) (*models.User, error)
	ValidateUser(registeredUser *models.RegisteredUser, user *models.User) (*models.User, error)
	GetUsers() (*[]models.User, error)
	GetUserById(id string) (*models.User, error)
	DeleteUser(id string) error
	UpdateUser(user *models.User) (*models.User, error)
	SaveBook(book *models.Book) (*models.Book, error)
	SaveBookSection(bookSection *models.BookSection) error
	GetBooks() (*[]models.Book, error)
	GetBookIndex(id string) (*[]models.BookContent, error)
	GetSectionsByBookId(bookId string) (*[]models.BookSection, error)
	GetBookSectionById(id string) (*models.BookSection, error)
	GetBookById(id string) (*models.Book, error)
	GetBookByAuthor(authorId string) (*[]models.Book, error)
	GetBookByCategory(category string) (*[]models.Book, error)
	DeleteBook(id string) error
	UpdateBook(book *models.Book) (*models.Book, error)
	Setup()
}
