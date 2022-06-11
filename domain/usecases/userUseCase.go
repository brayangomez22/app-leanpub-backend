package usecases

import (
	"errors"
	"leanpub-app/domain"
	"leanpub-app/domain/model"
)

type UserUseCase struct {
	datastore domain.DatabaseGateway
}

func NewUserUseCase(datastore domain.DatabaseGateway) UserUseCase {
	return UserUseCase{
		datastore: datastore,
	}
}

func (userUseCase UserUseCase) SaveUser(user *model.User) (*model.User, error) {
	var (
		registeredUser model.RegisteredUser
		User           model.User
	)
	registeredUser.Email, registeredUser.Password = user.Email, user.Password

	userUseCase.datastore.ValidateUser(&registeredUser, &User)
	if User.Email == user.Email {
		return nil, errors.New("REGISTERED_EMAIL")
	}

	return userUseCase.SaveUser(user)
}

func (userUseCase UserUseCase) ValidateUser(registeredUser *model.RegisteredUser, user *model.User) (*model.User, error) {
	return userUseCase.datastore.ValidateUser(registeredUser, user)
}

func (userUseCase UserUseCase) GetUser() (*[]model.User, error) {
	return userUseCase.datastore.GetUsers()
}

func (userUseCase UserUseCase) DeleteUser(id string) error {
	return userUseCase.datastore.DeleteUser(id)
}

func (userUseCase UserUseCase) UpdateUser(user *model.User) error {
	_, err := userUseCase.datastore.UpdateUser(user)
	return err
}
