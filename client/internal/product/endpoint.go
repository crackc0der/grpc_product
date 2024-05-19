package product

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

type ServiceInterface interface {
	GetProducts(ctx context.Context) ([]*Product, error)
	GetProduct(ctx context.Context, id int64) (*Product, error)
	AddProduct(ctx context.Context, prod *Product) (*Product, error)
	DeleteProduct(ctx context.Context, id int64) (bool, error)
	UpdateProduct(ctx context.Context, prod *Product) (*Product, error)
}

type Endpoint struct {
	service ServiceInterface
	log     *slog.Logger
}

func NewEndpoint(service *Service, log *slog.Logger) *Endpoint {
	return &Endpoint{service: service, log: log}
}

func (e *Endpoint) GetProducts(writer http.ResponseWriter, request *http.Request) {
	products, err := e.service.GetProducts(request.Context())
	if err != nil {
		e.log.Error("error in API's endpoint.GetProducts: " + err.Error())
	}

	if err := json.NewEncoder(writer).Encode(&products); err != nil {
		e.log.Error("error in API's endpoint.GetProducts: " + err.Error())
	}
}

func (e *Endpoint) GetProduct(writer http.ResponseWriter, request *http.Request) {
	param := request.PathValue("id")

	productID, err := strconv.Atoi(param)
	if err != nil {
		e.log.Error("error in API's endpoint.GetProduct: " + err.Error())
	}

	product, err := e.service.GetProduct(request.Context(), int64(productID))
	if err != nil {
		e.log.Error("error in API's enpoint.GetProduct: " + err.Error())
	}

	if err = json.NewEncoder(writer).Encode(&product); err != nil {
		e.log.Error("error in API's enpoint.GetProduct: " + err.Error())
	}
}

func (e *Endpoint) AddProduct(writer http.ResponseWriter, request *http.Request) {
	var prod Product

	if err := json.NewDecoder(request.Body).Decode(&prod); err != nil {
		e.log.Error("error in API's enpoint.AddProduct: " + err.Error())
	}

	product, err := e.service.AddProduct(request.Context(), &prod)
	if err != nil {
		e.log.Error("error in API's enpoint.AddProduct: " + err.Error())
	}

	if err := json.NewEncoder(writer).Encode(product); err != nil {
		e.log.Error("error in API's enpoint.AddProduct: " + err.Error())
	}
}

func (e *Endpoint) DeleteProduct(writer http.ResponseWriter, request *http.Request) {
	param := request.PathValue("id")

	productID, err := strconv.Atoi(param)
	if err != nil {
		e.log.Error("error in API's endpoint.GetProduct: " + err.Error())
	}

	result, err := e.service.DeleteProduct(request.Context(), int64(productID))
	if err != nil {
		e.log.Error("error in API's endpoint.DeleteProduct: " + err.Error())
	}

	if err = json.NewEncoder(writer).Encode(result); err != nil {
		e.log.Error("error in API's endpoint.DeleteProduct: " + err.Error())
	}
}

func (e *Endpoint) UpdateProduct(writer http.ResponseWriter, request *http.Request) {
	var prod Product

	if err := json.NewDecoder(request.Body).Decode(&prod); err != nil {
		e.log.Error("error in API's endpoint.UpdateProduct: " + err.Error())
	}

	product, err := e.service.UpdateProduct(request.Context(), &prod)
	if err != nil {
		e.log.Error("error in API's endpoint.UpdateProduct: " + err.Error())
	}

	if err = json.NewEncoder(writer).Encode(product); err != nil {
		e.log.Error("error in API's endpoint.UpdateProduct: " + err.Error())
	}
}
