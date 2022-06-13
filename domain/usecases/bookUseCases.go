package usecases

import (
	"github.com/google/uuid"
	"leanpub-app/domain"
	"leanpub-app/domain/models"
	"leanpub-app/domain/models/dtos"
	"log"
)

type BookUseCase struct {
	datastore domain.DatabaseGateway
}

func NewBookUseCase(datastore domain.DatabaseGateway) BookUseCase {
	return BookUseCase{
		datastore: datastore,
	}
}

func (bookUseCase BookUseCase) SaveBook(book *dtos.BookDto) (*models.Book, error) {
	var bookSection []models.BookSection
	var newContents []models.BookContent
	for _, content := range book.Content {
		var sections []string
		for _, section := range content.Sections {
			id, _ := uuid.NewRandom()

			newBookSection := models.BookSection{
				Id: id.String(),
				Title: section.Title,
				Content: section.Content,
			}

			bookSection = append(bookSection, newBookSection)
			sections = append(sections, newBookSection.Id)
		}

		newContent := models.BookContent{
			Chapter: content.Chapter,
			Sections: sections,
		}

		newContents = append(newContents, newContent)
	}

	for _, section := range bookSection {
		err := bookUseCase.datastore.SaveBookSections(&section)
		if err != nil {
			log.Print(err)
		}
	}

	id, _ := uuid.NewRandom()

	newBook := models.Book{
		Id: id.String(),
		Authors: book.Authors,
		AuthorCount: book.AuthorCount,
		Title: book.Title,
		AboutTheBook: book.AboutTheBook,
		Description: book.Description,
		Content: newContents,
		CoverImage: book.CoverImage,
		MinimumPrice: book.MinimumPrice,
		SuggestedPrice: book.SuggestedPrice,
		Reviews: book.Reviews,
		State: book.State,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
		LanguageName: book.LanguageName,
		LanguageCode: book.LanguageCode,
		Categories: book.Categories,
		ReadingOptions: book.ReadingOptions,
	}

	return bookUseCase.datastore.SaveBook(&newBook)
}

func (bookUseCase BookUseCase) GetBooks() (*[]models.Book, error) {
	return bookUseCase.datastore.GetBooks()
}

func (bookUseCase BookUseCase) GetBookIndex(id string) (*[]models.BookContent, error) {
	return bookUseCase.datastore.GetBookIndex(id)
}

func (bookUseCase BookUseCase) GetBookById(id string) (*models.Book, error) {
	return bookUseCase.datastore.GetBookById(id)
}

func (bookUseCase BookUseCase) GetBookByAuthor(authorId string) (*[]models.Book, error) {
	return bookUseCase.datastore.GetBookByAuthor(authorId)
}

func (bookUseCase BookUseCase) GetBookByCategory(category string) (*[]models.Book, error) {
	return bookUseCase.datastore.GetBookByCategory(category)
}

func (bookUseCase BookUseCase) DeleteBook(id string) error {
	return bookUseCase.datastore.DeleteBook(id)
}

func (bookUseCase BookUseCase) UpdateBook(book *models.Book) (*models.Book, error) {
	return bookUseCase.datastore.UpdateBook(book)
}
