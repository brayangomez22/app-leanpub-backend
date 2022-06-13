package usecases

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"leanpub-app/app/test"
	"leanpub-app/domain/models"
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