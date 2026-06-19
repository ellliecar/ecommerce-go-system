# Arquitectura del Sistema E-commerce en Go

## Diagrama de Clases
```mermaid
classDiagram
    class IProduct {
        <<interface>>
        +GetID() string
        +GetName() string
        +GetPrice() float64
        +GetStock() int
        +IsAvailable(int) bool
        +ReduceStock(int) error
    }
    class Product {
        -id string
        -name string
        -price float64
        -stock int
        +validate() error
    }
    IProduct <|-- Product

    class OrderService {
        -mu sync.RWMutex
        -catalog map
        -orderQueue []
        +ProcessOrder(...) (string, error)
    }

    class Server {
        +HealthCheck()
        +CreateOrder()
        +HandleConcurrency()
        ... (8 servicios)
    }

    OrderService --> IProduct
    Server --> OrderService
