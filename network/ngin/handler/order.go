package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zutim/ztm-demo/common/apis"
	"github.com/zutim/ztm-demo/common/pb/order"
	"github.com/zutim/ztm-demo/network/ngin/resp"
	"github.com/zutim/ztm-demo/network/ngin/respCode"
)

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) OrderCreate(ctx *gin.Context) {
	var in *order.OrderCreateReq
	if err := ctx.BindJSON(&in); err != nil {
		resp.FailResponse(ctx, err.Error(), respCode.FAILED)
		return
	}
	res, err := apis.Order().Create(ctx, in)
	if err != nil {
		resp.FailResponse(ctx, err.Error(), respCode.FAILED)
		return
	}
	resp.SuccessResponse(ctx, res)
}
