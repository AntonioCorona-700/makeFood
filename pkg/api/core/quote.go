package core

import (
	"fmt"
	"math/rand"

	"github.com/AntonioCorona-700/makeFood/pkg/api/datastore"
	"github.com/AntonioCorona-700/makeFood/pkg/api/model"
)

var ds datastore.Quotes

func Create(quote model.Quote) (model.Quote, error) {
	id := generateNextId()
	quote.SetId(id)
	err := ds.Create(id, quote)
	if err != nil {
		return model.Quote{}, err
	}
	return quote, nil
}

func generateNextId() int {
	return rand.Int()
}

func init() {
	fmt.Println("Initializing core quotes")

	ds = datastore.NewInMemory()

	fmt.Println("Done with core quotes initialization")
}
