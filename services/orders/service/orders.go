package service /// online business logic, apart from transport or grpc layer
import (
	"context"

	"github.com/Euclid0192/golang-grpc/services/common/genproto/orders"
)

// / MOCK data
var ordersDb = make([]*orders.Order, 0)

type OrderService struct {
	// store
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *orders.Order) error {
	ordersDb = append(ordersDb, order)

	return nil
}
