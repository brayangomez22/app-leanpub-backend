package app

import "net/http"

func enableCore(w *http.ResponseWriter) {
	(*w).Header().Set("Access-control-allow-Methods", "*")
	(*w).Header().Set("Access-control-allow-Origin", "*")
	(*w).Header().Set("Access-control-allow-Headers", "*")
}

func (app Application) Setup() {
	app.datastore.Setup()
	app.Router.Use(app.routeMiddleware)
	app.Router.HandleFunc("/users", app.SaveUser).Methods(http.MethodPost, http.MethodOptions)
	app.Router.HandleFunc("/users/validate", app.ValidateUser).Methods(http.MethodPost, http.MethodOptions)
	app.Router.HandleFunc("/users", app.GetUsers).Methods(http.MethodGet, http.MethodOptions)
	app.Router.HandleFunc("/users/{id}", app.GetUserById).Methods(http.MethodGet, http.MethodOptions)
	app.Router.HandleFunc("/users/{id}", app.DeleteUser).Methods(http.MethodDelete, http.MethodOptions)
	app.Router.HandleFunc("/users", app.UpdateUser).Methods(http.MethodPut, http.MethodOptions)
	app.Router.HandleFunc("/books", app.SaveBook).Methods(http.MethodPost, http.MethodOptions)
	app.Router.HandleFunc("/books", app.GetBooks).Methods(http.MethodGet, http.MethodOptions)
	app.Router.HandleFunc("/books/index/{id}", app.GetBookIndex).Methods(http.MethodGet, http.MethodOptions)
	app.Router.HandleFunc("/books/sections/{bookId}", app.GetSectionsByBookId).Methods(http.MethodGet, http.MethodOptions)
	app.Router.HandleFunc("/books/section/{id}", app.GetBookSectionById).Methods(http.MethodGet, http.MethodOptions)
	app.Router.HandleFunc("/books/{id}", app.GetBookById).Methods(http.MethodGet, http.MethodOptions)
	app.Router.HandleFunc("/books/author/{authorId}", app.GetBookByAuthor).Methods(http.MethodGet, http.MethodOptions)
	app.Router.HandleFunc("/books/category/{category}", app.GetBookByCategory).Methods(http.MethodGet, http.MethodOptions)
	app.Router.HandleFunc("/books/{id}", app.DeleteBook).Methods(http.MethodDelete, http.MethodOptions)
	app.Router.HandleFunc("/books", app.UpdateBook).Methods(http.MethodPut, http.MethodOptions)
}

func (app Application) routeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCore(&w)
		if (*r).Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}
