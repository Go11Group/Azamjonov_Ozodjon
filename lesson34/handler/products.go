package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"users_and_produncts/model"
)

// Create
func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	product := model.Product{}
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error with geting json from request body: %s", err)))
		return
	}

	err = h.Products.CreateProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error with adding new product to database: %s", err)))
		return
	}
	w.Write([]byte("yangi product muvaffaqiyatli qo'shildi"))
}

// Read
func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	filter := model.FilterProduct{}
	if query.Has("id") {
		id, err := strconv.Atoi(query.Get("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Error with id is not in correct type request body: %s", err)))
			return
		}
		filter.Id = &id
	}
	if query.Has("name") {
		name := query.Get("name")
		filter.Name = &name
	}
	if query.Has("description") {
		description := query.Get("description")
		filter.Description = &description
	}
	if query.Has("price") {
		price, err := strconv.ParseFloat(query.Get("price"), 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Error with price is not in correct type request body: %s", err)))
			return
		}
		filter.Price = &price
	}
	if query.Has("stock_quantity") {
		stock_quantity, err := strconv.Atoi(query.Get("stock_quantity"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Error with stack_quantity is not in correct type request body: %s", err)))
			return
		}
		filter.Stock_quantity = &stock_quantity
	}
	products, err := h.Products.GetProduct(filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error with getting product from database: %s", err)))
		return
	}
	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error with converting data to json: %s", err)))
		return
	}
}

// Update
func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productToUpdate := model.Product{}
	err := json.NewDecoder(r.Body).Decode(&productToUpdate)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error with Geting json from request body: %s", err)))
		return
	}
	err = h.Products.UpdateProduct(&productToUpdate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error with update product in database: %s", err)))
		return
	}
	w.Write([]byte("Product muvaffaqiyatli o'zgartilildi"))
}

// Delete
func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if !query.Has("id") {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: there is no id in Request Query"))
		return
	}
	id, err := strconv.Atoi(query.Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: id is not in correct type in Request Query"))
		return
	}
	err = h.Products.DeleteProduct(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error with deleting product from database"))
		return
	}
	w.Write([]byte("Product muvaffaqiyatli o'chirildi"))
}
