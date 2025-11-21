package orders

import (
	"errors"
)

var (
	ErrProductNotFound = errors.New("Product not found")
	ErrProductSoldOut  = errors.New("Product is sold out")
)
