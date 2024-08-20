package shop

import (
	"context"

	"xzdp/biz/dal/mysql"
	shop "xzdp/biz/model/shop"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type ShopListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewShopListService(Context context.Context, RequestContext *app.RequestContext) *ShopListService {
	return &ShopListService{RequestContext: RequestContext, Context: Context}
}

func (h *ShopListService) Run(req *shop.Empty) (resp *[]*shop.ShopType, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	// todo edit your code
	shopTypeList, err := mysql.QueryShopType(h.Context)
	if err != nil {
		return nil, err
	}

	return &shopTypeList, nil
}
