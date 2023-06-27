package service

import (
	"context"
	"github.com/zutim/ztm-demo/common/apis"
	"github.com/zutim/ztm-demo/common/pb/order"
	"github.com/zutim/ztm-demo/common/pb/user"
	"github.com/zutim/ztm-demo/internal/order/repo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrderService struct {
	order.UnimplementedOrderServer
	orderRepo repo.OrderRepo
}

func (o *OrderService) Create(ctx context.Context, in *order.OrderCreateReq) (*order.OrderCreateRes, error) {
	o.orderRepo.Create()
	userRes, err := apis.User().Create(ctx, &user.UserCreateReq{
		Name: in.Name,
	})

	if err != nil {
		return nil, err
	}

	return &order.OrderCreateRes{
		Name: in.Name,
		Age:  in.Age,
		Id:   userRes.Id,
	}, nil
}
func (o *OrderService) Get(ctx context.Context, in *order.OrderGetReq) (*order.OrderCreateRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func NewOrderService(orderRepo repo.OrderRepo) order.OrderServer {
	return &OrderService{
		orderRepo: orderRepo,
	}
}
