package types

import (
	"context"

	"github.com/Euclid0192/golang-grpc/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
}
