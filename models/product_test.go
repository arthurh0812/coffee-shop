package models

import (
	"testing"

	"github.com/arthurh0812/coffee-shop/products/pkg/models"
)

func TestProduct_Validate(t *testing.T) {
	p := models.NewProduct("", "", 130, "")
	err := p.Validate()
	if err == nil {
		t.Errorf("err should not be nil, got nil")
	}
	t.Log(err)
}
