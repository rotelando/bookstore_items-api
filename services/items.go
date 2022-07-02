package services

import (
	items "github.com/rotelando/bookstore_items-api/domain/items"
	"github.com/rotelando/bookstore_items-api/utils/errors"
)

var (
	ItemService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, *errors.RestErr)
	Get(string) (*items.Item, *errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item) (*items.Item, *errors.RestErr) {
	return nil, nil
}

func (s *itemsService) Get(itemId string) (*items.Item, *errors.RestErr) {
	return nil, nil
}
