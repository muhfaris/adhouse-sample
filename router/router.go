package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/muhfaris/adhouse-sample/configs"
	ph "github.com/muhfaris/adhouse-sample/product/handler"
	"github.com/muhfaris/adhouse-sample/user/handler"
)

var (
	// GET http method
	GET = []string{http.MethodOptions, http.MethodGet}

	// POST http method
	POST = []string{http.MethodOptions, http.MethodPost}

	// PATCH http method
	PATCH = []string{http.MethodOptions, http.MethodPatch}

	// DELETE http method
	DELETE = []string{http.MethodOptions, http.MethodDelete}
)

func HandlerV1(config *configs.Config, r *mux.Router) {
	v1 := r.PathPrefix("/v1").Subrouter()

	uh := handler.NewUserHandler(config.Connection.DB, config.Connection.Logger)
	v1.Handle("/users", customServe{config.Connection.Logger, uh.AddUserHandler}).Methods(POST...)
	v1.Handle("/login", customServe{config.Connection.Logger, uh.LoginHandler}).Methods(POST...)

	ph := ph.NewProductHandler(config.Connection.DB, config.Connection.Logger)
	v1.Handle("/products", customServe{config.Connection.Logger, ph.GetProductByID}).Methods(GET...)
}
