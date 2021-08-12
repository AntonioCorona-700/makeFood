package datastore

import "github.com/AntonioCorona-700/makeFood/pkg/api/model"

type Quotes interface {
	Create(id int, obj model.Quote) error
	Get(id int) (model.Quote, error)
}
