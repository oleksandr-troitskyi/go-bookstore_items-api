package items

import (
	"errors"
	"github.com/troitskyA/go-bookstore_items-api/clients/elasticsearch"
	"github.com/troitskyA/go-bookstore_utils-lib/rest_errors"
)

const (
	indexItems = "items"
)

func (i *Item) Save() *rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying ot save item", errors.New("database error"))
	}
	i.Id = result.Id
	return nil
}
