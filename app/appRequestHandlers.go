package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"leanpub-app/domain/models"
	"leanpub-app/domain/models/dtos"
	"net/http"
)

func (app Application) SaveUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userSaved, err := app.userUseCases.SaveUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(userSaved)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(data)
}

func (app Application) ValidateUser(w http.ResponseWriter, r *http.Request) {
	var userData models.RegisteredUser
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	validateUser, err := app.userUseCases.ValidateUser(&userData, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(validateUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(data)
}

func (app Application) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := app.userUseCases.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(data)
}

func (app Application) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	user, err := app.userUseCases.GetUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(data)
}

func (app Application) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := app.userUseCases.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app Application) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Id == "" {
		http.Error(w, "USER_NOT_FOUND", http.StatusBadRequest)
		return
	}

	updatedUser, err := app.userUseCases.UpdateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(data)
}

func (app Application) SaveBook(w http.ResponseWriter, r *http.Request) {
	var book dtos.BookDto
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bookSaved, err := app.bookUseCases.SaveBook(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(bookSaved)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(data)
}

func (app Application) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := app.bookUseCases.GetBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(data)
}

func (app Application) GetBookIndex(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	book, err := app.bookUseCases.GetBookIndex(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(data)
}

func (app Application) GetSectionsByBookId(w http.ResponseWriter, r *http.Request) {
	bookId := mux.Vars(r)["bookId"]
	sections, err := app.bookUseCases.GetSectionsByBookId(bookId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(&sections)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(data)
}

func (app Application) GetBookSectionById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	section, err := app.bookUseCases.GetBookSectionById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(&section)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(data)
}

func (app Application) GetBookById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	book, err := app.bookUseCases.GetBookById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(data)
}

func (app Application) GetBooksByAuthor(w http.ResponseWriter, r *http.Request) {
	authorId := mux.Vars(r)["authorId"]
	book, err := app.bookUseCases.GetBooksByAuthor(authorId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(data)
}

func (app Application) GetBooksByCategory(w http.ResponseWriter, r *http.Request) {
	category := mux.Vars(r)["category"]
	books, err := app.bookUseCases.GetBooksByCategory(category)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(&books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(data)
}

func (app Application) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := app.bookUseCases.DeleteBook(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app Application) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if book.Id == "" {
		http.Error(w, "USER_NOT_FOUND", http.StatusBadRequest)
		return
	}

	updatedBook, err := app.bookUseCases.UpdateBook(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(updatedBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(data)
}

func (app Application) SaveShoppingCart(w http.ResponseWriter, r *http.Request)  {
	var shoppingCart models.ShoppingCart
	err := json.NewDecoder(r.Body).Decode(&shoppingCart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shoppingCartSaved, err := app.shoppingCartUseCases.SaveShoppingCart(&shoppingCart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(shoppingCartSaved)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(data)
}

func (app Application) GetShoppingCarts(w http.ResponseWriter, r *http.Request) {
	shoppingCarts, err := app.shoppingCartUseCases.GetShoppingCarts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(shoppingCarts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(data)
}

func (app Application) GetShoppingCartById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	shoppingCart, err := app.shoppingCartUseCases.GetShoppingCartById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(shoppingCart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(data)
}

func (app Application) DeleteShoppingCart(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := app.shoppingCartUseCases.DeleteShoppingCart(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app Application) UpdateShoppingCart(w http.ResponseWriter, r *http.Request) {
	var shoppingCart models.ShoppingCart
	err := json.NewDecoder(r.Body).Decode(&shoppingCart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if shoppingCart.Id == "" {
		http.Error(w, "SHOPPING_CART_NOT_FOUND", http.StatusBadRequest)
		return
	}

	updatedShoppingCart, err := app.shoppingCartUseCases.UpdateShoppingCart(&shoppingCart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(updatedShoppingCart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(data)
}