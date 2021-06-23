package db

import (
	"fmt"
	"time"

	"github.com/arthurh0812/coffee-shop/models"
	models2 "github.com/arthurh0812/coffee-shop/products/pkg/models"
)

type ProductDB struct {
}

func GetProducts() models2.Products {
	return productList
}

func CreateProduct(p *models2.Product) *models2.Product {
	newProduct := p
	newProduct.CreatedAt = time.Now()
	productList = append(productList, newProduct)
	return newProduct
}

func GetProductByID(id models.ObjectID) *models2.Product {
	for _, product := range productList {
		if product.ID == id {
			return product
		}
	}
	return nil
}

func UpdateProductByID(id models.ObjectID, update *models2.Product) *models2.Product {
	for i, prod := range productList {
		if prod.ID == id {
			productList[i] = update
			update.UpdatedAt = time.Now()
		}
	}
	return update
}

func DeleteProductByID(id models.ObjectID) error {
	toDelete := -1
	for i, p := range productList {
		if p.ID == id {
			toDelete = i
			break
		}
	}
	if toDelete == -1 {
		return fmt.Errorf("could not find product with ID '%s'", id)
	}
	productList = append(productList[:toDelete], productList[toDelete+1:]...)
	return nil
}

var productList = []*models2.Product{
	{
		ID:          models.NewObjectID(),
		Name:        "Latte",
		Price:       2.75,
		Description: "Nice frothy milk coffee for breakfast.",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          models.NewObjectID(),
		Name:        "Espresso",
		Price:       1.99,
		Description: "Short and strong coffee without milk.",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}
