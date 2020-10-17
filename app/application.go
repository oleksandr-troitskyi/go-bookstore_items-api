package app

import (
	"github.com/gorilla/mux"
	"github.com/troitskyA/go-bookstore_items-api/clients/elasticsearch"
	"net/http"
	"time"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()

	mapUrls()

	srv := &http.Server{
		Handler:      router,
		Addr:         "go_bookstore_items_api.loc:8000",
		WriteTimeout: 500 * time.Second,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}

}
