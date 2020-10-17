package services

import (
	"github.com/troitskyA/go-bookstore_items-api/domain/items"
	"github.com/troitskyA/go-bookstore_utils-lib/rest_errors"
	"net/http"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, *rest_errors.RestErr)
	Get(string) (*items.Item, *rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item) (*items.Item, *rest_errors.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Get(string) (*items.Item, *rest_errors.RestErr) {
	return nil, &rest_errors.RestErr{
		Status:  http.StatusNotImplemented,
		Message: "implement me!",
		Error:   "not_implemented",
	}
}
