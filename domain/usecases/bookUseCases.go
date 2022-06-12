package usecases

import (
	"leanpub-app/domain"
	"leanpub-app/domain/model"
)

type BookUseCase struct {
	datastore domain.DatabaseGateway
}

func NewBookUseCase(datastore domain.DatabaseGateway) BookUseCase {
	return BookUseCase{
		datastore: datastore,
	}
}

func (bookUseCase BookUseCase) SaveBook(book *model.Book) (*model.Book, error) {
	return bookUseCase.datastore.SaveBook(book)
}

func (bookUseCase BookUseCase) GetBooks() (*[]model.Book, error) {
	return bookUseCase.datastore.GetBooks()
}

func (bookUseCase BookUseCase) GetBookById(id string) (*model.Book, error) {
	return bookUseCase.datastore.GetBookById(id)
}

func (bookUseCase BookUseCase) GetBookByAuthor(authorId string) (*model.Book, error) {
	return bookUseCase.datastore.GetBookByAuthor(authorId)
}

func (bookUseCase BookUseCase) GetBookByCategory(category string) (*[]model.Book, error) {
	return bookUseCase.datastore.GetBookByCategory(category)
}

func (bookUseCase BookUseCase) DeleteBook(id string) error {
	return bookUseCase.datastore.DeleteBook(id)
}

func (bookUseCase BookUseCase) UpdateBook(book *model.Book) (*model.Book, error) {
	return bookUseCase.datastore.UpdateBook(book)
}
