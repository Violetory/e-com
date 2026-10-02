package main

import (
	repo "github.com/Violetory/e-com/internal/adapters/postgresql/sqlc"
	"github.com/jackc/pgx/v5"
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

	// 初始化商品服务和处理器
	productService := products.NewService(repo.New(app.db))
	productHandler := products.NewHandler(productService)

	// 健康检查
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("非常好👍")); err != nil {
			log.Printf("write root response: %v", err)
		}
	})

	// 获取商品列表
	r.Get("/product/list", productHandler.ListProducts)

	// 获取商品详情
	r.Get("/product/get", productHandler.GetProduct)

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
	db     *pgx.Conn
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
