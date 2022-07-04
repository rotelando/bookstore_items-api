package items

import (
	"github.com/rotelando/bookstore_items-api/clients/elasticsearch"
	"github.com/rotelando/bookstore_items-api/utils/errors"
)

const (
	indexItems = "items"
)

func (item *Item) Save() *errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, item)
	if err != nil {
		return errors.NewInternalServerError("error when trying to save item")
	}
	item.Id = result.Id
	return nil
}
