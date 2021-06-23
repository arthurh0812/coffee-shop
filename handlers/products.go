package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/arthurh0812/coffee-shop/db"
	"github.com/arthurh0812/coffee-shop/models"
	models2 "github.com/arthurh0812/coffee-shop/products/pkg/models"
	"github.com/arthurh0812/coffee-shop/schema"
)

type Products struct {
	handler
}

var products *Products

func NewProducts(l *log.Logger) *Products {
	if products == nil { // singleton
		products = &Products{handler: newHandler("Products", l)}
	}
	return products
}

type ProductKey struct{}

type ProductIDKey struct{}

func (p *Products) GetAllProducts(w http.ResponseWriter, _ *http.Request) {
	products := db.GetProducts()
	w.Header().Add("Content-Type", "application/json")
	err := schema.EncodeProductsResponse(w, &schema.ProductsResponse{
		Products: products,
		Count:    len(products),
		Message:  "It's alright!",
		Status:   http.StatusOK,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to send JSON response: %v", err), http.StatusInternalServerError)
	}
}

func (p *Products) GetProductByID(w http.ResponseWriter, req *http.Request) {
	id := req.Context().Value(ProductIDKey{}).(models.ObjectID) // must do type assertion
	product := db.GetProductByID(id)
	p.sendGetByID(w, product)
}

func (p *Products) sendGetByID(w http.ResponseWriter, product *models2.Product) {
	w.Header().Add("Content-Type", "application/json")
	if product == nil {
		p.sendGetByIDError(w, "Product could not be found.", http.StatusNotFound)
		return
	}
	p.sendGetByIDData(w, &schema.ProductResponse{
		Product: product,
		Message: "Successfully retrieved product by ID!",
		Status:  http.StatusOK,
	})

}

func (p *Products) sendGetByIDError(w http.ResponseWriter, msg string, statusCode int) {
	w.WriteHeader(statusCode)
	err := schema.EncodeProductResponse(w, &schema.ProductResponse{
		Product: nil,
		Message: msg,
		Status:  statusCode,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to send JSON response: %v", err), http.StatusInternalServerError)
	}
}

func (p *Products) sendGetByIDData(w http.ResponseWriter, response *schema.ProductResponse) {
	err := schema.EncodeProductResponse(w, response)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to send JSON response: %v", err), http.StatusInternalServerError)
	}
}

func (p *Products) PreCreateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req, err := schema.DecodeCreateProductsRequest(r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to decode JSON from request body: %v", err), http.StatusInternalServerError)
			return
		}
		ctx := context.WithValue(r.Context(), ProductKey{}, req)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func (p *Products) CreateProduct(w http.ResponseWriter, r *http.Request) {
	req := r.Context().Value(ProductKey{}).(*schema.CreateProductRequest) // must do type assertion
	product := models2.NewProduct(req.Name, req.SKU, req.Price, models.NewUserID()).SetDescription(req.Description)
	if err := product.Validate(); err != nil {
		http.Error(w, fmt.Sprintf("validate input: %v", err), http.StatusBadRequest)
		return
	}
	createdProduct := db.CreateProduct(product)
	w.Header().Add("Content-Type", "application/json")
	err := schema.EncodeCreateProductResponse(w, &schema.CreateProductResponse{
		CreatedProduct: createdProduct,
		Message:        "Successfully created new product!",
		Status:         http.StatusCreated,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to send JSON response: %v", err), http.StatusInternalServerError)
	}
}

func (p *Products) PreUpdateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req, err := schema.DecodeUpdateProductRequest(r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to decode JSON: %v", err), http.StatusInternalServerError)
			return
		}
		ctx := context.WithValue(r.Context(), ProductKey{}, req)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func (p *Products) UpdateProductByID(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(ProductIDKey{}).(models.ObjectID)             // must do type assertion
	req := r.Context().Value(ProductKey{}).(*schema.UpdateProductRequest) // must do type assertion
	product := db.GetProductByID(id)
	update := models2.NewProduct(req.Name, product.SKU, req.Price, product.CreatorID).SetDescription(req.Description)
	product.Update(update)
	if err := product.Validate(); err != nil {
		http.Error(w, fmt.Sprintf("validate input: %v", err), http.StatusBadRequest)
		return
	}
	updatedProduct := db.UpdateProductByID(id, product)
	p.sendUpdateResponse(w, updatedProduct)
}

func (p *Products) sendUpdateResponse(w http.ResponseWriter, updatedProduct *models2.Product) {
	w.Header().Add("Content-Type", "application/json")
	if updatedProduct == nil {
		p.sendUpdateProductError(w, "Product could not be found.", http.StatusNotFound)
		return
	}
	p.sendUpdateProductData(w, &schema.UpdateProductResponse{
		UpdatedProduct: updatedProduct,
		Message:        "Successfully modified product",
		Status:         http.StatusOK,
	})

}

func (p *Products) sendUpdateProductError(w http.ResponseWriter, msg string, statusCode int) {
	w.WriteHeader(statusCode)
	err := schema.EncodeUpdateProductResponse(w, &schema.UpdateProductResponse{
		UpdatedProduct: nil,
		Message:        msg,
		Status:         statusCode,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to send JSON response: %v", err), http.StatusInternalServerError)
	}
}

func (p *Products) sendUpdateProductData(w http.ResponseWriter, request *schema.UpdateProductResponse) {
	err := schema.EncodeUpdateProductResponse(w, request)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to send JSON response: %v", err), http.StatusInternalServerError)
	}
}

func (p *Products) ExtractID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		i, ok := mux.Vars(req)["id"]
		if !ok {
			http.Error(w, "please provide an ID parameter in the URI", http.StatusBadRequest)
			return
		}
		id, err := models.ToObjectID(i)
		if err != nil {
			http.Error(w, fmt.Sprintf("invalid product ID: %v", err), http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(req.Context(), ProductIDKey{}, id)
		req = req.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}

func (p *Products) DeleteProductByID(w http.ResponseWriter, req *http.Request) {
	id := req.Context().Value(ProductIDKey{}).(models.ObjectID) // must do type assertion
	err := db.DeleteProductByID(id)
	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		p.sendDeleteProductError(w, err.Error(), http.StatusNotFound)
		return
	}
	p.sendDeleteProductData(w, &schema.DeleteProductResponse{
		Message: fmt.Sprintf("Successfully deleted product with ID '%s'", id),
		Status:  http.StatusOK,
	})
}

func (p *Products) sendDeleteProductError(w http.ResponseWriter, msg string, statusCode int) {
	w.WriteHeader(statusCode)
	err := schema.EncodeDeleteProductResponse(w, &schema.DeleteProductResponse{
		Message: msg,
		Status:  statusCode,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to send JSON: %v", err), http.StatusInternalServerError)
	}
}

func (p *Products) sendDeleteProductData(w http.ResponseWriter, response *schema.DeleteProductResponse) {
	err := schema.EncodeDeleteProductResponse(w, response)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to send JSON: %v", err), http.StatusInternalServerError)
	}
}
