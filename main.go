package main

import (
	"log"
	"net/http"
	"os"

	muxmiddleware "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/arthurh0812/coffee-shop/env"
	"github.com/arthurh0812/coffee-shop/handlers"
	"github.com/arthurh0812/coffee-shop/middlewares"
	server2 "github.com/arthurh0812/coffee-shop/products/pkg/server"
)

var (
	logger = log.New(os.Stdout, "[INFO] ", log.LstdFlags)

	productsHandler *handlers.Products
	usersHandler    *handlers.Users
	helloHandler    *handlers.Hello
	goodbyeHandler  *handlers.GoodBye
	viewsHandler    *handlers.Views
)

func init() {
	productsHandler = handlers.NewProducts(logger)
	usersHandler = handlers.NewUsers(logger)
	helloHandler = handlers.NewHello(logger)
	goodbyeHandler = handlers.NewGoodBye(logger)
	viewsHandler = handlers.NewViews(logger)
}

func main() {
	router := mux.NewRouter()

	apiRouter := router.Host("api." + env.Env["host"]).Subrouter()
	apiRouter.Use(muxmiddleware.CORS(muxmiddleware.AllowedOrigins([]string{"http://www.localhost:8080"})))

	wwwRouter := router.Host("www." + env.Env["host"]).Subrouter()
	wwwRouter.Use(func(handler http.Handler) http.Handler {
		return middlewares.Print(logger, handler)
	})

	getRouter := wwwRouter.Methods(http.MethodGet).Subrouter()
	getRouter.Path("/").HandlerFunc(viewsHandler.Homepage)
	getRouter.Path("/hello").HandlerFunc(helloHandler.Get)
	getRouter.Path("/goodbye").HandlerFunc(goodbyeHandler.Get)

	productsRouter := apiRouter.Path("/products").Subrouter()
	productsRouter.Methods(http.MethodGet).HandlerFunc(productsHandler.GetAllProducts)
	productsRouter.Methods(http.MethodPost).Handler(productsHandler.PreCreateProduct(http.HandlerFunc(productsHandler.CreateProduct)))

	productsIDRouter := apiRouter.Path("/products/{id:[0-9a-zA-Z]+}").Subrouter()
	productsIDRouter.Use(productsHandler.ExtractID)
	productsIDRouter.Methods(http.MethodGet).HandlerFunc(productsHandler.GetProductByID)
	productsIDRouter.Methods(http.MethodPost).Handler(productsHandler.PreUpdateProduct(http.HandlerFunc(productsHandler.UpdateProductByID)))
	productsIDRouter.Methods(http.MethodDelete).HandlerFunc(productsHandler.DeleteProductByID)

	usersRouter := apiRouter.Path("/users").Subrouter()
	usersRouter.Methods(http.MethodGet).HandlerFunc(usersHandler.GetAllUsers)
	usersRouter.Methods(http.MethodPost).Handler(usersHandler.PreCreateUser(http.HandlerFunc(usersHandler.CreateUser)))

	usersIDRouter := apiRouter.Path("/users/{id:[0-9a-zA-Z]+}").Subrouter()
	usersIDRouter.Use(usersHandler.ExtractID)
	usersIDRouter.Methods(http.MethodGet).HandlerFunc(usersHandler.GetUserByID)
	usersIDRouter.Methods(http.MethodPost).Handler(usersHandler.PreUpdateUser(http.HandlerFunc(usersHandler.UpdateUser)))
	usersIDRouter.Methods(http.MethodDelete).HandlerFunc(usersHandler.DeleteUserByID)

	apiRouter.Path("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		_, err := w.Write([]byte("Hello"))
		if err != nil {
			http.Error(w, "failed to send data", http.StatusInternalServerError)
		}
	})

	srv := server2.NewServer(router)

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	server2.WaitAndGracefulShutdown(srv, logger) // blocks
}
