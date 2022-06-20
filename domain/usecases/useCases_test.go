package usecases

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"leanpub-app/app/test"
	"leanpub-app/domain/models"
	"leanpub-app/domain/models/dtos"
	"testing"
	"time"
)

func TestSaveUserIsOk(t *testing.T) {
	app := test.CreateApp()

	user := &models.User{
		Id:              "1234567890",
		Email:           "test@example.com",
		Password:        "test1234",
		Name:            "test",
		About:           "test",
		AvatarUrl:       "test",
		CreatedAt:       time.Time{},
		IsAdmin:         false,
		HasSubscription: false,
		IsAuthor:        false,
		SocialNetworks:  []models.SocialNetwork{{Name: "Facebook", Url: "url"}},
		UpdatedAt:       time.Time{},
	}

	app.DataStore.On("SaveUser", mock.Anything).Return(user, nil)
	app.DataStore.On("ValidateUser", mock.Anything, mock.Anything).Return(user, nil)

	_, err := UserUseCase{
		datastore: app.DataStore,
	}.SaveUser(user)

	assert.Nil(t, err)
	app.DataStore.MethodCalled("SaveUser", mock.Anything)
	app.DataStore.MethodCalled("ValidateUser", mock.Anything)
}

func TestSaveUserIsWrongConnectionFailed(t *testing.T) {
	app := test.CreateApp()

	user := &models.User{
		Id:              "1234567890",
		Email:           "test@example.com",
		Password:        "test1234",
		Name:            "test",
		About:           "test",
		AvatarUrl:       "test",
		CreatedAt:       time.Time{},
		IsAdmin:         false,
		HasSubscription: false,
		IsAuthor:        false,
		SocialNetworks:  []models.SocialNetwork{{Name: "Facebook", Url: "url"}},
		UpdatedAt:       time.Time{},
	}

	app.DataStore.On("SaveUser", mock.Anything).Return(nil, errors.New("CONNECTION_FAIL"))
	app.DataStore.On("ValidateUser", mock.Anything, mock.Anything).Return(nil, nil)

	_, err := UserUseCase{
		datastore: app.DataStore,
	}.SaveUser(user)

	assert.NotNil(t, err, "CONNECTION_FAIL")
	app.DataStore.MethodCalled("SaveUser", mock.Anything)
}

func TestValidateUserIsOk(t *testing.T) {
	app := test.CreateApp()

	user := &models.User{
		Id:              "1234567890",
		Email:           "test@example.com",
		Password:        "test1234",
		Name:            "test",
		About:           "test",
		AvatarUrl:       "test",
		CreatedAt:       time.Time{},
		IsAdmin:         false,
		HasSubscription: false,
		IsAuthor:        false,
		SocialNetworks:  []models.SocialNetwork{{Name: "Facebook", Url: "url"}},
		UpdatedAt:       time.Time{},
	}

	registerUser := &models.RegisteredUser{
		Email:    "test@example.com",
		Password: "test1234",
	}

	app.DataStore.On("ValidateUser", mock.Anything, mock.Anything).Return(user, nil)

	_, err := UserUseCase{
		datastore: app.DataStore,
	}.ValidateUser(registerUser, user)

	assert.Nil(t, err)
	app.DataStore.MethodCalled("ValidateUser", mock.Anything)
}

func TestValidateUserIsWrongConnectionFailed(t *testing.T) {
	app := test.CreateApp()

	user := &models.User{
		Id:              "1234567890",
		Email:           "test@example.com",
		Password:        "test1234",
		Name:            "test",
		About:           "test",
		AvatarUrl:       "test",
		CreatedAt:       time.Time{},
		IsAdmin:         false,
		HasSubscription: false,
		IsAuthor:        false,
		SocialNetworks:  []models.SocialNetwork{{Name: "Facebook", Url: "url"}},
		UpdatedAt:       time.Time{},
	}

	registerUser := &models.RegisteredUser{
		Email:    "test@example.com",
		Password: "test1234",
	}

	app.DataStore.On("ValidateUser", mock.Anything, mock.Anything).Return(nil, errors.New("UNREGISTERED_USER"))

	_, err := UserUseCase{
		datastore: app.DataStore,
	}.ValidateUser(registerUser, user)

	assert.NotNil(t, err, "UNREGISTERED_USER")
	app.DataStore.MethodCalled("ValidateUser", mock.Anything)
}

func TestUpdateUserIsOk(t *testing.T) {
	app := test.CreateApp()

	user := &models.User{
		Id:              "1234567890",
		Email:           "test@example.com",
		Password:        "test1234",
		Name:            "test",
		About:           "test",
		AvatarUrl:       "test",
		CreatedAt:       time.Time{},
		IsAdmin:         false,
		HasSubscription: false,
		IsAuthor:        false,
		SocialNetworks:  []models.SocialNetwork{{Name: "Facebook", Url: "url"}},
		UpdatedAt:       time.Time{},
	}

	app.DataStore.On("UpdateUser", mock.Anything).Return(user, nil)

	_, err := UserUseCase{
		datastore: app.DataStore,
	}.UpdateUser(user)

	assert.Nil(t, err)
	app.DataStore.MethodCalled("UpdateUser", mock.Anything)
}

func TestUpdateUserIsWrongConnectionFailed(t *testing.T) {
	app := test.CreateApp()

	user := &models.User{
		Id:              "1234567890",
		Email:           "test@example.com",
		Password:        "test1234",
		Name:            "test",
		About:           "test",
		AvatarUrl:       "test",
		CreatedAt:       time.Time{},
		IsAdmin:         false,
		HasSubscription: false,
		IsAuthor:        false,
		SocialNetworks:  []models.SocialNetwork{{Name: "Facebook", Url: "url"}},
		UpdatedAt:       time.Time{},
	}

	app.DataStore.On("UpdateUser", mock.Anything).Return(nil, errors.New("CONNECTION_FAIL"))

	_, err := UserUseCase{
		datastore: app.DataStore,
	}.UpdateUser(user)

	assert.NotNil(t, err, "CONNECTION_FAIL")
	app.DataStore.MethodCalled("UpdateUser", mock.Anything)
}

func TestDeleteUserIsOk(t *testing.T) {
	app := test.CreateApp()
	Id := "xxx1234"

	app.DataStore.On("DeleteUser", mock.Anything).Return(nil)

	err := UserUseCase{
		datastore: app.DataStore,
	}.DeleteUser(Id)

	assert.Nil(t, err)
	app.DataStore.MethodCalled("DeleteUser", mock.Anything)
}

func TestDeleteUserIsWrongConnectionFailed(t *testing.T) {
	app := test.CreateApp()
	Id := "xxx1234"

	app.DataStore.On("DeleteUser", mock.Anything).Return(errors.New("CONNECTION_FAIL"))

	err := UserUseCase{
		datastore: app.DataStore,
	}.DeleteUser(Id)

	assert.NotNil(t, err, "CONNECTION_FAIL")
	app.DataStore.MethodCalled("DeleteUser", mock.Anything)
}

func TestGetUsersIsOk(t *testing.T) {
	app := test.CreateApp()

	users := []models.User{{
		Id:              "1234567890",
		Email:           "test@example.com",
		Password:        "test1234",
		Name:            "test",
		About:           "test",
		AvatarUrl:       "test",
		CreatedAt:       time.Time{},
		IsAdmin:         false,
		HasSubscription: false,
		IsAuthor:        false,
		SocialNetworks:  []models.SocialNetwork{{Name: "Facebook", Url: "url"}},
		UpdatedAt:       time.Time{}},
	}

	app.DataStore.On("GetUsers", mock.Anything).Return(&users, nil)

	_, err := UserUseCase{
		datastore: app.DataStore,
	}.GetUsers()

	assert.Nil(t, err)
	app.DataStore.MethodCalled("GetUsers", mock.Anything)
}

func TestGetUsersIsWrongConnectionFailed(t *testing.T) {
	app := test.CreateApp()

	app.DataStore.On("GetUsers").Return(nil, errors.New("CONNECTION_FAIL"))

	_, err := UserUseCase{
		datastore: app.DataStore,
	}.GetUsers()

	assert.NotNil(t, err, "CONNECTION_FAIL")
	app.DataStore.MethodCalled("GetUsers")
}

func TestGetUserByIdIsOk(t *testing.T) {
	app := test.CreateApp()

	user := &models.User{
		Id:              "1234567890",
		Email:           "test@example.com",
		Password:        "test1234",
		Name:            "test",
		About:           "test",
		AvatarUrl:       "test",
		CreatedAt:       time.Time{},
		IsAdmin:         false,
		HasSubscription: false,
		IsAuthor:        false,
		SocialNetworks:  []models.SocialNetwork{{Name: "Facebook", Url: "url"}},
		UpdatedAt:       time.Time{},
	}
	Id := "1234567890"

	app.DataStore.On("GetUserById", mock.Anything).Return(user, nil)

	_, err := UserUseCase{
		datastore: app.DataStore,
	}.GetUserById(Id)

	assert.Nil(t, err)
	app.DataStore.MethodCalled("GetUserById", mock.Anything)
}

func TestGetUserByIdIsWrongConnectionFailed(t *testing.T) {
	app := test.CreateApp()
	Id := "xxx1234"

	app.DataStore.On("GetUserById", mock.Anything).Return(nil, errors.New("CONNECTION_FAIL"))

	_, err := UserUseCase{
		datastore: app.DataStore,
	}.GetUserById(Id)

	assert.NotNil(t, err, "CONNECTION_FAIL")
	app.DataStore.MethodCalled("GetUserById", mock.Anything)
}

func TestSaveBookIsOk(t *testing.T) {
	app := test.CreateApp()

	bookDto := &dtos.BookDto{
		Id:           "312312",
		Authors:      []models.Author{{AuthorId: "211212"}},
		AuthorCount:  1,
		Title:        "test",
		AboutTheBook: "test",
		Description:  "test",
		Content: []dtos.BookContentDto{{
			Chapter: "test",
			Sections: []models.BookSection{{
				Id:      "12312",
				Title:   "test",
				Content: "test",
			}},
		}},
		CoverImage:     "test",
		MinimumPrice:   321.12,
		SuggestedPrice: 21312.12,
		Reviews:        10,
		State:          "UNPUBLISHED",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		LanguageName:   "test",
		LanguageCode:   "test",
		Categories:     []string{"test", "test", "test", "test", "test", "test"},
		ReadingOptions: []models.ReadingOption{{
			Option:      "test",
			Description: "test",
		}},
	}
	savedBook := &models.Book{
		Id:           "312312",
		Authors:      []models.Author{{AuthorId: "211212"}},
		AuthorCount:  1,
		Title:        "test",
		AboutTheBook: "test",
		Description:  "test",
		Content: []models.BookContent{{
			Chapter: "test",
			Sections: []models.BookSectionId{{
				SectionId: "12312312312",
			}},
		}},
		CoverImage:     "test",
		MinimumPrice:   321.12,
		SuggestedPrice: 21312.12,
		Reviews:        10,
		State:          "UNPUBLISHED",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		LanguageName:   "test",
		LanguageCode:   "test",
		Categories:     []string{"test", "test", "test", "test", "test", "test"},
		ReadingOptions: []models.ReadingOption{{
			Option:      "test",
			Description: "test",
		}},
	}

	app.DataStore.On("SaveBook", mock.Anything).Return(savedBook, nil)
	app.DataStore.On("SaveBookSections", mock.Anything).Return(nil)

	_, err := BookUseCase{
		datastore: app.DataStore,
	}.SaveBook(bookDto)

	assert.Nil(t, err)
	app.DataStore.MethodCalled("SaveBook", mock.Anything)
	app.DataStore.MethodCalled("SaveBookSections", mock.Anything)
}

func TestSaveBookIsWrongConnectionFailed(t *testing.T) {
	app := test.CreateApp()

	bookDto := &dtos.BookDto{
		Id:           "312312",
		Authors:      []models.Author{{AuthorId: "211212"}},
		AuthorCount:  1,
		Title:        "test",
		AboutTheBook: "test",
		Description:  "test",
		Content: []dtos.BookContentDto{{
			Chapter: "test",
			Sections: []models.BookSection{{
				Id:      "12312",
				Title:   "test",
				Content: "test",
			}},
		}},
		CoverImage:     "test",
		MinimumPrice:   321.12,
		SuggestedPrice: 21312.12,
		Reviews:        10,
		State:          "UNPUBLISHED",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		LanguageName:   "test",
		LanguageCode:   "test",
		Categories:     []string{"test", "test", "test", "test", "test", "test"},
		ReadingOptions: []models.ReadingOption{{
			Option:      "test",
			Description: "test",
		}},
	}

	app.DataStore.On("SaveBook", mock.Anything).Return(nil, errors.New("CONNECTION_FAIL"))
	app.DataStore.On("SaveBookSections", mock.Anything).Return(nil)

	_, err := BookUseCase{
		datastore: app.DataStore,
	}.SaveBook(bookDto)

	assert.NotNil(t, err, "CONNECTION_FAIL")
	app.DataStore.MethodCalled("SaveBook", mock.Anything)
	app.DataStore.MethodCalled("SaveBookSections", mock.Anything)
}

func TestSaveBookSectionsIsOk(t *testing.T) {
	app := test.CreateApp()

	bookSections := []interface{}{
		bson.D{{"id", "12312312"}, {"tile", "test 1"}, {"content", "test 1"}},
		bson.D{{"id", "312312"}, {"title", "test 2"}, {"content", "test 2"}},
	}

	app.DataStore.On("SaveBookSections", mock.Anything).Return(nil)

	err := BookUseCase{
		datastore: app.DataStore,
	}.SaveBookSections(bookSections)

	assert.Nil(t, err)
	app.DataStore.MethodCalled("SaveBookSections", mock.Anything)
}

func TestGetBooksIsOk(t *testing.T) {
	app := test.CreateApp()

	book := &[]models.Book{{
		Id:           "312312",
		Authors:      []models.Author{{AuthorId: "211212"}},
		AuthorCount:  1,
		Title:        "test",
		AboutTheBook: "test",
		Description:  "test",
		Content: []models.BookContent{{
			Chapter: "test",
			Sections: []models.BookSectionId{{
				SectionId: "12312312312",
			}},
		}},
		CoverImage:     "test",
		MinimumPrice:   321.12,
		SuggestedPrice: 21312.12,
		Reviews:        10,
		State:          "UNPUBLISHED",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		LanguageName:   "test",
		LanguageCode:   "test",
		Categories:     []string{"test", "test", "test", "test", "test", "test"},
		ReadingOptions: []models.ReadingOption{{
			Option:      "test",
			Description: "test",
		}},
	}}

	app.DataStore.On("GetBooks", mock.Anything).Return(book, nil)

	_, err := BookUseCase{
		datastore: app.DataStore,
	}.GetBooks()

	assert.Nil(t, err)
	app.DataStore.MethodCalled("GetBooks", mock.Anything)
}

func TestGetBooksIsWrongConnectionFailed(t *testing.T) {
	app := test.CreateApp()

	app.DataStore.On("GetBooks", mock.Anything).Return(nil, errors.New("CONNECTION_FAIL"))

	_, err := BookUseCase{
		datastore: app.DataStore,
	}.GetBooks()

	assert.NotNil(t, err, "CONNECTION_FAIL")
	app.DataStore.MethodCalled("GetBooks", mock.Anything)
}

func TestGetBookIndex(t *testing.T) {
	app := test.CreateApp()

	id := "12312312"
	bookIndex := &models.BookIndex{
		Content: []models.BookContent{{
			Chapter: "test",
			Sections: []models.BookSectionId{{
				SectionId: "312312",
			}},
		}},
	}

	section := &models.BookSection{
		Id: "21312",
		Title: "test",
		Content: "test",
	}

	app.DataStore.On("GetBookIndex", mock.Anything).Return(bookIndex, nil)
	app.DataStore.On("GetBookSectionById", mock.Anything).Return(section, nil)

	_, err := BookUseCase{
		datastore: app.DataStore,
	}.GetBookIndex(id)

	assert.Nil(t, err)
	app.DataStore.MethodCalled("GetBookIndex", mock.Anything)
	app.DataStore.MethodCalled("GetBookSectionById", mock.Anything)
}

func TestGetBookIndexWrongConnectionFailed(t *testing.T) {
	app := test.CreateApp()

	id := "12312312"
	app.DataStore.On("GetBookIndex", mock.Anything).Return(nil, errors.New("CONNECTION_FAIL"))

	_, err := BookUseCase{
		datastore: app.DataStore,
	}.GetBookIndex(id)

	assert.NotNil(t, err, "CONNECTION_FAIL")
	app.DataStore.MethodCalled("GetBookIndex", mock.Anything)
}

func TestGetSectionsByBookIdIsOk(t *testing.T) {
	app := test.CreateApp()

	bookId := "12312312"
	sections := &models.BookSections{
		Sections: []models.BookSection{{
			Id: "312312",
			Title: "test",
			Content: "test",
		}},
	}

	app.DataStore.On("GetSectionsByBookId", mock.Anything).Return(sections, nil)

	_, err := BookUseCase{
		datastore: app.DataStore,
	}.GetSectionsByBookId(bookId)

	assert.Nil(t, err)
	app.DataStore.MethodCalled("GetSectionsByBookId", mock.Anything)
}

func TestGetSectionsByBookIdWrongConnectionFailed(t *testing.T) {
	app := test.CreateApp()
	bookId := "12312312"

	app.DataStore.On("GetSectionsByBookId", mock.Anything).Return(nil, errors.New("CONNECTION_FAIL"))

	_, err := BookUseCase{
		datastore: app.DataStore,
	}.GetSectionsByBookId(bookId)

	assert.NotNil(t, err, "CONNECTION_FAIL")
	app.DataStore.MethodCalled("GetSectionsByBookId", mock.Anything)
}

func TestGetBookSectionByIdIsOk(t *testing.T) {
	app := test.CreateApp()

	bookSection := &models.BookSection{
		Id: "31231212",
		Title: "test",
		Content: "test",
	}
	id := "21312312"

	app.DataStore.On("GetBookSectionById", mock.Anything).Return(bookSection, nil)

	_, err := BookUseCase{
		datastore: app.DataStore,
	}.GetBookSectionById(id)

	assert.Nil(t, err)
	app.DataStore.MethodCalled("GetBookSectionById", mock.Anything)
}

func TestGetBookSectionByIdWrongConnectionFailed(t *testing.T) {
	app := test.CreateApp()

	id := "21312312"

	app.DataStore.On("GetBookSectionById", mock.Anything).Return(nil, errors.New("CONNECTION_FAIL"))

	_, err := BookUseCase{
		datastore: app.DataStore,
	}.GetBookSectionById(id)

	assert.NotNil(t, err, errors.New("CONNECTION_FAIL"))
	app.DataStore.MethodCalled("GetBookSectionById", mock.Anything)
}

func TestGetBookByIdIsOk(t *testing.T) {
	app := test.CreateApp()

	book := &models.Book{
		Id:           "312312",
		Authors:      []models.Author{{AuthorId: "211212"}},
		AuthorCount:  1,
		Title:        "test",
		AboutTheBook: "test",
		Description:  "test",
		Content: []models.BookContent{{
			Chapter: "test",
			Sections: []models.BookSectionId{{
				SectionId: "12312312312",
			}},
		}},
		CoverImage:     "test",
		MinimumPrice:   321.12,
		SuggestedPrice: 21312.12,
		Reviews:        10,
		State:          "UNPUBLISHED",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		LanguageName:   "test",
		LanguageCode:   "test",
		Categories:     []string{"test", "test", "test", "test", "test", "test"},
		ReadingOptions: []models.ReadingOption{{
			Option:      "test",
			Description: "test",
		}},
	}
	id := "21312312"

	app.DataStore.On("GetBookById", mock.Anything).Return(book, nil)

	_, err := BookUseCase{
		datastore: app.DataStore,
	}.GetBookById(id)

	assert.Nil(t, err)
	app.DataStore.MethodCalled("GetBookById", mock.Anything)
}

func TestGetBookByIdWrongConnectionFailed(t *testing.T) {
	app := test.CreateApp()

	id := "21312312"

	app.DataStore.On("GetBookById", mock.Anything).Return(nil, errors.New("CONNECTION_FAIL"))

	_, err := BookUseCase{
		datastore: app.DataStore,
	}.GetBookById(id)

	assert.NotNil(t, err, errors.New("CONNECTION_FAIL"))
	app.DataStore.MethodCalled("GetBookById", mock.Anything)
}

func TestGetBooksByAuthorIsOk(t *testing.T) {
	app := test.CreateApp()

	books := &[]models.Book{{
		Id:           "312312",
		Authors:      []models.Author{{AuthorId: "211212"}},
		AuthorCount:  1,
		Title:        "test",
		AboutTheBook: "test",
		Description:  "test",
		Content: []models.BookContent{{
			Chapter: "test",
			Sections: []models.BookSectionId{{
				SectionId: "12312312312",
			}},
		}},
		CoverImage:     "test",
		MinimumPrice:   321.12,
		SuggestedPrice: 21312.12,
		Reviews:        10,
		State:          "UNPUBLISHED",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		LanguageName:   "test",
		LanguageCode:   "test",
		Categories:     []string{"test", "test", "test", "test", "test", "test"},
		ReadingOptions: []models.ReadingOption{{
			Option:      "test",
			Description: "test",
		}},
	}}
	authorId := "21312312"

	app.DataStore.On("GetBooksByAuthor", mock.Anything).Return(books, nil)

	_, err := BookUseCase{
		datastore: app.DataStore,
	}.GetBooksByAuthor(authorId)

	assert.Nil(t, err)
	app.DataStore.MethodCalled("GetBooksByAuthor", mock.Anything)
}

func TestGetBooksByAuthorWrongConnectionFailed(t *testing.T) {
	app := test.CreateApp()

	authorId := "21312312"

	app.DataStore.On("GetBooksByAuthor", mock.Anything).Return(nil, errors.New("CONNECTION_FAIL"))

	_, err := BookUseCase{
		datastore: app.DataStore,
	}.GetBooksByAuthor(authorId)

	assert.NotNil(t, err, errors.New("CONNECTION_FAIL"))
	app.DataStore.MethodCalled("GetBooksByAuthor", mock.Anything)
}

func TestGetBooksByCategoryIsOk(t *testing.T) {
	app := test.CreateApp()

	books := &[]models.Book{{
		Id:           "312312",
		Authors:      []models.Author{{AuthorId: "211212"}},
		AuthorCount:  1,
		Title:        "test",
		AboutTheBook: "test",
		Description:  "test",
		Content: []models.BookContent{{
			Chapter: "test",
			Sections: []models.BookSectionId{{
				SectionId: "12312312312",
			}},
		}},
		CoverImage:     "test",
		MinimumPrice:   321.12,
		SuggestedPrice: 21312.12,
		Reviews:        10,
		State:          "UNPUBLISHED",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		LanguageName:   "test",
		LanguageCode:   "test",
		Categories:     []string{"test", "test", "test", "test", "test", "test"},
		ReadingOptions: []models.ReadingOption{{
			Option:      "test",
			Description: "test",
		}},
	}}
	category := "Go"

	app.DataStore.On("GetBooksByCategory", mock.Anything).Return(books, nil)

	_, err := BookUseCase{
		datastore: app.DataStore,
	}.GetBooksByCategory(category)

	assert.Nil(t, err)
	app.DataStore.MethodCalled("GetBooksByCategory", mock.Anything)
}

func TestGetBooksByCategoryWrongConnectionFailed(t *testing.T) {
	app := test.CreateApp()

	category := "Go"

	app.DataStore.On("GetBooksByCategory", mock.Anything).Return(nil, errors.New("CONNECTION_FAIL"))

	_, err := BookUseCase{
		datastore: app.DataStore,
	}.GetBooksByCategory(category)

	assert.NotNil(t, err, errors.New("CONNECTION_FAIL"))
	app.DataStore.MethodCalled("GetBooksByCategory", mock.Anything)
}

func TestDeleteBookIsOk(t *testing.T) {
	app := test.CreateApp()

	id := "21312312"

	app.DataStore.On("DeleteBook", mock.Anything).Return(nil)

	err := BookUseCase{
		datastore: app.DataStore,
	}.DeleteBook(id)

	assert.Nil(t, err)
	app.DataStore.MethodCalled("DeleteBook", mock.Anything)
}

func TestDeleteBookWrongConnectionFailed(t *testing.T) {
	app := test.CreateApp()

	id := "21312312"

	app.DataStore.On("DeleteBook", mock.Anything).Return(errors.New("CONNECTION_FAIL"))

	err := BookUseCase{
		datastore: app.DataStore,
	}.DeleteBook(id)

	assert.NotNil(t, err, errors.New("CONNECTION_FAIL"))
	app.DataStore.MethodCalled("DeleteBook", mock.Anything)
}

func TestUpdateBookIsOk(t *testing.T) {
	app := test.CreateApp()

	book := &models.Book{
		Id:           "312312",
		Authors:      []models.Author{{AuthorId: "211212"}},
		AuthorCount:  1,
		Title:        "test",
		AboutTheBook: "test",
		Description:  "test",
		Content: []models.BookContent{{
			Chapter: "test",
			Sections: []models.BookSectionId{{
				SectionId: "12312312312",
			}},
		}},
		CoverImage:     "test",
		MinimumPrice:   321.12,
		SuggestedPrice: 21312.12,
		Reviews:        10,
		State:          "UNPUBLISHED",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		LanguageName:   "test",
		LanguageCode:   "test",
		Categories:     []string{"test", "test", "test", "test", "test", "test"},
		ReadingOptions: []models.ReadingOption{{
			Option:      "test",
			Description: "test",
		}},
	}

	app.DataStore.On("UpdateBook", mock.Anything).Return(book, nil)

	_, err := BookUseCase{
		datastore: app.DataStore,
	}.UpdateBook(book)

	assert.Nil(t, err)
	app.DataStore.MethodCalled("UpdateBook", mock.Anything)
}

func TestUpdateBookWrongConnectionFailed(t *testing.T) {
	app := test.CreateApp()

	book := &models.Book{
		Id:           "312312",
		Authors:      []models.Author{{AuthorId: "211212"}},
		AuthorCount:  1,
		Title:        "test",
		AboutTheBook: "test",
		Description:  "test",
		Content: []models.BookContent{{
			Chapter: "test",
			Sections: []models.BookSectionId{{
				SectionId: "12312312312",
			}},
		}},
		CoverImage:     "test",
		MinimumPrice:   321.12,
		SuggestedPrice: 21312.12,
		Reviews:        10,
		State:          "UNPUBLISHED",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		LanguageName:   "test",
		LanguageCode:   "test",
		Categories:     []string{"test", "test", "test", "test", "test", "test"},
		ReadingOptions: []models.ReadingOption{{
			Option:      "test",
			Description: "test",
		}},
	}

	app.DataStore.On("UpdateBook", mock.Anything).Return(nil, errors.New("CONNECTION_FAIL"))

	_, err := BookUseCase{
		datastore: app.DataStore,
	}.UpdateBook(book)

	assert.NotNil(t, err, errors.New("CONNECTION_FAIL"))
	app.DataStore.MethodCalled("UpdateBook", mock.Anything)
}