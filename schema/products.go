package schema

import (
	"encoding/json"
	"io"

	models2 "github.com/arthurh0812/coffee-shop/products/pkg/models"
)

type ProductResponse struct {
	Product *models2.Product `json:"product"`
	Message string           `json:"message"`
	Status  int              `json:"status"`
}

func EncodeProductResponse(w io.Writer, res *ProductResponse) error {
	return json.NewEncoder(w).Encode(res)
}

type ProductsResponse struct {
	Products models2.Products `json:"products"`
	Count    int              `json:"count"`
	Message  string           `json:"message"`
	Status   int              `json:"status"`
}

func EncodeProductsResponse(w io.Writer, res *ProductsResponse) error {
	return json.NewEncoder(w).Encode(res)
}

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	SKU         string  `json:"sku"`
}

func DecodeCreateProductsRequest(r io.Reader) (*CreateProductRequest, error) {
	req := &CreateProductRequest{}
	err := json.NewDecoder(r).Decode(req)
	return req, err
}

type CreateProductResponse struct {
	CreatedProduct *models2.Product `json:"createdProduct"`
	Message        string           `json:"message"`
	Status         int              `json:"status"`
}

func EncodeCreateProductResponse(w io.Writer, res *CreateProductResponse) error {
	return json.NewEncoder(w).Encode(res)
}

type UpdateProductRequest struct {
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
}

func DecodeUpdateProductRequest(r io.Reader) (*UpdateProductRequest, error) {
	req := &UpdateProductRequest{}
	err := json.NewDecoder(r).Decode(req)
	return req, err
}

type UpdateProductResponse struct {
	UpdatedProduct *models2.Product `json:"updatedProduct"`
	Message        string           `json:"message"`
	Status         int              `json:"status"`
}

func EncodeUpdateProductResponse(w io.Writer, res *UpdateProductResponse) error {
	return json.NewEncoder(w).Encode(res)
}

type DeleteProductResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func EncodeDeleteProductResponse(w io.Writer, res *DeleteProductResponse) error {
	return json.NewEncoder(w).Encode(res)
}
