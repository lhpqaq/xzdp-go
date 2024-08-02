package service

import (
	"context"

	shop "xzdp/biz/model/shop"

	"github.com/cloudwego/hertz/pkg/app"
)

type ShopOfTypeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewShopOfTypeService(Context context.Context, RequestContext *app.RequestContext) *ShopOfTypeService {
	return &ShopOfTypeService{RequestContext: RequestContext, Context: Context}
}

func (h *ShopOfTypeService) Run(req *shop.ShopOfTypeReq) (resp *[]*shop.Shop, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return &[]*shop.Shop{}, nil
}
