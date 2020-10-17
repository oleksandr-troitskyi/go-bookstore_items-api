package elasticsearch

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"github.com/troitskyA/go-bookstore_utils-lib/logger"
	"net/http"
	"time"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(c *elastic.Client)
	Index(string, interface{}) (*elastic.IndexResponse, error)
}

type esClient struct {
	client *elastic.Client
}

func Init() {
	log := logger.GetLogger()
	client, err := elastic.NewClient(
		elastic.SetURL("http://go_bookstore_items_api.loc:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(log),
		elastic.SetInfoLog(log),
		elastic.SetHeaders(http.Header{
			"X-Caller-Id": []string{"..."},
		}),
	)

	if err != nil {
		panic(err)
	}

	Client.setClient(client)
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

func (c *esClient) Index(index string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().
		Index(index).
		BodyJson(doc).
		Type("item").
		Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to index document in index %s", index), err)
		return nil, err
	}

	return result, nil
}
