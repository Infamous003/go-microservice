package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Infamous003/go-microservice/data"
	"github.com/go-chi/chi/v5"
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
		http.Error(w, "[ERROR] Unable to marshal products", http.StatusInternalServerError)
		return
	}
}

func (p *Product) GetProductByID(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "[ERROR] Invalid URL", http.StatusBadRequest)
		return
	}

	prod, err := data.GetProductByID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(prod)
	if err != nil {
		http.Error(w, "[ERROR] Unable to marshal product", http.StatusInternalServerError)
		return
	}
}
