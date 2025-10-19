package data

import (
	"errors"
	"time"
)

var (
	ErrProductNotFound = errors.New("Product not found")
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float32
	CreatedOn   string
	UpdatedOn   string
	DeletedOn   string
}

// This type is for a List of Products
type Products []*Product

func GetProducts() Products {
	return productsList
}

func GetProductByID(id int) (*Product, error) {
	for _, p := range productsList {
		if p.ID == id {
			return p, nil
		}
	}
	return nil, ErrProductNotFound
}

var productsList = Products{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and string coffee without milk",
		Price:       1.99,
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Oreo",
		Description: "Milky dark Oreo",
		Price:       1.99,
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
