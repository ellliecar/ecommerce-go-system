package models

import (
	"fmt"
	"ecommerce-go-system/internal/interfaces"
)

// Product implementa IProduct. Encapsulamiento: campos minúsculos = privados en Go
type Product struct {
	id    string
	name  string
	price float64
	stock int
}

// NewProduct constructor con validación de invariantes (POO: estado consistente desde el inicio)
func NewProduct(id, name string, price float64, stock int) (interfaces.IProduct, error) {
	p := &Product{id: id, name: name, price: price, stock: stock}
	if err := p.validate(); err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Product) validate() error {
	if p.price <= 0 {
		return fmt.Errorf("invariant violation: price must be > 0")
	}
	if p.stock < 0 {
		return fmt.Errorf("invariant violation: stock cannot be negative")
	}
	return nil
}

// Getters (única vía de acceso a estado privado)
func (p *Product) GetID() string        { return p.id }
func (p *Product) GetName() string      { return p.name }
func (p *Product) GetPrice() float64    { return p.price }
func (p *Product) GetStock() int        { return p.stock }

func (p *Product) IsAvailable(qty int) bool {
	return p.stock >= qty
}

func (p *Product) ReduceStock(qty int) error {
	if !p.IsAvailable(qty) {
		return fmt.Errorf("insufficient stock for product %s", p.id)
	}
	p.stock -= qty
	return nil
}
