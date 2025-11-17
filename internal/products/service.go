package products

import "context"

type Service interface {
	GetProducts(ctx context.Context) error
}

type svc struct {
	// repository
}

func NewService() Service {
	return &svc{}
}

func (s *svc) GetProducts(ctx context.Context) error {
	return nil
}
