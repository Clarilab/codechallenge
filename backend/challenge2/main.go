package main

import (
	"github.com/Clarilab/codechallenge/backend/challenge2/database"
	"github.com/savsgio/atreugo/v11"
)

func main() {
	repo := database.NewRepository()
	api := &API{
		repo: repo,
	}

	config := atreugo.Config{
		Addr: "0.0.0.0:8000",
	}

	server := atreugo.New(config)

	server.GET("/search/{name}", api.SearchForName)

	server.POST("/search", api.Search)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

// API implements the rest api handlers
type API struct {
	repo database.Repository
}

func (api *API) SearchForName(ctx *atreugo.RequestCtx) error {
	// Your code goes here
	return nil
}

func (api *API) Search(ctx *atreugo.RequestCtx) error {
	// Your code goes here
	return nil
}
