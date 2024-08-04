package service

import (
	"context"

	shop "xzdp/biz/model/shop"

	"xzdp/biz/dal/mysql"

	"github.com/cloudwego/hertz/pkg/app"
)

type ShopInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewShopInfoService(Context context.Context, RequestContext *app.RequestContext) *ShopInfoService {
	return &ShopInfoService{RequestContext: RequestContext, Context: Context}
}

func (h *ShopInfoService) Run(id int64) (resp *shop.Shop, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return mysql.QueryByID(h.Context, id)
}
