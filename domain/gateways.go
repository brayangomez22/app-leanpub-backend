package domain

import "leanpub-app/domain/model"

type DatabaseGateway interface {
	SaveUser(user *model.User) (*model.User, error)
	ValidateUser(registeredUser *model.RegisteredUser, user *model.User) (*model.User, error)
	GetUsers() (*[]model.User, error)
	DeleteUser(id string) error
	UpdateUser(user *model.User) (*model.User, error)
	Setup()
}
