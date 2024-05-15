package main

import (
	product_grpc "client/api/note_v1"
	"client/config"
	"client/internal/category"
	"client/internal/product"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Run() {
	mux := http.NewServeMux()
	timeout := 10

	config, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	conn, err := grpc.Dial(config.GrpcServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	categoryClient := product_grpc.NewCategoryClient(conn)

	categoryRepository := category.NewRepository(categoryClient)
	categoryService := category.NewService(categoryRepository)
	categoryEndpoint := category.NewEndpoint(categoryService, logger)

	productClient := product_grpc.NewProductClient(conn)

	productRepository := product.NewRepository(productClient)
	productService := product.NewService(productRepository)
	productEndpoint := product.NewEndpoint(productService, logger)

	mux.HandleFunc("GET /v1/categories", categoryEndpoint.GetCategories)
	mux.HandleFunc("POST /v1/categories", categoryEndpoint.AddCategory)
	mux.HandleFunc("PUT /v1/categories", categoryEndpoint.UpdateCategory)
	mux.HandleFunc("DELETE /v1/categories/{id}", categoryEndpoint.DeleteCategory)

	mux.HandleFunc("GET /v1/products", productEndpoint.GetProducts)
	mux.HandleFunc("GET /v1/products/{id}", productEndpoint.GetProduct)
	mux.HandleFunc("POST /v1/products", productEndpoint.AddProduct)
	mux.HandleFunc("PUT /v1/products", productEndpoint.UpdateProduct)
	mux.HandleFunc("DELETE /v1/products/{id}", productEndpoint.DeleteProduct)
	//nolint
	srv := http.Server{
		Addr:              config.HTTPPort,
		Handler:           mux,
		ReadHeaderTimeout: time.Second * time.Duration(timeout),
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
