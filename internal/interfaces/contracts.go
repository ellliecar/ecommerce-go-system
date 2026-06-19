package interfaces

// IProduct define el contrato para entidades de producto (Encapsulamiento + Abstracción)
type IProduct interface {
	GetID() string
	GetName() string
	GetPrice() float64
	GetStock() int
	ReduceStock(qty int) error
	IsAvailable(qty int) bool
}

// IOrderProcessor define el contrato para procesamiento de pedidos (Principio ISP)
type IOrderProcessor interface {
	ProcessOrder(productID string, qty int, userID string) (string, error)
	GetQueueStatus() int
}
