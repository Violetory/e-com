package products

import "context"

type Service interface {
	ListProducts(context context.Context) error
}

type service struct {
	// 数据库
}

func NewService() Service {
	return &service{}
}

func (s *service) ListProducts(context context.Context) error {
	return nil
}
