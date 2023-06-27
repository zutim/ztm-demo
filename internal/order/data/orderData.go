package data

import "github.com/zutim/ztm-demo/internal/order/repo"

type orderData struct {
}

func (o *orderData) Create() {
	return
}

func NewOrderData() repo.OrderRepo {
	return &orderData{}
}
