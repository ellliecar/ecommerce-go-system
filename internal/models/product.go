package models

import "fmt"

type Product struct {
	id    string
	name  string
	price float64
	stock int
}

func NewProduct(id, name string, price float64, stock int) (*Product, error) {
	if price <= 0 || stock < 0 {
		return nil, fmt.Errorf("datos inválidos del producto")
	}
	return &Product{id: id, name: name, price: price, stock: stock}, nil
}

func (p *Product) GetID() string    { return p.id }
func (p *Product) GetName() string  { return p.name }
func (p *Product) GetPrice() float64 { return p.price }
func (p *Product) GetStock() int     { return p.stock }

func (p *Product) ReduceStock(qty int) error {
	if qty <= 0 {
		return fmt.Errorf("cantidad inválida")
	}
	if p.stock < qty {
		return fmt.Errorf("stock insuficiente")
	}
	p.stock -= qty
	return nil
}
