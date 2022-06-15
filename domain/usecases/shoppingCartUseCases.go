package usecases

import (
	"leanpub-app/domain"
	"leanpub-app/domain/models"
)

type ShoppingCartUseCase struct {
	datastore domain.DatabaseGateway
}

func NewShoppingCartUseCase(datastore domain.DatabaseGateway) ShoppingCartUseCase {
	return ShoppingCartUseCase{
		datastore: datastore,
	}
}

func (useCase ShoppingCartUseCase) SaveShoppingCart(shoppingCart *models.ShoppingCart) (*models.ShoppingCart, error)  {
	return useCase.datastore.SaveShoppingCart(shoppingCart)
}

func (useCase ShoppingCartUseCase) GetShoppingCarts() (*[]models.ShoppingCart, error) {
	return useCase.datastore.GetShoppingCarts()
}

func (useCase ShoppingCartUseCase) GetShoppingCartById(id string) (*models.ShoppingCart, error) {
	return useCase.datastore.GetShoppingCartById(id)
}

func (useCase ShoppingCartUseCase) DeleteShoppingCart(id string) error {
	return useCase.datastore.DeleteShoppingCart(id)
}

func (useCase ShoppingCartUseCase) UpdateShoppingCart(shoppingCart *models.ShoppingCart)	(*models.ShoppingCart, error) {
	return useCase.datastore.UpdateShoppingCart(shoppingCart)
}