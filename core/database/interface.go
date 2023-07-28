package database

import "main/core/model"

type Interface interface {
	FindAll() ([]model.Product, error)
	FindOne(key string) (model.Product, error)
	Create(product model.Product) error
	Update(key string, product model.Product) error
	Delete(key string) error
}
