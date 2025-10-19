package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Infamous003/go-microservice/data"
)

type Product struct{}

// Product constructor
func NewProduct() *Product {
	return &Product{}
}

func (p *Product) GetProducts(w http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	err := json.NewEncoder(w).Encode(products)

	if err != nil {
		http.Error(w, "[ERROR] Unable to Marshal products", http.StatusInternalServerError)
		return
	}
}
