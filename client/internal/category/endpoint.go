package category

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

type ServiceInterface interface {
	GetCategories(ctx context.Context) ([]*Category, error)
	AddCategory(ctx context.Context, cat *Category) (*Category, error)
	UpdateCategory(ctx context.Context, cat *Category) (*Category, error)
	DeleteCategory(ctx context.Context, id int64) (bool, error)
}

type Endpoint struct {
	service ServiceInterface
	log     *slog.Logger
}

func NewEndpoint(service *Service, log *slog.Logger) *Endpoint {
	return &Endpoint{service: service, log: log}
}

func (e *Endpoint) GetCategories(writer http.ResponseWriter, request *http.Request) {
	categories, err := e.service.GetCategories(request.Context())
	if err != nil {
		e.log.Error("error in API's endpoint.GetCategories: " + err.Error())
	}

	if err = json.NewEncoder(writer).Encode(&categories); err != nil {
		e.log.Error("error in API's endpoint.GetCategories: " + err.Error())
	}
}

func (e *Endpoint) AddCategory(writer http.ResponseWriter, request *http.Request) {
	var cat Category
	if err := json.NewDecoder(request.Body).Decode(&cat); err != nil {
		e.log.Error("error in API's endpoint.AddCategory: " + err.Error())
	}

	category, err := e.service.AddCategory(request.Context(), &cat)
	if err != nil {
		e.log.Error("error in API's endpoint.method AddCategory: " + err.Error())
	}

	if err = json.NewEncoder(writer).Encode(&category); err != nil {
		e.log.Error("error in API's endpoint.AddCategory: " + err.Error())
	}
}

func (e *Endpoint) UpdateCategory(writer http.ResponseWriter, request *http.Request) {
	var cat Category
	if err := json.NewDecoder(request.Body).Decode(&cat); err != nil {
		e.log.Error("error in API's endpoint.UpdateCategory: " + err.Error())
	}

	category, err := e.service.UpdateCategory(request.Context(), &cat)
	if err != nil {
		e.log.Error("error in API's endpoint.UpdateCategory: " + err.Error())
	}

	if err = json.NewEncoder(writer).Encode(&category); err != nil {
		e.log.Error("error in API's endpoint.UpdateCategory: " + err.Error())
	}
}

func (e *Endpoint) DeleteCategory(writer http.ResponseWriter, request *http.Request) {
	param := request.PathValue("id")

	categoryID, err := strconv.Atoi(param)
	if err != nil {
		e.log.Error("error in API's endpoint.DeleteCategory: " + err.Error())
	}

	result, err := e.service.DeleteCategory(request.Context(), int64(categoryID))
	if err != nil {
		e.log.Error("error in API's endpoint.DeleteCategory: " + err.Error())
	}

	if err = json.NewEncoder(writer).Encode(&result); err != nil {
		e.log.Error("error in API's endpoint.DeleteCategory: " + err.Error())
	}
}
