# Arquitectura del Sistema - Sistema de Gestión de E-commerce (Golang)

**Estudiante:** Elizabeth Cardona  
**Asignatura:** Programación Orientada a Objetos  
**Institución:** Universidad Internacional del Ecuador (UIDE)  
**Fecha:** 28 de junio de 2026

---

## Diagrama de Arquitectura General
[Cliente HTTP]
|
v
[net/http Server :8080]
|
v
[api.Handlers] --> [services.OrderService]
|
+---------------+---------------+
| | |
v v v
[map[string]IProduct] [sync.RWMutex] [encoding/json]
|
v
[models.Product]


---

## Mapeo POO Clásico → Go

| Concepto POO | Implementación en Go | Beneficio |
|-------------|---------------------|-----------|
| Clase | `struct` + métodos con receptor | Claridad estructural, sin herencia oculta |
| Encapsulamiento | Campos `minúsculas` = privados | Visibilidad explícita, fácil de auditar |
| Constructor | Función `NewXxx()` que valida | Estado consistente desde el inicio |
| Interfaz | Contrato implícito por métodos | Composición sobre herencia, bajo acoplamiento |
| Excepciones | Retorno `(value, error)` | Flujo de control explícito y predecible |
| Polimorfismo | Interfaces + implementación implícita | Sustituibilidad sin acoplamiento concreto |

---

## Estrategia de Encapsulamiento en Go

```go
// internal/models/product.go
type Product struct {
    id    string      // Privado: no accesible desde otros paquetes
    name  string      // Privado
    price float64     // Privado
    stock int         // Privado
}

// Constructor con validación de invariantes
func NewProduct(id, name string, price float64, stock int) (IProduct, error) {
    p := &Product{id: id, name: name, price: price, stock: stock}
    if err := p.validate(); err != nil {
        return nil, err  // Estado inválido nunca escapa del constructor
    }
    return p, nil
}

// Validación privada de invariantes de dominio
func (p *Product) validate() error {
    if p.price <= 0 {
        return fmt.Errorf("invariant: price must be > 0")
    }
    if p.stock < 0 {
        return fmt.Errorf("invariant: stock cannot be negative")
    }
    return nil
}

// Getters públicos: única vía de acceso al estado interno
func (p *Product) GetPrice() float64 { return p.price }

// Mutación controlada con pre-validación de negocio
func (p *Product) ReduceStock(qty int) error {
    if !p.IsAvailable(qty) {
        return fmt.Errorf("insufficient stock for product %s", p.id)
    }
    p.stock -= qty  // Única mutación permitida del campo stock
    return nil
}
