package product

import (
	product_grpc "client/api/note_v1"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type Endpoint struct {
	service *Service
	log     *slog.Logger
}

func NewEndpoint(service *Service, log *slog.Logger) *Endpoint {
	return &Endpoint{service: service, log: log}
}

func (e *Endpoint) GetProducts(writer http.ResponseWriter, request *http.Request) {
	products, err := e.service.GetProducts(context.Background())
	if err != nil {
		e.log.Error("error in endpoint.GetProducts: " + err.Error())
	}

	if err := json.NewEncoder(writer).Encode(&products); err != nil {
		e.log.Error("error in endpoint.GetProducts: " + err.Error())
	}
}

func (e *Endpoint) GetProduct(writer http.ResponseWriter, request *http.Request) {
	id := &product_grpc.ProductRequest{}

	if err := json.NewDecoder(request.Body).Decode(&id); err != nil {
		e.log.Error("error in enpoint.GetProductAAAAAAAAAAA: " + err.Error())
	}
	fmt.Println(id)
	product, err := e.service.GetProduct(context.Background(), id)
	if err != nil {
		e.log.Error("error in enpoint.GetProduct: " + err.Error())
	}

	if err = json.NewEncoder(writer).Encode(&product); err != nil {
		e.log.Error("error in enpoint.GetProduct: " + err.Error())
	}
}

func (e *Endpoint) AddProduct(writer http.ResponseWriter, request *http.Request) {
	prod := &product_grpc.ProductMessage{}

	if err := json.NewDecoder(request.Body).Decode(prod); err != nil {
		e.log.Error("error in enpoint.AddProduct: " + err.Error())
	}

	product, err := e.service.AddProduct(context.Background(), prod)
	if err != nil {
		e.log.Error("error in enpoint.AddProduct: " + err.Error())
	}

	if err := json.NewEncoder(writer).Encode(product); err != nil {
		e.log.Error("error in enpoint.AddProduct: " + err.Error())
	}
}

func (e *Endpoint) DeleteProduct(writer http.ResponseWriter, request *http.Request) {
	id := &product_grpc.ProductRequest{}

	if err := json.NewDecoder(request.Body).Decode(id); err != nil {
		e.log.Error("error in Delete.AddProduct: " + err.Error())
	}

	result, err := e.service.DeleteProduct(context.Background(), id)
	if err != nil {
		e.log.Error("error in Delete.AddProduct: " + err.Error())
	}

	if err = json.NewEncoder(writer).Encode(result); err != nil {
		e.log.Error("error in Delete.AddProduct: " + err.Error())
	}
}

func (e *Endpoint) UpdateProduct(writer http.ResponseWriter, request *http.Request) {
	prod := &product_grpc.ProductMessage{}

	if err := json.NewDecoder(request.Body).Decode(prod); err != nil {
		e.log.Error("error in Update.AddProduct: " + err.Error())
	}

	product, err := e.service.UpdateProduct(context.Background(), prod)
	if err != nil {
		e.log.Error("error in Update.AddProduct: " + err.Error())
	}

	if err = json.NewEncoder(writer).Encode(product); err != nil {
		e.log.Error("error in Update.AddProduct: " + err.Error())
	}
}
