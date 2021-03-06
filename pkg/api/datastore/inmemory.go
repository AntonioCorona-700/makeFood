package datastore

import (
	"errors"
	"fmt"

	"github.com/AntonioCorona-700/makeFood/pkg/api/model"
)

type InMemory struct {
	quotes map[int]model.Quote
}

func NewInMemory() *InMemory {
	return &InMemory{
		quotes: make(map[int]model.Quote),
	}
}

func (m *InMemory) Create(id int, quote model.Quote) error {
	if _, exists := m.quotes[id]; exists {
		return errors.New(fmt.Sprintf("entry with id: %d already exists", id))
	}

	m.quotes[id] = quote

	return nil
}

func (m *InMemory) Get(id int) (model.Quote, error) {
	return model.Quote{}, errors.New("Unimplemented")
}
