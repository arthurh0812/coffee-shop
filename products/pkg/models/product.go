package models

import (
	"encoding/json"
	"io"
	"log"
	"regexp"
	"time"

	"github.com/go-playground/validator"

	"github.com/arthurh0812/coffee-shop/models"
)

type Product struct {
	ID        models.ObjectID `json:"id" validate:"required"`
	CreatorID models.UserID   `json:"creatorID" validate:"required"`

	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"required,gt=0,lte=100"`
	SKU         string  `json:"sku" validate:"required,sku"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `exp.MatchString(s)json:"updatedAt,omitempty"`
}

func (p *Product) Validate() error {
	validate := validator.New()
	err := validate.RegisterValidation("sku", SKUValidation)
	if err != nil {
		log.Fatal(err)
	}
	return validate.Struct(p)
}

func SKUValidation(lvl validator.FieldLevel) bool {
	s := lvl.Field().String()

	exp := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := exp.FindAllString(s, -1)

	if len(matches) != 1 {
		return false
	}
	return true
}

func NewProduct(name, sku string, price float32, uid models.UserID) *Product {
	return &Product{
		ID:        models.NewObjectID(),
		CreatorID: uid,
		Name:      name,
		Price:     price,
		SKU:       sku,
		CreatedAt: time.Now(),
	}
}

func (p *Product) SetDescription(desc string) *Product {
	p.Description = desc
	return p
}

func (p *Product) Update(update *Product) {
	if update.Name != "" {
		p.Name = update.Name
	}
	if update.Description != "" {
		p.Description = update.Description
	}
	if update.Price != 0 {
		p.Price = update.Price
	}
	p.SKU = update.SKU
}

func (p *Product) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(p)
}

type Products []*Product

func (p Products) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(p)
}
