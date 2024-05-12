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

	mux.HandleFunc("GET /categories", categoryEndpoint.GetCategories)
	mux.HandleFunc("POST /category", categoryEndpoint.AddCategory)
	mux.HandleFunc("PUT /category", categoryEndpoint.UpdateCategory)
	mux.HandleFunc("DELETE /category", categoryEndpoint.DeleteCategory)

	mux.HandleFunc("GET /products", productEndpoint.GetProducts)
	mux.HandleFunc("GET /product", productEndpoint.GetProduct)
	mux.HandleFunc("POST /product", productEndpoint.AddProduct)
	mux.HandleFunc("PUT /product", productEndpoint.UpdateProduct)
	mux.HandleFunc("DELETE /product", productEndpoint.DeleteProduct)
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
