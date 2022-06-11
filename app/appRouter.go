package app

import "net/http"

func enableCore(w *http.ResponseWriter){
	(*w).Header().Set("Access-control-allow-Methods","*")
	(*w).Header().Set("Access-control-allow-Origin","*")
	(*w).Header().Set("Access-control-allow-Headers","*")
}

func (app Application) Setup() {
	app.datastore.Setup()
	app.Router.Use(app.routeMiddleware)
	app.Router.HandleFunc("/users/save", app.SaveUser).Methods(http.MethodPost, http.MethodOptions)
	app.Router.HandleFunc("/users/validate", app.ValidateUser).Methods(http.MethodPost, http.MethodOptions)
	app.Router.HandleFunc("/users/get", app.GetUsers).Methods(http.MethodGet, http.MethodOptions)
	app.Router.HandleFunc("/users/delete/{id}", app.DeleteUser).Methods(http.MethodDelete, http.MethodOptions)
	app.Router.HandleFunc("/users/update", app.UpdateUser).Methods(http.MethodPut, http.MethodOptions)
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