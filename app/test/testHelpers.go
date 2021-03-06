package test

import (
	"github.com/stretchr/testify/mock"
	"leanpub-app/domain/models"
)

type DbGateway struct {
	mock.Mock
}

func NewDbGateway() DbGateway {
	return DbGateway{}
}

func (db DbGateway) Setup() {}

func (db DbGateway) SaveUser(user *models.User) (*models.User, error) {
	args := db.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (db DbGateway) ValidateUser(registeredUser *models.RegisteredUser, user *models.User) (*models.User, error) {
	args := db.Called(registeredUser, user)
	if args.Get(0) == nil || args.Get(1) == nil {
		return nil, args.Error(1)
	}
	return args.Get(1).(*models.User), args.Error(1)
}

func (db DbGateway) GetUsers() (*[]models.User, error) {
	args := db.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]models.User), args.Error(1)
}

func (db DbGateway) GetUserById(id string) (*models.User, error){
	args := db.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (db DbGateway) DeleteUser(id string) error {
	args := db.Called(id)
	return args.Error(0)
}

func (db DbGateway) UpdateUser(user *models.User) (*models.User, error){
	args := db.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (db DbGateway) SaveBook(book *models.Book) (*models.Book, error) {
	args := db.Called(book)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Book), args.Error(1)
}

func (db DbGateway) SaveBookSection(bookSection *models.BookSection) error {
	args := db.Called(bookSection)
	return args.Error(0)
}

func (db DbGateway) SaveBookSections(bookSections []interface{}) error {
	args := db.Called(bookSections)
	return args.Error(0)
}

func (db DbGateway) GetBooks() (*[]models.Book, error) {
	args := db.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]models.Book), args.Error(1)
}

func (db DbGateway) GetBookIndex(id string) (*models.BookIndex, error) {
	args := db.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.BookIndex), args.Error(1)
}

func (db DbGateway) GetSectionsByBookId(bookId string) (*models.BookSections, error) {
	args := db.Called(bookId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.BookSections), args.Error(1)
}

func (db DbGateway) GetBookSectionById(id string) (*models.BookSection, error) {
	args := db.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.BookSection), args.Error(1)
}

func (db DbGateway) GetBookById(id string) (*models.Book, error) {
	args := db.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Book), args.Error(1)
}

func (db DbGateway) GetBooksByAuthor(authorId string) (*[]models.Book, error) {
	args := db.Called(authorId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]models.Book), args.Error(1)
}

func (db DbGateway) GetBooksByCategory(category string) (*[]models.Book, error) {
	args := db.Called(category)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]models.Book), args.Error(1)
}

func (db DbGateway) DeleteBook(id string) error {
	args := db.Called(id)
	return args.Error(0)
}

func (db DbGateway) UpdateBook(book *models.Book) (*models.Book, error) {
	args := db.Called(book)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Book), args.Error(1)
}

func (db DbGateway) SaveShoppingCart(shoppingCart *models.ShoppingCart) (*models.ShoppingCart, error) {
	panic("implement me")
}

func (db DbGateway) GetShoppingCarts() (*[]models.ShoppingCart, error) {
	panic("implement me")
}

func (db DbGateway) GetShoppingCartById(id string) (*models.ShoppingCart, error) {
	panic("implement me")
}

func (db DbGateway) DeleteShoppingCart(id string) error {
	panic("implement me")
}

func (db DbGateway) UpdateShoppingCart(shoppingCart *models.ShoppingCart) (*models.ShoppingCart, error) {
	panic("implement me")
}