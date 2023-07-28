package model

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID        string    `json:"_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func (p *Product) GenerateID() {
	p.ID = uuid.NewString()
}

func (p *Product) SetCreatedAt() {
	p.CreatedAt = time.Now()
}

func (p *Product) SetUpdatedAt() {
	p.UpdatedAt = time.Now()
}

func GetTimeFormat() string {
	return "2010-01-02T15:04:05-0700"
}

func InterfaceToModel(data interface{}) (Product, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return Product{}, err
	}

	var product Product
	err = json.Unmarshal(bytes, &product)
	return product, err
}
