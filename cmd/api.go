package main

import (
	"fmt"
	"net/http"
	"time"

	repo "github.com/Sarthakischauhan/internal/adapters/postgresql/sqlc"
	"github.com/Sarthakischauhan/internal/orders"
	"github.com/Sarthakischauhan/internal/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
)

type application struct {
	config config
	db     *pgx.Conn
}

// add receivers for app object
func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	// middleware layer
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	// timeout when returning a response
	r.Use(middleware.Timeout(60 * time.Second))

	productService := products.NewService(repo.New(app.db), app.db)
	productHandler := products.NewHandler(productService)

	orderService := orders.NewService(repo.New(app.db), app.db)
	orderHandler := orders.NewHandler(orderService)

	r.Get("/products", productHandler.GetProductsHandler)
	r.Post("/create-product", productHandler.AddProductsHandler)

	r.Post("/create-order", orderHandler.CreateOrderHandler)

	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("running a new code"))
	// })

	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Minute,
	}
	fmt.Printf("Running this app now on: %s", app.config.addr)

	return srv.ListenAndServe()
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
