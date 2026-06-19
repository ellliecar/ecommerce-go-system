package errors

import "errors"

// Errores tipados del dominio (Manejo explícito de errores, sin panic)
var (
	ErrProductNotFound   = errors.New("product not found in catalog")
	ErrDuplicateRequest  = errors.New("duplicate request in current session")
	ErrInsufficientStock = errors.New("insufficient stock")
	ErrPaymentFailed     = errors.New("payment processing failed")
)
