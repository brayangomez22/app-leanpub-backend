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
	SaveBookSections(bookSections []interface{}) error
	GetBooks() (*[]models.Book, error)
	GetBookIndex(id string) (*models.BookIndex, error)
	GetSectionsByBookId(bookId string) (*models.BookSections, error)
	GetBookSectionById(id string) (*models.BookSection, error)
	GetBookById(id string) (*models.Book, error)
	GetBooksByAuthor(authorId string) (*[]models.Book, error)
	GetBooksByCategory(category string) (*[]models.Book, error)
	DeleteBook(id string) error
	UpdateBook(book *models.Book) (*models.Book, error)
	SaveShoppingCart(shoppingCart *models.ShoppingCart) (*models.ShoppingCart, error)
	GetShoppingCarts() (*[]models.ShoppingCart, error)
	GetShoppingCartById(id string) (*models.ShoppingCart, error)
	DeleteShoppingCart(id string) error
	UpdateShoppingCart(shoppingCart *models.ShoppingCart)	(*models.ShoppingCart, error)
	Setup()
}
