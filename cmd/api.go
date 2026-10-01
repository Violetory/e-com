package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Violetory/e-com/internal/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Mount
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.RequestID)              // 用于速率限制
	r.Use(middleware.ClientIPFromRemoteAddr) // 用于限制和追踪
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) // 恢复功能

	// Timeout middleware
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("非常好👍")); err != nil {
			log.Printf("write root response: %v", err)
		}
	})

	productService := products.NewService()
	productHandler := products.NewHandler(productService)
	r.Get("/product/list", productHandler.ListProducts)

	return r
}

func (app *application) run(handler http.Handler) error {
	server := &http.Server{
		Addr:         app.config.addr,
		Handler:      handler,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("服务启动啦～地址：%s", server.Addr)

	return server.ListenAndServe()
}

type application struct {
	config config
	// logger
	// db driver
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
