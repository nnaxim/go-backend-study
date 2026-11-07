package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	repo "github.com/nnaxim/go-backend-study/internal/adapters/postgresql/sqlc"
	"github.com/nnaxim/go-backend-study/internal/orders"
	"github.com/nnaxim/go-backend-study/internal/products"
)

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.ClientIPFromRemoteAddr)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)


	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("it's alive"))
	})

	productService := products.NewService(repo.New(app.db))
	productHandler := products.NewHandler(productService)

	r.Get("/products", productHandler.ListProducts)
	r.Get("/product/{id}", productHandler.FindProductByID)

	orderService := orders.NewService(repo.New(app.db), app.db)
	orderHandler := orders.NewHandler(orderService)
	r.Post("/order", orderHandler.PlaceOrder)
	return r
}

func (app *application) run(h http.Handler) error {
	srv := http.Server{
		Addr: app.config.addr,
		Handler: h,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 30,
		IdleTimeout: time.Minute,
	}

	fmt.Printf("server has started at addres: %s", app.config.addr)

	return srv.ListenAndServe()
}

type application struct {
	config config
	db *pgx.Conn
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}