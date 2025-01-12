package handler

import (
	"context"

	"github.com/Euclid0192/golang-grpc/services/common/genproto/orders"
	"github.com/Euclid0192/golang-grpc/services/orders/types"
	"google.golang.org/grpc"
)

type OrdersGrpcHanlder struct {
	// service injection
	ordersServices types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrdersService(grpc *grpc.Server, grpcOrdersService types.OrderService) {
	gRPCHandler := &OrdersGrpcHanlder{
		ordersServices: grpcOrdersService,
	}

	// register the OrderService Server
	orders.RegisterOrderServiceServer(grpc, gRPCHandler)
}

func (h *OrdersGrpcHanlder) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderId:    42,
		CustomerID: 2,
		ProductID:  1,
		Quantity:   10,
	}

	/// Send to service -> service interact with DB (store, edit...)
	err := h.ordersServices.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}

	return res, nil

}
