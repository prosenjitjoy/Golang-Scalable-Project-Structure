package controller

import (
	"main/core/database"
	"main/core/model"
	"time"
)

type Controller struct {
	repository database.Interface
}

func NewController(repository database.Interface) Interface {
	return &Controller{
		repository: repository,
	}
}

func (c *Controller) ListAll() ([]model.Product, error) {
	// perform additional logic based on different database api
	// example - for dynamodb convert dynamodb attribute to struct
	// example - validate id
	return c.repository.FindAll()
}

func (c *Controller) ListOne(id string) (model.Product, error) {
	// perform additional logic based on different database api
	// example - for dynamodb convert dynamodb attribute to struct
	// example - validate id
	return c.repository.FindOne(id)
}

func (c *Controller) Create(product *model.Product) (string, error) {
	product.CreatedAt = time.Now()
	err := c.repository.Create(*product)
	return product.ID, err
}

func (c *Controller) Update(id string, product *model.Product) error {
	product.UpdatedAt = time.Now()
	return c.repository.Update(id, *product)
}

func (c *Controller) Remove(id string) error {
	// perform additional logic based on different database api
	// example - for dynamodb convert dynamodb attribute to struct
	// example - validate id
	return c.repository.Delete(id)
}
