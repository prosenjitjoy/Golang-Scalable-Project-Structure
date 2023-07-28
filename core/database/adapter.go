package database

import (
	"context"
	"github.com/arangodb/go-driver"
	"main/core/model"
)

type Database struct {
	database   driver.Database
	collection driver.Collection
	logMode    bool
}

func NewAdapter(db driver.Database, collectionName string) Interface {
	return &Database{
		database:   db,
		collection: OpenCollection(db, collectionName),
		logMode:    false,
	}
}

func (db *Database) FindAll() ([]model.Product, error) {
	query := "FOR item IN products RETURN item"
	cursor, err := db.database.Query(context.TODO(), query, nil)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	var products []model.Product
	for {
		var product model.Product
		_, err = cursor.ReadDocument(context.TODO(), &product)

		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, err
		}

		products = append(products, product)
	}
	return products, nil
}

func (db *Database) FindOne(key string) (model.Product, error) {
	var product model.Product
	_, err := db.collection.ReadDocument(context.TODO(), key, &product)
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func (db *Database) Create(product model.Product) error {
	_, err := db.collection.CreateDocument(context.TODO(), product)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) Update(key string, product model.Product) error {
	_, err := db.collection.UpdateDocument(context.TODO(), key, product)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) Delete(key string) error {
	_, err := db.collection.RemoveDocument(context.TODO(), key)
	return err
}
