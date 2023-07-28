package controller

import "main/core/model"

type Interface interface {
	ListAll() ([]model.Product, error)
	ListOne(id string) (model.Product, error)
	Create(entity *model.Product) (string, error)
	Update(id string, entity *model.Product) error
	Remove(id string) error
}
